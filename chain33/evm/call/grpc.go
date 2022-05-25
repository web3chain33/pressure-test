package call

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	chainCommon "github.com/33cn/chain33/common"
	"github.com/33cn/chain33/rpc/grpcclient"
	"google.golang.org/grpc"

	chainAddress "github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	evmAbi "github.com/33cn/plugin/plugin/dapp/evm/executor/abi"
	evmtypes "github.com/33cn/plugin/plugin/dapp/evm/types"
)

const (
	yccChainId     = 999
	defaultFeeRate = 100000
)

// CallContract 成功部署后的合约
type CallContract struct {
	ContractAddr string
	ParaName     string
	Abi          string
	DeployerPri  crypto.PrivKey
}

// LocalCreateYCCEVMTx 本地构造ycc的evm交易
func (c *CallContract) LocalCreateYCCEVMTx(parameter string) (*chainTypes.Transaction, error) {
	exec := c.ParaName + evmtypes.ExecutorName
	toAddr := chainAddress.ExecAddress(exec)

	_, packedParameter, err := evmAbi.Pack(parameter, c.Abi, false)
	if err != nil {
		return nil, err
	}

	action := evmtypes.EVMContractAction{
		Para:         packedParameter,
		ContractAddr: c.ContractAddr,
	}

	tx := &chainTypes.Transaction{Execer: []byte(exec), Payload: chainTypes.Encode(&action), Fee: 0, To: toAddr}
	// 十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = yccChainId
	tx.Sign(chainTypes.SECP256K1, c.DeployerPri)

	return tx, nil
}

// LocalCreateYCCEVMGroupTx 本地构造ycc的evm交易组
func (c *CallContract) LocalCreateYCCEVMGroupTx(parameters, privkeys []string) (*chainTypes.Transaction, error) {
	exec := c.ParaName + evmtypes.ExecutorName
	toAddr := chainAddress.ExecAddress(exec)

	parameterCount := len(parameters)
	privkeysCount := len(privkeys)

	txList := make([]*chainTypes.Transaction, 0, parameterCount)
	for i := 0; i < parameterCount; i++ {
		_, packedParameter, err := evmAbi.Pack(parameters[i], c.Abi, false)
		if err != nil {
			return nil, err
		}

		action := evmtypes.EVMContractAction{
			Para:         packedParameter,
			ContractAddr: c.ContractAddr,
		}

		tx := &chainTypes.Transaction{Execer: []byte(exec), Payload: chainTypes.Encode(&action), Fee: 0, To: toAddr}
		// 十倍手续费保证成功
		tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		tx.Nonce = random.Int63()
		tx.ChainID = yccChainId
		txList = append(txList, tx)
	}
	tg, err := chainTypes.CreateTxGroup(txList, 10*defaultFeeRate)
	if err != nil {
		return nil, err
	}
	for j := 0; j < len(txList); j++ {
		if privkeysCount > 0 && privkeysCount == parameterCount {
			tg.SignN(j, chainTypes.SECP256K1, chainUtil.HexToPrivkey(privkeys[j]))
		} else {
			tg.SignN(j, chainTypes.SECP256K1, c.DeployerPri)
		}
	}
	return tg.Tx(), nil
}

type DeployContract struct {
	Endpoint     string
	ParaName     string
	Bin          string
	Abi          string
	Parameter    string
	DeployerPri  crypto.PrivKey
	DeployerAddr string
}

// LocalCreateDeployTx create deploy contract tx
func (d *DeployContract) LocalCreateDeployTx() (*chainTypes.Transaction, error) {
	exec := d.ParaName + evmtypes.ExecutorName
	toAddr := chainAddress.ExecAddress(exec)

	bCode, err := chainCommon.FromHex(d.Bin)
	if err != nil {
		return nil, err
	}
	if d.Parameter != "" {
		packData, err := evmAbi.PackContructorPara(d.Parameter, d.Abi)
		if err != nil {
			return nil, err
		}

		bCode = append(bCode, packData...)
	}

	action := evmtypes.EVMContractAction{
		Code:         bCode,
		ContractAddr: toAddr,
	}

	tx := &chainTypes.Transaction{
		Execer:    []byte(exec),
		Payload:   chainTypes.Encode(&action),
		Signature: nil,
		To:        toAddr,
		ChainID:   yccChainId,
	}
	// 十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()

	tx.Sign(chainTypes.SECP256K1, d.DeployerPri)
	return tx, nil
}

// Deploy contract return txhash,ContractAddr
func (d *DeployContract) Deploy() (string, string, error) {
	// 创建本地合约交易
	tx, err := d.LocalCreateDeployTx()
	if err != nil {
		return "", "", err
	}
	txHash := chainCommon.ToHex(tx.Hash())

	conn, err := grpc.Dial(grpcclient.NewMultipleURL(d.Endpoint), grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpcclient.NewMultipleURL err:", err)
		return "", "", err
	}

	client := chainTypes.NewChain33Client(conn)

	// grpc发送交易
	res, err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", "", err
	}
	if !res.IsOk {
		return "", "", fmt.Errorf("SendTransaction fail %v", string(res.Msg))
	}

	// 获取合约地址
	contractAddr := LocalGetContractAddr(d.DeployerAddr, txHash)

	return txHash, contractAddr, nil
}

func LocalGetContractAddr(caller, txhash string) string {
	return chainAddress.HashToAddress(chainAddress.NormalVer,
		chainAddress.ExecPubKey(caller+strings.TrimPrefix(txhash, "0x"))).String()
}
