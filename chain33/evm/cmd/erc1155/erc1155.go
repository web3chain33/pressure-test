package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/33cn/chain33/rpc/grpcclient"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
	"gitlab.33.cn/proof/backend-micro/pkg/tx/txpool"
	"gitlab.33.cn/proof/pressure-test/chain33/evm/call"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
	"time"
)

type Job struct {
	Parameter    string
	Privkey      string
	ContractAddr string
}

type Conf struct {
	PoolSize        int      `yaml:"PoolSize"`
	OperationType   int      `yaml:"OperationType"`
	Rate            int      `yaml:"Rate"`
	GroupSie        int      `yaml:"GroupSie"`
	GrpcTxNum       int      `yaml:"GrpcTxNum"`
	DeployerPrivkey string   `yaml:"DeployerPrivkey"`
	DeployerAddr    string   `yaml:"DeployerAddr"`
	RpcType         int      `yaml:"RpcType"` // 1:jsonrpc  2:grpc
	Chain           []string `yaml:"Chain"`
	ParaName        []string `yaml:"ParaName"`
	ContractAddr    []string `yaml:"ContractAddr"`
}

type Addr struct {
	Address string
	PrivKey string
}

var AddressList []Addr

var (
	//configFile  = flag.String("f", "etc/config.yaml", "the config file")
	//addressFile = flag.String("a", "etc/address.json", "the address file")

	configFile  = flag.String("f", "/Users/qiangu/gocode/src/gitlab.33.cn/proof/pressure-test/chain33/evm/etc/config.yaml", "the config file")
	addressFile = flag.String("a", "/Users/qiangu/gocode/src/gitlab.33.cn/proof/pressure-test/chain33/evm/etc/address.json", "the address file")
)

