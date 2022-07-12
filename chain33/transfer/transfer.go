package main

import (
	"context"
	"fmt"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/rpc/grpcclient"
	chainTypes "github.com/33cn/chain33/types"
	"github.com/web3chain33/pressure-test/chain33/util"
	"google.golang.org/grpc"
	"runtime"
	"sync"
	"time"
)

var cpuNum = runtime.NumCPU()
var Cfg *Config

func main() {
	runtime.GOMAXPROCS(cpuNum)
	Setup()

	paras := []*Para{}
	for _, v := range Cfg.Paras {
		paras = append(paras, v)
	}
	parasLen := len(paras)
	util.SetParasLen(parasLen)
	util.InitTy(Cfg.ChainType)
	var wg sync.WaitGroup
	wg.Add(parasLen)
	for _, v := range paras {
		go func(p *Para) {
			p.Run()
			wg.Done()
		}(v)
	}
	wg.Wait()

}

type Para struct {
	Name         string `json:",default=user.p.para1."`
	JrpcEndpoint string `json:",default=http://172.16.103.233:8911"` // jrpc端口
	GrpcEndpoint string `json:",default=172.16.103.233:8912"`        //grpc端口
}

func (p *Para) Run() {
	//grpc1
	maxMsgSize := 4 * 1024 * 1024 //最大传输数据 最大区块大小 与服务端一致
	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize))

	conn, err := grpc.Dial(grpcclient.NewMultipleURL(p.GrpcEndpoint),
		grpc.WithInsecure(), diaOpt)

	if err != nil {
		panic(err)
	}
	client := chainTypes.NewChain33Client(conn)

	// 生成合约结构体
	t := &util.Transfer{
		ParaName: p.Name,
		TxNum:    Cfg.Txnum,
		ToAddr:   Cfg.ToAddr,
		Amount:   Cfg.Amount,
	}

	log.Info(fmt.Sprintf("平行链%v开始构造%v笔交易", p.Name, Cfg.Txnum))
	time1 := time.Now().Unix()

	//构造普通交易
	txs, err := t.LocalTransferFast(Cfg.UserPrivateKey)
	if err != nil {
		panic(err)
	}
	time2 := time.Now().Unix()
	log.Info(fmt.Sprintf("平行链%v构造用时:%vs", p.Name, time2-time1))
	log.Info(fmt.Sprintf("平行链%v成功构造交易:%v笔", p.Name, len(txs)))

	//一笔交易是一个交易组，一个交易组目前的大小是接近20*0.5kb
	// grpc默认限制4mb接收大小，可以弄400个交易组一次性发送
	var wg sync.WaitGroup
	time1 = time.Now().Unix()
	g := len(txs) / Cfg.GrpcTxNum
	if len(txs)%Cfg.GrpcTxNum != 0 {
		g += 1
	}

	wg.Add(g)
	log.Info(fmt.Sprintf("平行链%v开始发送，每次发送:%v笔,共发送%v次", p.Name, Cfg.GrpcTxNum, g))
	for i := 0; i < g; i++ {
		go func(n int) {
			_, err := client.SendTransactions(context.Background(),
				&chainTypes.Transactions{Txs: txs[n*Cfg.GrpcTxNum : util.Min((n+1)*Cfg.GrpcTxNum, len(txs))]})
			if err != nil {
				panic(err)
			}
			wg.Done()
			log.Info(fmt.Sprintf("平行链%v第%v次发送完毕", p.Name, n+1))
		}(i)
	}

	wg.Wait()
	time2 = time.Now().Unix()
	log.Info(fmt.Sprintf("平行链%v发送用时:%vs", p.Name, time2-time1))
}
