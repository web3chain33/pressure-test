package util

import (
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	"math/rand"
	"sync"
	"time"
)

type Proof struct {
	ParaName string
	TxNum    int
	Note     string
}

// 创建存证交易
func (t *Proof) localProofTx(prikey crypto.PrivKey) (*chainTypes.Transaction, error) {
	execer := t.ParaName + "none"
	execAddr := address.ExecAddress(execer)

	tx := &chainTypes.Transaction{Execer: []byte(execer), Payload: []byte("proot"), To: execAddr}

	//十倍手续费保证成功
	tx.Fee, _ = tx.GetRealFee(10 * defaultFeeRate)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	tx.ChainID = yccChainId
	tx.Sign(Ty, prikey)

	return tx, nil
}

// LocalProofFast 高性能创建转账
func (t *Proof) LocalProofFast(pristr string) ([]*chainTypes.Transaction, error) {
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
				res, _ := t.localProofTx(prikey)
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