func main() {
	fmt.Println("开始测试")
	flag.Parse()

	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Println("ioutil.ReadFile err: ", err)
	}

	var c Conf
	err = yaml.Unmarshal(content, &c)
	if err != nil {
		fmt.Println("读取配置失败 err:", err)
		return
	}

	err = InitAddress(*addressFile)
	if err != nil {
		fmt.Println("初始化地址失败", err)
		return
	}

	if c.RpcType == 2 {

		grpcJobChain1 := make(chan *chainTypes.Transaction, 5000)
		grpcJobChain2 := make(chan *chainTypes.Transaction, 5000)
		grpcJobChain3 := make(chan *chainTypes.Transaction, 5000)
		grpcJobChain4 := make(chan *chainTypes.Transaction, 5000)

		cha := []chan *chainTypes.Transaction{grpcJobChain1, grpcJobChain2, grpcJobChain3, grpcJobChain4}
		InitGrpcJobChain(cha, c.ContractAddr[0], c.ParaName[0], c.DeployerPrivkey, c.OperationType, c.Rate)
		time.Sleep(3 * time.Second)

		for i := 0; i < len(cha); i++ {
			go Send(c.PoolSize, c.Chain[0], cha[i])
		}

		time.Sleep(10 * time.Minute)
		return
	}

	if c.RpcType == 3 {
		start := time.Now()

		nftId := 0
		r := c.Rate/len(c.ContractAddr) + 1
		job3 := make([][][]*chainTypes.Transaction, 0, len(c.ContractAddr))
		for k := 0; k < len(c.ContractAddr); k++ {
			jobLists := make([][]*chainTypes.Transaction, 0, r)
			for i := 0; i < r; i++ {
				l := make([]*chainTypes.Transaction, 0, len(AddressList))
				jobLists = append(jobLists, l)
			}

			id := InitGrpcJobList(nftId, jobLists, c.ContractAddr[k], c.ParaName[k], c.DeployerPrivkey, c.OperationType, r, c.GroupSie)
			nftId = id
			job3 = append(job3, jobLists)
		}

		stop := time.Now()
		fmt.Println("交易构造完毕，开始发送, 构造开始时间: ", start.String(), "结束时间: ", stop.String(), "耗时：", stop.Unix()-start.Unix())
		for h := 0; h < len(job3); h++ {
			for j := 0; j < len(job3[h]); j++ {
				go SendList(c.Chain[h], job3[h][j], c.GrpcTxNum)
			}
		}

		time.Sleep(20 * time.Minute)
		return
	}

	if c.RpcType == 4 {
		start := time.Now()

		contractAddrLen := len(c.ContractAddr)

		nftId := 0
		r := c.Rate/len(c.ContractAddr) + 1
		job3 := make([][][]*chainTypes.Transaction, 0, len(c.ContractAddr))

		poolSize := 6
		wg := &sync.WaitGroup{}
		wg.Add(poolSize * len(c.ContractAddr))

		groupChains := make([]chan *TxGroupParams, 0, contractAddrLen)
		resultChains := make([]chan *chainTypes.Transaction, 0, contractAddrLen)

		job3Len := 0
		for k := 0; k < len(c.ContractAddr); k++ {
			jobLists := make([][]*chainTypes.Transaction, 0, r)
			for i := 0; i < r; i++ {
				l := make([]*chainTypes.Transaction, 0, len(AddressList))
				jobLists = append(jobLists, l)
				job3Len++
			}
			job3 = append(job3, jobLists)

			groupChain := make(chan *TxGroupParams, 1000)
			groupChains = append(groupChains, groupChain)
			resultChain := make(chan *chainTypes.Transaction, 1000)
			resultChains = append(resultChains, resultChain)

			go PollCreateTxGroup(poolSize, c.ContractAddr[k], c.ParaName[k], c.DeployerPrivkey, groupChains[k], resultChains[k], wg)
			go ChainToJobList(resultChains[k], job3[k])
			go InitGrpcTxGroupChain(nftId, c.OperationType, r, c.GroupSie, groupChains[k])

			nftId += len(AddressList) * r

		}
		wg.Wait()
		time.Sleep(1 * time.Second)

		createStop := time.Now()
		fmt.Println("交易构造完毕，开始发送, 构造开始时间: ", start.String(), "结束时间: ", createStop.String(), "耗时：", createStop.Unix()-start.Unix())

		wg.Add(job3Len)
		for h := 0; h < len(job3); h++ {
			for j := 0; j < len(job3[h]); j++ {
				go SendListWaitGroup(c.Chain[h], job3[h][j], c.GrpcTxNum, wg)
			}
		}

		wg.Wait()
		time.Sleep(1 * time.Second)
		sendStop := time.Now()
		fmt.Println("交易发送完毕，发送结束时间", sendStop.String(), "耗时：", sendStop.Unix() - createStop.Unix())
		return
	}

	jobChain := make(chan *Job, 5000)
	err = InitJobChain(jobChain, c.ContractAddr[0], c.DeployerPrivkey, c.OperationType, c.Rate)
	if err != nil {
		fmt.Println("main InitJobChain err", err)
		return
	}

	time.Sleep(3 * time.Second)

	client := call.NewJsonClient(c.Chain[0], c.ParaName[0],
		c.ContractAddr[0], abi)

	resultChain := make(chan int, 5000)

	start := time.Now().Unix()
	defer func(start int64) {
		stop := time.Now().Unix()
		fmt.Printf("开始发送：%v , 结束发送：%v , 耗时: %v s \n", start, stop, stop-start)
	}(start)
	time.Sleep(1 * time.Second)
	CreatePool(c.PoolSize, jobChain, resultChain, client, c.DeployerAddr, c.DeployerPrivkey)

	resultNum := 0
	successNum := 0
	failNum := 0
	defer func() {
		fmt.Println("发行成功：", successNum, "发行失败：", failNum)
	}()

	for status := range resultChain {
		fmt.Println("resultChain status", status)
		if status == call.Success {
			successNum++
		} else {
			failNum++
		}

		resultNum++
		tokenRate := 1
		if c.OperationType == 1 || c.OperationType == 3 {
			tokenRate = c.Rate
		}
		if resultNum >= len(AddressList)*tokenRate {
			return
		}
	}
}

func InitAddress(addressFile string) error {
	content, err := ioutil.ReadFile(addressFile)
	if err != nil {
		return err
	}

	AddressList = make([]Addr, 0, 10000)
	err = json.Unmarshal(content, &AddressList)
	if err != nil {
		return err
	}
	return nil
}

