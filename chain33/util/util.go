package util

import (
	"context"
	"fmt"
	chainCommon "github.com/33cn/chain33/common"
	chainAddress "github.com/33cn/chain33/common/address"
	ethAddr "github.com/33cn/chain33/system/address/eth"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	evmAbi "github.com/33cn/plugin/plugin/dapp/evm/executor/abi"
	evmCommon "github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	evmTypes "github.com/33cn/plugin/plugin/dapp/evm/types"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	yccChainId     = 999
	defaultFeeRate = chainTypes.DefaultMinFee
	groupSize      = int(chainTypes.MaxTxGroupSize)
)

var cpuNum = runtime.NumCPU()

// 并发的平行链数量
var parasSize = 4

// Ty = signID 用于签名 参考  https://github.com/33cn/chain33/blob/master/types/sign.md 文档
var Ty = int32(chainTypes.SECP256K1)
var AddrType = int32(chainAddress.DefaultID)

// CallContract 需要部署的合约
type DeployeContract struct {
	ParaName    string
	Abi         string
	Bin         string
	Parameter   string
	Client      chainTypes.Chain33Client
	CallAddr    string
	CallPrivkey string
}

// CallContract 成功部署后的合约
type CallContract struct {
	ContractAddr string
	ParaName     string
	Abi          string
}

func SetParasLen(l int) {
	parasSize = l
}

func InitTy(chianType string) {
	if chianType == "ycc" {
		AddrType = int32(ethAddr.ID)
	}
	Ty = chainTypes.EncodeSignID(chainTypes.SECP256K1, AddrType)
}

// Deploy contract return txhash,ContractAddr
func (d *DeployeContract) Deploy() (string, string, error) {
	//创建本地合约交易
	tx, err := d.LocalCreateDeployTx()
	if err != nil {
		return "", "", err
	}
	txHash := chainCommon.ToHex(tx.Hash())

	// grpc发送交易
	res, err := d.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", "", err
	}
	if !res.IsOk {
		return "", "", fmt.Errorf("SendTransaction fail %v", string(res.Msg))
	}

	//获取合约地址
	contractAddr := LocalGetContractAddr(d.CallAddr, tx.Hash())

	return txHash, contractAddr, nil
}

