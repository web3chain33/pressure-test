package util

import (
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	"math/rand"
	"sync"
	"time"
)

type Transfer struct {
	ParaName string
	TxNum    int
	ToAddr   string
	Amount   int64
}

// 创建转账交易
func (t *Transfer) localTransferTx(prikey crypto.PrivKey) (*chainTypes.Transaction, error) {
	execer := t.ParaName + "coins"
	execAddr := address.ExecAddress(execer)
	transfer := &cty.CoinsAction{}
	v := &cty.CoinsAction_Transfer{Transfer: &chainTypes.AssetsTransfer{Amount: t.Amount, To: t.ToAddr}}
	transfer.Value = v
	transfer.Ty = cty.CoinsActionTransfer

	tx := &chainTypes.Transaction{Execer: []byte(execer), Payload: chainTypes.Encode(transfer), To: execAddr}

	//十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = yccChainId
	tx.Sign(Ty, prikey)

	return tx, nil
}

// LocalTransferFast 高性能创建转账
func (t *Transfer) LocalTransferFast(pristr string) ([]*chainTypes.Transaction, error) {
	prikey := chainUtil.HexToPrivkey(pristr)
	ch := make(chan *chainTypes.Transaction)
	gsize := t.TxNum * parasSize / cpuNum
	//gsize笔起一个协程
	var wg sync.WaitGroup
	// 向上取整
	g := t.TxNum / gsize
	if t.TxNum%gsize != 0 {
		g += 1
	}
	wg.Add(g)
	for i := 0; i < g; i++ {
		go func(n int) {
			for j := 0; j < Min(gsize, t.TxNum-n*gsize); j++ {
				res, _ := t.localTransferTx(prikey)
				ch <- res
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	txs := []*chainTypes.Transaction{}
	for v := range ch {
		txs = append(txs, v)
	}

	return txs, nil
}