func InitJobChain(jobChain chan *Job, contractAddr, deployerPrivkey string, operationType, rate int) error {
	go func(jobChain chan *Job, contractAddr, privkey string) {
		nftId := 0
		if operationType == 1 {
			for i := 0; i < len(AddressList); i++ {
				for j := 0; j < rate; j++ {
					nftId++
					job := &Job{
						Parameter:    fmt.Sprintf("mint(%q, %v)", AddressList[i].Address, nftId),
						Privkey:      privkey,
						ContractAddr: contractAddr,
					}
					jobChain <- job
				}
			}
		} else if operationType == 2 {
			for i := 0; i < len(AddressList); i++ {
				ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
				nftId += 5

				idsByte, _ := json.Marshal(ids)
				job := &Job{
					Parameter:    fmt.Sprintf("batchMint(%q, %v)", AddressList[i].Address, string(idsByte)),
					Privkey:      privkey,
					ContractAddr: contractAddr,
				}
				jobChain <- job
			}
		} else if operationType == 3 {
			addrLen := len(AddressList)
			for i := 0; i < addrLen; i++ {
				for j := 0; j < rate; j++ {
					nftId++
					job := &Job{
						Parameter:    fmt.Sprintf("transfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, nftId),
						Privkey:      AddressList[i].PrivKey,
						ContractAddr: contractAddr,
					}
					jobChain <- job
				}
			}
		} else if operationType == 4 {
			addrLen := len(AddressList)
			for i := 0; i < addrLen; i++ {
				ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
				nftId += 5

				idsByte, _ := json.Marshal(ids)
				job := &Job{
					Parameter:    fmt.Sprintf("batchTransfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, string(idsByte)),
					Privkey:      AddressList[i].PrivKey,
					ContractAddr: contractAddr,
				}
				jobChain <- job
			}
		}
	}(jobChain, contractAddr, deployerPrivkey)

	return nil
}

func CreatePool(num int, jobChan chan *Job, resultChain chan int, cli *call.JsonClient, deployerAddr, deployerPrivkey string) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, cli *call.JsonClient) {
			// cli := NewClient(c.Chain.Endpoint, c.Chain.ParaName, c.Contract.Addr, c.Contract.Abi)
			// 执行运算
			for job := range jobChan {
				status, _ := cli.SendContractGroup(job.Parameter, job.Privkey, deployerAddr, deployerPrivkey)
				resultChain <- status
			}
		}(jobChan, cli)
	}
}