func (d *DeployeContract) LocalCreateDeployTx() (*chainTypes.Transaction, error) {
	exec := d.ParaName + evmTypes.ExecutorName
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

	action := evmTypes.EVMContractAction{
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
	//十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	prikey := chainUtil.HexToPrivkey(d.CallPrivkey)
	tx.Sign(chainTypes.SECP256K1, prikey)

	return tx, nil
}

// LocalGetContractAddr 本地构建合约地址 交易哈希需要去掉0x前缀
// 参考
//func LocalGetContractAddr(caller, txhash string) string {
//	return chainAddress.HashToAddress(chainAddress.NormalVer,
//		chainAddress.ExecPubKey(caller+strings.TrimPrefix(txhash, "0x"))).String()
//}

// LocalGetContractAddr 本地构建合约地址
// 代码仓库 github.com/33cn/plugin v1.67.3-0.20220517092344-565e980cc752
// 参考位置 github.com/33cn/plugin/plugin/dapp/evm/executor/exec.go ;func innerExec ;value contractAddrStr
func LocalGetContractAddr(caller string, txhash []byte) string {
	return evmCommon.NewContractAddress(*evmCommon.StringToAddress(caller), txhash).String()
}

// LocalCreateUnSignYCCEVMTx 本地构造ycc的evm未签名交易
func (c *CallContract) LocalCreateUnSignYCCEVMTx(parameter string) (*chainTypes.Transaction, error) {
	return c.localCreateYCCEVMTx(parameter)
}

// LocalCreateSignYCCEVMTx 本地构造ycc的evm签名交易
func (c *CallContract) LocalCreateSignYCCEVMTx(pristr, parameter string) (*chainTypes.Transaction, error) {
	tx, err := c.localCreateYCCEVMTx(parameter)
	if err != nil {
		return nil, err
	}
	prikey := chainUtil.HexToPrivkey(pristr)

	tx.Sign(Ty, prikey)

	return tx, nil
}

//// LocalCreateYCCEVMTx 本地构造ycc的evm默认签名交易
//func (c *CallContract) LocalCreateYCCEVMTx(parameter string) (*chainTypes.Transaction, error) {
//	tx, err := c.localCreateYCCEVMTx(parameter)
//	if err != nil {
//		return nil, err
//	}
//	tx.Sign(chainTypes.SECP256K1, pri)
//
//	return tx, nil
//}

func (c *CallContract) localCreateYCCEVMTx(parameter string) (*chainTypes.Transaction, error) {
	exec := c.ParaName + evmTypes.ExecutorName
	toAddr, err := chainAddress.GetExecAddress(exec, AddrType)
	if err != nil {
		return nil, err
	}

	_, packedParameter, err := evmAbi.Pack(parameter, c.Abi, false)
	if err != nil {
		return nil, err
	}

	action := evmTypes.EVMContractAction{
		Para:         packedParameter,
		ContractAddr: c.ContractAddr,
	}

	tx := &chainTypes.Transaction{Execer: []byte(exec), Payload: chainTypes.Encode(&action), Fee: 0, To: toAddr}
	//十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = yccChainId

	return tx, nil
}

// LocalTxGroup 批量生成本地签名后交易,多个交易会分成N个交易组形式
func (c *CallContract) LocalTxGroup(pristr string, parameters ...string) ([]*chainTypes.Transaction, error) {
	return c.localTxGroup(pristr, parameters...)
}

// LocalTxGroupFast 高性能
func (c *CallContract) LocalTxGroupFast(pristr string, parameters ...string) ([]*chainTypes.Transaction, error) {
	plen := len(parameters)
	ch := make(chan []*chainTypes.Transaction)
	gsize := plen * parasSize / cpuNum
	//gsize笔起一个协程
	var wg sync.WaitGroup
	// 向上取整
	g := plen / gsize
	if plen%gsize != 0 {
		g += 1
	}
	wg.Add(g)
	for i := 0; i < g; i++ {
		go func(n int) {
			res, _ := c.localTxGroup(pristr, parameters[n*gsize:Min(plen, (n+1)*gsize)]...)
			ch <- res
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	txs := []*chainTypes.Transaction{}
	for v := range ch {
		txs = append(txs, v...)
	}
	return txs, nil
}

// Min 返回小的
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (c *CallContract) localTxGroup(pristr string, parameters ...string) ([]*chainTypes.Transaction, error) {
	plen := len(parameters)
	//存放最终交易
	txs := make([]*chainTypes.Transaction, 0, (plen/groupSize)+1)
	//用来缓存临时交易
	tmp := make([]*chainTypes.Transaction, groupSize, groupSize)
	prikey := chainUtil.HexToPrivkey(pristr)
	i := 0
	//先循环生成交易
	for ; i < plen; i++ {
		tx, err := c.localCreateYCCEVMTx(parameters[i])
		if err != nil {
			return nil, err
		}
		tmp[i%groupSize] = tx
		if i%groupSize != groupSize-1 {
			continue
		}

		//需要分割,十倍手续费保证成功
		tx2, err := chainTypes.CreateTxGroup(tmp, 10*defaultFeeRate)
		if err != nil {
			return nil, err
		}
		for j := 0; j < groupSize; j++ {
			err := tx2.SignN(j, Ty, prikey)
			if err != nil {
				return nil, err
			}
		}
		txs = append(txs, tx2.Tx())
	}

	//将最后一组取出来,如果是单笔交易直接存进去，如果是多笔交易再进行一次交易组构造
	k := plen % groupSize
	switch k {
	//交易组正好装完,不做处理
	case 0:
	//最后一组里只有一笔
	case 1:
		tx := tmp[0]
		tx.Sign(Ty, prikey)
		txs = append(txs, tx)
	default:
		// 取最后一组的交易 左包右闭
		tx2, err := chainTypes.CreateTxGroup(tmp[0:k], 10*defaultFeeRate)
		if err != nil {
			return nil, err
		}
		for j := 0; j < k; j++ {
			err := tx2.SignN(j, Ty, prikey)
			if err != nil {
				return nil, err
			}
		}
		txs = append(txs, tx2.Tx())
	}

	return txs, nil
}
