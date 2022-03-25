package call

import (
	chainAddress "github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	evmAbi "github.com/33cn/plugin/plugin/dapp/evm/executor/abi"
	evmtypes "github.com/33cn/plugin/plugin/dapp/evm/types"
	"math/rand"
	"time"
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
	//十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = yccChainId
	tx.Sign(chainTypes.SECP256K1, c.DeployerPri)

	return tx, nil
}

// LocalCreateYCCEVMGroupTx 本地构造ycc的evm交易组
func (c *CallContract) LocalCreateYCCEVMGroupTx(parameters []string, privkeys []string) (*chainTypes.Transaction, error) {
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
		//十倍手续费保证成功
		tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		tx.Nonce = random.Int63()
		tx.ChainID = yccChainId
		if privkeysCount > 0 && privkeysCount == parameterCount {
			tx.Sign(chainTypes.SECP256K1, chainUtil.HexToPrivkey(privkeys[i]))
		} else {
			tx.Sign(chainTypes.SECP256K1, c.DeployerPri)
		}
		txList = append(txList, tx)
	}
	tg, err :=  chainTypes.CreateTxGroup(txList, 10 * defaultFeeRate)
	if err != nil {
		return nil, err
	}
	for j :=0; j < len(txList); j++ {
		if privkeysCount > 0 && privkeysCount == parameterCount {
			tg.SignN(j, chainTypes.SECP256K1, chainUtil.HexToPrivkey(privkeys[j]))
		} else {
			tg.SignN(j, chainTypes.SECP256K1, c.DeployerPri)
		}
	}
	return tg.Tx(), nil
}
