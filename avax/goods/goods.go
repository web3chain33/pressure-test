package goods

import (
	"crypto/ecdsa"
	"github.com/chendehai/pressure-test/avax"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	evmabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Client struct {
	Evm          *avax.Client
	abiJson      string
	abi          evmabi.ABI
	OwnerKey     *ecdsa.PrivateKey
	OwnerAddr    common.Address
	ContractAddr *common.Address
}

func NewClient(evm *avax.Client, abiJson string, ownerKey *ecdsa.PrivateKey, contractAddrString common.Address) (*Client, error) {
	var c Client
	c.Evm = evm
	c.abiJson = abiJson
	abi, err := evmabi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, errors.Wrap(err, "NewClient evmabi.JSON")
	}
	c.abi = abi
	c.OwnerKey = ownerKey
	c.ContractAddr = &contractAddrString
	c.OwnerAddr = crypto.PubkeyToAddress(ownerKey.PublicKey)
	return &c, nil
}

func NewAccountClient(evmClient *avax.Client, rootKey *ecdsa.PrivateKey, abiJson string, abiBin []byte) (*Client, error) {
	prikey, err := avax.NewAccount()
	if err != nil {
		return nil, errors.Wrap(err, "NewAccountClient NewAccount")
	}
	mintAddr := crypto.PubkeyToAddress(prikey.PublicKey)
	// fmt.Printf("generate addr:%s, prikey:%s\n", mintAddr.String(), hex.EncodeToString(crypto.FromECDSA(prikey)))

	_, err = evmClient.TransferAvax(rootKey, &mintAddr, 50)
	if err != nil {
		return nil, errors.Wrap(err, "NewAccountClient TransferAvax")
	}
	// fmt.Printf("transfer %d avax from %s to %s success, hash:%s\n", 50, testAddr.String(), mintAddr.String(), TaRes.Tx.Hash().String())

	_, contractAddr, err := evmClient.SyncDeployContractByBin(prikey, abiBin)
	if err != nil {
		return nil, errors.Wrap(err, "NewAccountClient SyncDeployContractByBinFile")
	}
	// fmt.Println("deploy contract succeed:" + hash + ",addr:" + contractAddr.String())

	return NewClient(evmClient, abiJson, prikey, *contractAddr)
}

func (c *Client) BatchMint(owner common.Address, ids []*big.Int) (string, uint64, error) {
	data, err := c.abi.Pack("batchMint", owner, ids)
	if err != nil {
		return "", 0, err
	}
	return c.Evm.SendTransaction(avax.TxReq{
		ToAddr: c.ContractAddr,
		FromPK: c.OwnerKey,
		Data:   data,
	})
}

func (c *Client) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	data, err := c.abi.Pack("balanceOf", owner, id)
	if err != nil {
		return nil, err
	}

	resp, err := c.Evm.CallContract(avax.TxReq{
		ToAddr:   c.ContractAddr,
		FromAddr: owner,
		Data:     data,
	})
	if err != nil {
		return nil, err
	}
	ans := big.NewInt(0)
	ans.SetBytes(resp)
	return ans, nil
}

func (c *Client) GetSuccessNum(owner common.Address) (int64, error) {
	data, err := c.abi.Pack("getSuccessNum")
	if err != nil {
		return 0, err
	}

	resp, err := c.Evm.CallContract(avax.TxReq{
		ToAddr:   c.ContractAddr,
		FromAddr: owner,
		Data:     data,
	})
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, nil
	}

	return big.NewInt(0).SetBytes(resp).Int64(), nil
}
