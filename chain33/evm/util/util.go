package util

import (
	chainAddress "github.com/33cn/chain33/common/address"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	evmAbi "github.com/33cn/plugin/plugin/dapp/evm/executor/abi"
	evmtypes "github.com/33cn/plugin/plugin/dapp/evm/types"
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
var parasLen = 4

// CallContract 成功部署后的合约
type CallContract struct {
	ContractAddr string
	ParaName     string
	Abi          string
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
	tx.Sign(chainTypes.SECP256K1, prikey)

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
	gsize := plen * parasLen / cpuNum
	//gsize笔起一个携程
	var wg sync.WaitGroup
	g := plen/gsize + 1
	if plen%gsize == 0 {
		g -= 1
	}
	wg.Add(g)
	for i := 0; i < g; i++ {
		go func(n int) {
			res, _ := c.localTxGroup(pristr, parameters[n*gsize:min(plen, (n+1)*gsize)]...)
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

func min(a, b int) int {
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
			err := tx2.SignN(j, chainTypes.SECP256K1, prikey)
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
		tx.Sign(chainTypes.SECP256K1, prikey)
		txs = append(txs, tx)
	default:
		// 取最后一组的交易 左包右闭
		tx2, err := chainTypes.CreateTxGroup(tmp[0:k], 10*defaultFeeRate)
		if err != nil {
			return nil, err
		}
		for j := 0; j < k; j++ {
			err := tx2.SignN(j, chainTypes.SECP256K1, prikey)
			if err != nil {
				return nil, err
			}
		}
		txs = append(txs, tx2.Tx())
	}

	return txs, nil
}
