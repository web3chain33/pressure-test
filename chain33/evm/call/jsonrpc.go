package call

import (
	"gitlab.33.cn/proof/backend-micro/pkg/evm"
	"fmt"
	"strconv"

	chainCommon "github.com/33cn/chain33/common"
	chainTypes "github.com/33cn/chain33/types"
	chainUtil "github.com/33cn/chain33/util"
)

const (
	Success = 1
	Fail    = 2
)

type Contract struct {
	endpoint string
	paraName string
	addr     string
	abi      string
}

func NewtContract(endpoint, paraName, addr, abi string) Contract {
	return Contract{
		endpoint: endpoint,
		paraName: paraName,
		addr:     addr,
		abi:      abi,
	}
}

type JsonClient struct {
	Endpoint string
	ParaName string
	Addr     string
	Abi      string
	Cli      *evm.Client
}

func NewJsonClient(endpoint, paraName, addr, abi string ) *JsonClient {
	client := JsonClient{
		Endpoint: endpoint,
		ParaName:paraName,
		Addr:addr,
		Abi: abi,
		Cli: evm.NewClient(paraName, endpoint),
	}
	return &client
}

func (c *JsonClient) SendContract(parameter, privkey, deployerAddr, deployerPrivkey string) (status int, hash string) {
	// 1)
	// 创建操作合约交易
	resp1, err := c.Cli.CreateCallTx(&evm.CreateCallTxReq{
		Abi:          c.Abi,
		Parameter:    parameter,
		Expire:       "300s",
		ContractAddr: c.Addr,
	})
	if err != nil {
		return Fail, ""
	}

	fmt.Printf("resp %+v\n", resp1)

	// 2)
	// 估算部署交易或者调用交易需要的 gas
	resp2, err := c.Cli.EstimateGas(&evm.EstimateGasReq{
		Tx:   resp1.Result,
		From: deployerAddr,
	})
	if err != nil {
		return Fail, ""
	}

	fee, _ := strconv.Atoi(resp2.Gas)
	// 3)
	// 签名

	// fee := 1000000 // 固定0.01 bty
	resp3, err := c.Cli.SignRawTx(&evm.SignRawTxReq{
		Privkey: privkey,
		TxHex:   resp1.Result,
		Expire:  "300s",
		// Index:     0,
		// Token:     "",
		Fee: int64(fee),
		// NewToAddr: "",
	})
	if err != nil {
		return Fail, ""
	}

	// 4)
	// 发送交易
	transaction, err := c.Cli.SendTransaction(&evm.SendTransactionReq{
		Data: resp3.Result,
	})
	if err != nil {
		return Fail, ""
	}

	return Success, transaction.Result
}

func (c *JsonClient) SendContractGroup(parameter, privkey, deployerAddr, deployerPrivkey string) (status int, hash string) {
	// 1)
	// 创建操作合约交易
	resp1, err := c.Cli.CreateCallTx(&evm.CreateCallTxReq{
		Abi:          c.Abi,
		Parameter:    parameter,
		Expire:       "300s",
		ContractAddr: c.Addr,
	})
	if err != nil {
		fmt.Printf("func CreateDeployTxerr error: %s", err.Error())
		return Fail, ""
	}

	t1, err := evm.GetTxFromHex(resp1.Result)
	if err != nil {
		return Fail, ""
	}

	// 2)
	// 估算部署交易或者调用交易需要的 gas
	resp2, err := c.Cli.EstimateGas(&evm.EstimateGasReq{
		Tx:   resp1.Result,
		From: deployerAddr,
	})
	if err != nil {
		fmt.Printf("func EstimateGasReq error: %s\n", err.Error())
		return Fail, ""
	}

	gas, _ := strconv.Atoi(resp2.Gas)

	t1.Fee = int64(gas * 11)

	txGroup, err := c.Cli.CreateWithholdTxGroup(c.ParaName, t1)
	if err != nil {
		fmt.Printf("func EstimateGasReq error: %s\n", err.Error())
		return Fail, ""
	}

	err = txGroup.SignN(0, 1,
		chainUtil.HexToPrivkey(deployerPrivkey))
	if err != nil {
		return Fail, ""
	}

	err = txGroup.SignN(1, 1, chainUtil.HexToPrivkey(privkey))
	if err != nil {
		return Fail, ""
	}
	tg := txGroup.Tx()

	resp3, err := c.Cli.SendTransaction(&evm.SendTransactionReq{
		Data: chainCommon.ToHex(chainTypes.Encode(tg)),
	})
	if err != nil {
		fmt.Printf("func SignRawTx error: %s\n", err.Error())
		return Fail, ""
	}

	return Success, resp3.Result
}
