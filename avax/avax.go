package avax

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	DefaultGasLimit         = 100000
	WeiPerAvax      float64 = 1000000000000000000
)

var (
	NonceMap     = make(map[string]uint64)
	nonceManager = NonceManager{m: make(map[string]uint64), Mutex: sync.Mutex{}}
)

type NonceManager struct {
	m map[string]uint64
	sync.Mutex
}

func (m *NonceManager) GetNonce(addr common.Address, c *Client) uint64 {
	m.Lock()
	key := addr.String()
	Nonce := NonceMap[key]
	if Nonce == 0 {
		Nonce, _ = c.PendingNonceAt(context.Background(), addr)
	}
	NonceMap[key] = Nonce + 1
	m.Unlock()
	return Nonce
}

type Client struct {
	*ethclient.Client
	chainID  *big.Int
	signer   types.Signer
	GasPrice *big.Int
}

func NewClient(url string) (*Client, error) {
	ec, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	var client Client
	client.Client = ec
	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	client.chainID = chainID
	client.signer = types.LatestSignerForChainID(chainID)

	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	client.GasPrice = gasPrice

	return &client, nil
}

type TxReq struct {
	Nonce    uint64
	FromAddr common.Address
	FromPK   *ecdsa.PrivateKey
	ToAddr   *common.Address
	Data     []byte
	GasLimit uint64
	GasPrice *big.Int
	Value    float64
}

type TxResult struct {
	Tx     *types.Transaction
	Rcpt   *types.Receipt
	Header *types.Header
}

func NewAccount() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return privateKey, err
}

func (c *Client) CallContract(txReq TxReq) ([]byte, error) {
	var msg ethereum.CallMsg
	msg.GasPrice = c.GasPrice
	msg.From = txReq.FromAddr
	msg.Gas = txReq.GasLimit
	if msg.Gas == 0 {
		msg.Gas = DefaultGasLimit
	}
	msg.To = txReq.ToAddr
	msg.Data = txReq.Data
	return c.Client.CallContract(context.Background(), msg, nil)
}

func (c *Client) TransferAvax(fromKey *ecdsa.PrivateKey, to *common.Address, amount float64) (*TxResult, error) {
	return c.SyncSendTransaction(TxReq{
		FromPK:   fromKey,
		ToAddr:   to,
		Value:    amount,
		GasPrice: big.NewInt(27000000000),
	})
}

func (c *Client) SyncSendTransaction(txReq TxReq) (*TxResult, error) {
	hash, _, err := c.SendTransaction(txReq)
	if err != nil {
		return nil, err
	}
	return c.WaitTransaction(hash)
}

func (c *Client) SendTransaction(txReq TxReq) (hash string, Nonce uint64, err error) {
	pk := txReq.FromPK

	sender := crypto.PubkeyToAddress(pk.PublicKey)

	if err != nil {
		return "", 0, err
	}
	gasLimit := txReq.GasLimit
	if gasLimit == 0 {
		gasLimit = DefaultGasLimit
	}
	gasPrice := txReq.GasPrice
	if gasPrice == nil {
		// gasPrice = c.GasPrice
		gasPrice, err = c.SuggestGasPrice(context.Background())
		if err != nil {
			return "", 0, err
		}
	}

	value, _ := big.NewFloat(0).Mul(big.NewFloat(txReq.Value), big.NewFloat(WeiPerAvax)).Int(nil)

	Nonce = nonceManager.GetNonce(sender, c)
	tx, err := types.SignNewTx(pk, c.signer, &types.LegacyTx{
		Nonce:    Nonce,
		To:       txReq.ToAddr,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     txReq.Data,
		Value:    value,
	})
	if err != nil {
		return "", Nonce, err
	}

	err = c.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		//if strings.Contains(err.Error(), "nonce too low") {
		//	Nonce, err = c.PendingNonceAt(context.Background(), sender)
		//	if err != nil {
		//		return "", err
		//	}
		//	return c.SendTransaction(txReq)
		//}
		return "", Nonce, err
	}
	return tx.Hash().String(), Nonce, nil
}

var errWaitTimeOut = errors.New("time out")

func (c *Client) WaitTransaction(hashHex string) (*TxResult, error) {
	hash := common.HexToHash(hashHex)
	ticker := time.NewTicker(time.Millisecond * 200)
	end := time.NewTimer(time.Second * 30)

	for {
		tx, isp, err := c.TransactionByHash(context.Background(), hash)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			fmt.Println("TransactionByHash", err)
			return nil, err
		}
		if !isp {
			rcv, err := c.TransactionReceipt(context.Background(), hash)
			if err != nil && !strings.Contains(err.Error(), "not found") {
				fmt.Println("TransactionReceipt", err)
				return nil, err
			}
			if err == nil && rcv != nil {
				header, err := c.HeaderByHash(context.Background(), rcv.BlockHash)
				if err != nil {
					return nil, err
				}
				return &TxResult{Tx: tx, Rcpt: rcv, Header: header}, nil
			}
		}
		select {
		case <-ticker.C:
			continue
		case <-end.C:
			return nil, errWaitTimeOut
		}
	}
}

func (c *Client) SyncDeployContractByBin(ownerKey *ecdsa.PrivateKey, buf []byte) (hash string, addr *common.Address, err error) {

	data := make([]byte, hex.DecodedLen(len(buf)))
	_, err = hex.Decode(data, buf)
	if err != nil {
		return "", nil, err
	}

	gas, err := c.EstimateGas(context.Background(), ethereum.CallMsg{Data: data, Gas: 4000000, From: crypto.PubkeyToAddress(ownerKey.PublicKey)})
	if err != nil {
		return "", nil, err
	}

	hash, _, err = c.SendTransaction(TxReq{
		FromPK:   ownerKey,
		Data:     data,
		GasLimit: gas,
		GasPrice: big.NewInt(27000000000),
	})
	if err != nil {
		return "", nil, err
	}
	txres, err := c.WaitTransaction(hash)
	if err != nil {
		return "", nil, err
	}
	return hash, &txres.Rcpt.ContractAddress, nil
}
