package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gitlab.33.cn/proof/pressure-test/chain33/evm/call"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	DeployerPrivkey string   `yaml:"DeployerPrivkey"`
	DeployerAddr    string   `yaml:"DeployerAddr"`
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
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	addressFile = flag.String("a", "etc/address.json", "the address file")
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

const abi = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"},{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"}],"name":"URI","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"accounts","type":"address[]"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"balanceOfBatch","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"batchMint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"}],"name":"batchTransfer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getSuccessNum","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256[]","name":"ids","type":"uint256[]"},{"internalType":"uint256[]","name":"amounts","type":"uint256[]"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeBatchTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"successNum","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"id","type":"uint256"}],"name":"transfer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"uri","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`


