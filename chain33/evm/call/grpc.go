package call

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	chainCommon "github.com/33cn/chain33/common"
	chainAddress "github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/rpc/grpcclient"
	_ "github.com/33cn/chain33/system/address/eth"
	ethAddr "github.com/33cn/chain33/system/address/eth"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	evmAbi "github.com/33cn/plugin/plugin/dapp/evm/executor/abi"
	evmCommon "github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	evmtypes "github.com/33cn/plugin/plugin/dapp/evm/types"
	"google.golang.org/grpc"
)

const (
	yccChainId     = 999
	defaultFeeRate = 100000
)

var (
	Ty       = int32(chainTypes.SECP256K1)
	AddrType = int32(chainAddress.DefaultID)
)

func InitTy(chianType string) {
	if chianType == "ycc" {
		AddrType = int32(ethAddr.ID)
		// 加载, 确保在evm使能高度前, eth地址驱动已使能
		driver, err := chainAddress.LoadDriver(AddrType, -1)
		if err != nil {
			panic(fmt.Sprintf("address driver must enable before %d", 0))
		}
		evmCommon.InitEvmAddressTypeOnce(driver)
	}
	Ty = chainTypes.EncodeSignID(chainTypes.SECP256K1, AddrType)
}

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
	toAddr, err := chainAddress.GetExecAddress(exec, AddrType)
	if err != nil {
		return nil, err
	}

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
	toAddr, err := chainAddress.GetExecAddress(exec, AddrType)
	if err != nil {
		return nil, err
	}

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
			tg.SignN(j, Ty, chainUtil.HexToPrivkey(privkeys[j]))
		} else {
			tg.SignN(j, Ty, c.DeployerPri)
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
	toAddr, err := chainAddress.GetExecAddress(exec, AddrType)
	if err != nil {
		return nil, err
	}

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

	tx.Sign(Ty, d.DeployerPri)
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
	contractAddr := LocalGetContractAddr(d.DeployerAddr, tx.Hash())

	return txHash, contractAddr, nil
}

func LocalGetContractAddr(caller string, txhash []byte) string {
	return evmCommon.NewContractAddress(*evmCommon.StringToAddress(caller), txhash).String()
}