func InitGrpcJobChain(grpcJobChain []chan *chainTypes.Transaction, contractAddr, paraName, deployerPrivkey string, operationType, rate int) {
	go func(grpcJobChain []chan *chainTypes.Transaction, contractAddr, paraName, deployerPrivkey string, operationType, rate int) {
		nftId := 0
		c := &call.CallContract{
			ContractAddr: contractAddr,
			ParaName:     paraName,
			Abi:          abi,
			DeployerPri:  chainUtil.HexToPrivkey(deployerPrivkey),
		}

		if operationType == 1 {
			for i := 0; i < len(AddressList); i++ {
				for j := 0; j < rate; j++ {
					nftId++
					tx, err := c.LocalCreateYCCEVMTx(fmt.Sprintf("mint(%q, %v)", AddressList[i].Address, nftId))
					if err != nil {
						fmt.Println("c.LocalCreateYCCEVMTx ,err: ", err)
						continue
					}
					y := nftId % len(grpcJobChain)
					grpcJobChain[y] <- tx
				}
			}
		} else if operationType == 2 {
			for i := 0; i < len(AddressList); i++ {
				ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
				nftId += 5

				idsByte, _ := json.Marshal(ids)
				tx, err := c.LocalCreateYCCEVMTx(fmt.Sprintf("batchMint(%q, %v)", AddressList[i].Address, string(idsByte)))
				if err != nil {
					fmt.Println("c.LocalCreateYCCEVMTx ,err: ", err)
					continue
				}
				y := nftId % len(grpcJobChain)
				grpcJobChain[y] <- tx
			}
		} else if operationType == 3 {
			addrLen := len(AddressList)
			for i := 0; i < addrLen; i++ {
				for j := 0; j < rate; j++ {
					nftId++
					tx, err := c.LocalCreateYCCEVMTx(fmt.Sprintf("transfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, nftId))
					if err != nil {
						fmt.Println("c.LocalCreateYCCEVMTx ,err: ", err)
						continue
					}
					y := nftId % len(grpcJobChain)
					grpcJobChain[y] <- tx
				}
			}
		} else if operationType == 4 {
			addrLen := len(AddressList)
			for i := 0; i < addrLen; i++ {
				ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
				nftId += 5

				idsByte, _ := json.Marshal(ids)

				tx, err := c.LocalCreateYCCEVMTx(fmt.Sprintf("batchTransfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, string(idsByte)))
				if err != nil {
					fmt.Println("c.LocalCreateYCCEVMTx ,err: ", err)
					continue
				}
				y := nftId % len(grpcJobChain)
				grpcJobChain[y] <- tx

			}
		}

	}(grpcJobChain, contractAddr, paraName, deployerPrivkey, operationType, rate)

}

type TxGroupParams struct {
	Params   []string
	Privkeys []string
}

func PollCreateTxGroup(poolSize int, contractAddr, paraName, deployerPrivkey string, groupChain chan *TxGroupParams, resultChain chan *chainTypes.Transaction, wg *sync.WaitGroup) {
	c := &call.CallContract{
		ContractAddr: contractAddr,
		ParaName:     paraName,
		Abi:          abi,
		DeployerPri:  chainUtil.HexToPrivkey(deployerPrivkey),
	}

	for i := 0; i < poolSize; i++ {
		go func(c *call.CallContract, groupChain chan *TxGroupParams, wg *sync.WaitGroup) {
			for param := range groupChain {
				tx, err := c.LocalCreateYCCEVMGroupTx(param.Params, param.Privkeys)
				if err != nil {
					fmt.Println("c.LocalCreateYCCEVMGroupTx ,err: ", err)
					continue
				}
				resultChain <- tx
			}
			wg.Done()
		}(c, groupChain, wg)
	}
}

func ChainToJobList(resultChain chan *chainTypes.Transaction, jobLists [][]*chainTypes.Transaction) {
	groupCount := 0
	for tx := range resultChain {
		groupCount++
		y := groupCount % len(jobLists)
		jobLists[y] = append(jobLists[y], tx)
	}
}

func InitGrpcTxGroupChain(nftId int, operationType, rate, groupSize int, groupChain chan *TxGroupParams) {

	txCount := 0
	params := make([]string, 0, groupSize)
	privkeys := make([]string, 0, groupSize)

	if operationType == 1 {
		for i := 0; i < len(AddressList); i++ {
			for j := 0; j < rate; j++ {
				nftId++
				txCount++
				params = append(params, fmt.Sprintf("mint(%q, %v)", AddressList[i].Address, nftId))

				if txCount >= groupSize {
					param := &TxGroupParams{
						Params:   params,
						Privkeys: privkeys,
					}
					groupChain <- param

					txCount = 0
					params = make([]string, 0, groupSize)
				}

			}
		}
		close(groupChain)
	}
}

func InitGrpcJobList(nftId int, jobLists [][]*chainTypes.Transaction, contractAddr, paraName, deployerPrivkey string, operationType, rate, groupSize int) int {
	c := &call.CallContract{
		ContractAddr: contractAddr,
		ParaName:     paraName,
		Abi:          abi,
		DeployerPri:  chainUtil.HexToPrivkey(deployerPrivkey),
	}

	txCount := 0
	params := make([]string, 0, groupSize)
	privkeys := make([]string, 0, groupSize)
	groupCount := 0
	if operationType == 1 {
		for i := 0; i < len(AddressList); i++ {
			for j := 0; j < rate; j++ {
				nftId++
				txCount++
				params = append(params, fmt.Sprintf("mint(%q, %v)", AddressList[i].Address, nftId))

				if txCount >= groupSize {
					tx, err := c.LocalCreateYCCEVMGroupTx(params, privkeys)
					if err != nil {
						fmt.Println("c.LocalCreateYCCEVMGroupTx ,err: ", err)
						continue
					}
					groupCount++
					y := groupCount % len(jobLists)
					jobLists[y] = append(jobLists[y], tx)

					txCount = 0
					params = make([]string, 0, groupSize)
				}
			}
		}
	} else if operationType == 2 {
		for i := 0; i < len(AddressList); i++ {
			ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
			nftId += 5

			idsByte, _ := json.Marshal(ids)

			txCount++
			params = append(params, fmt.Sprintf("batchMint(%q, %v)", AddressList[i].Address, string(idsByte)))

			if txCount >= groupSize {
				tx, err := c.LocalCreateYCCEVMGroupTx(params, privkeys)
				if err != nil {
					fmt.Println("c.LocalCreateYCCEVMGroupTx ,err: ", err)
					continue
				}
				groupCount++
				y := groupCount % len(jobLists)
				jobLists[y] = append(jobLists[y], tx)

				txCount = 0
				params = make([]string, 0, groupSize)
			}
		}
	} else if operationType == 3 {
		addrLen := len(AddressList)
		for i := 0; i < addrLen; i++ {
			for j := 0; j < rate; j++ {
				nftId++
				txCount++
				params = append(params, fmt.Sprintf("transfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, nftId))
				privkeys = append(privkeys, AddressList[i].PrivKey)

				if txCount >= groupSize {
					tx, err := c.LocalCreateYCCEVMGroupTx(params, privkeys)
					if err != nil {
						fmt.Println("c.LocalCreateYCCEVMGroupTx ,err: ", err)
						continue
					}
					groupCount++
					y := groupCount % len(jobLists)
					jobLists[y] = append(jobLists[y], tx)

					txCount = 0
					params = make([]string, 0, groupSize)
					privkeys = make([]string, 0, groupSize)
				}
			}
		}
	} else if operationType == 4 {
		addrLen := len(AddressList)
		for i := 0; i < addrLen; i++ {
			ids := []int{nftId + 1, nftId + 2, nftId + 3, nftId + 4, nftId + 5}
			nftId += 5

			idsByte, _ := json.Marshal(ids)

			txCount++
			params = append(params, fmt.Sprintf("batchTransfer(%q, %q, %v)", AddressList[i].Address, AddressList[addrLen-1-i].Address, string(idsByte)))
			privkeys = append(privkeys, AddressList[i].PrivKey)

			if txCount >= groupSize {
				tx, err := c.LocalCreateYCCEVMGroupTx(params, privkeys)
				if err != nil {
					fmt.Println("c.LocalCreateYCCEVMGroupTx ,err: ", err)
					continue
				}
				groupCount++
				y := groupCount % len(jobLists)
				jobLists[y] = append(jobLists[y], tx)

				txCount = 0
				params = make([]string, 0, groupSize)
				privkeys = make([]string, 0, groupSize)
			}

		}
	}
	return nftId
}

func SendListWaitGroup(endpoint string, jobList []*chainTypes.Transaction, grpcTxNum int, wg *sync.WaitGroup) {
	maxMsgSize := 20 * 1024 * 1024 //最大传输数据 最大区块大小
	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize))

	conn, err := grpc.Dial(grpcclient.NewMultipleURL(endpoint), grpc.WithInsecure(), diaOpt)
	if err != nil {
		fmt.Println("grpcclient.NewMultipleURL err:", err)
		return
	}
	client := chainTypes.NewChain33Client(conn)
	for i := 0; i < len(jobList); i += grpcTxNum {
		replys, err := client.SendTransactions(context.Background(), &chainTypes.Transactions{Txs: jobList[i : i+grpcTxNum]})

		if err != nil {
			fmt.Println("SendTransaction err:", err)
		}
		fmt.Println("SendTransactions replys, isOK: ", replys.ReplyList[0].IsOk)
	}
	wg.Done()
}

func SendList(endpoint string, jobList []*chainTypes.Transaction, grpcTxNum int) {
	maxMsgSize := 20 * 1024 * 1024 //最大传输数据 最大区块大小
	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize))

	conn, err := grpc.Dial(grpcclient.NewMultipleURL(endpoint), grpc.WithInsecure(), diaOpt)
	if err != nil {
		fmt.Println("grpcclient.NewMultipleURL err:", err)
		return
	}
	client := chainTypes.NewChain33Client(conn)
	for i := 0; i < len(jobList); i += grpcTxNum {
		replys, err := client.SendTransactions(context.Background(), &chainTypes.Transactions{Txs: jobList[i : i+grpcTxNum]})

		if err != nil {
			fmt.Println("SendTransaction err:", err)
		}
		fmt.Println("SendTransactions replys, isOK: ", replys.ReplyList[0].IsOk)
	}
}

func Send(poolSize int, endpoint string, grpcJobChain chan *chainTypes.Transaction) {
	grpc := txpool.CreateTxPool(&txpool.Cfg{
		Chs: txpool.Chs{
			Size: poolSize,
			Len:  100, //OpenRetry=true 会堵塞当前ch 这个len只在 OpenRetry=false有效
		},
		GrpcAddrs: []string{endpoint},
	})

	id := 0
	for tx := range grpcJobChain {
		id++
		grpc.SendTx(&txpool.TxMsg{
			Id:        int64(id),
			Time:      time.Now().Unix(),
			Tx:        tx,
			OpenRetry: true,
		})
	}

}

const abi = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"},{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"}],"name":"URI","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"accounts","type":"address[]"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"balanceOfBatch","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"batchMint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"batchTransfer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getSuccessNum","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeBatchTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"successNum","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"transfer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
