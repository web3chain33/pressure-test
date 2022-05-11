package call

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	l "github.com/33cn/chain33/common/log/log15"
	"github.com/chendehai/pressure-test/eth/solidity/goods"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var log = l.New("module", "call")

var keys = []string{
	`{"address":"273d6deaddd423fa9a5fcbe44a0c303e8d2a65d3","crypto":{"cipher":"aes-128-ctr","ciphertext":"8687713a07bd7cd14dc4bc9c368aadfa116a0a4c8468ba293382ad79bed68288","cipherparams":{"iv":"210e3c0fa46872980cc5f5ecfc9119da"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"90eb440b97c2ade69977028175314df235e291feba474ca262ef8b637a62c3e7"},"mac":"c5b53d3307a279ec50f9e3075bdc93d6c6cd91d300039028a6e9d488e16b0043"},"id":"c5b111d6-0056-4f55-98e5-8ba3879f4da5","version":3}`,
	`{"address":"444934ea925050f8c46d4a5739cf2deadbb2b29d","crypto":{"cipher":"aes-128-ctr","ciphertext":"e931942875b17b8d291e09c3e3cbc2ada5298c93211f59ad37513f451430762c","cipherparams":{"iv":"f2d0e12411372df25f49d02542b87200"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"13eb02896e9fbf842aa7f44c7e54f8e9fe93492c845ebd0ece3e38ff17a4fc9d"},"mac":"1f68478df310da79328f874d2a4369e871b3587cd9e54ff62af5eb818c7cf41b"},"id":"5f6706c3-e0d5-45a7-95bf-9e9d61a0b6c4","version":3}`,
	`{"address":"6fe3ab57e4b490b9cf6db8763d8e2948ead8102d","crypto":{"cipher":"aes-128-ctr","ciphertext":"8c41e34f17f1c1bcbfa0d1e851594333544ac56a213e3115d6b16e4726ec9907","cipherparams":{"iv":"0acc085629eb6d856277d41afa5e7a03"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"674d4d33158868736ac2f5be3e454c12a30c9176bc584071a8309af8f8305202"},"mac":"8a25e975cc1810c0721a7ca3712c359cd5a3ff617e5a53334edac1172f519926"},"id":"32f2230c-df2a-4264-a558-e0c2fd3f37a6","version":3}`,
	`{"address":"bd0282fba7edcfd57ac14093409db9228554cdda","crypto":{"cipher":"aes-128-ctr","ciphertext":"cc4bd269499cb9b0d2ed592f7ac71018a0d9138a119110694f5fd5f787e0c0c7","cipherparams":{"iv":"06893c836365920df052927f61eaa6f5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"464fb8d1aaf6b3829f1cf4b3474d51bf2a52af27e1070fbd361256d919b30f48"},"mac":"d49275d57f47b01b8fa187944a7f33e326a7b8748bb462f02a0e9d047ee549e5"},"id":"94af218a-60d7-4d90-a4c6-1576aa3a9182","version":3}`,
	`{"address":"93f730797216a8d22483eed0041b59bdd0d719ce","crypto":{"cipher":"aes-128-ctr","ciphertext":"c2fb57609347b55abdeaaf0ffe809e13bdc2d0773cecf55e190608521a2b4331","cipherparams":{"iv":"2cb522a2ea1da9302b52cd41a88754ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"65132109876bd640d67f15a916414b7a8e9b70d39bc12cecb01db0792a2a6513"},"mac":"da8bc2958639cc45909539e56eb0c099ab33aa2c55d2a1d9ab6fcc56f697f58a"},"id":"14c04fcc-b242-4f8f-a56e-b5afdf987808","version":3}`,
	`{"address":"89024430ac684ef2142aeafdf07ede9e0702e772","crypto":{"cipher":"aes-128-ctr","ciphertext":"e6df69145d8515fb8a4f90acb9bb3d642480a47d6603003269543678801d24d1","cipherparams":{"iv":"cafc780c0fe496ea3eb191ce2149e1cb"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"95e9b12d873e1db369ca7c832e170ca949462a3571e2a6e5301417f050a240de"},"mac":"aea360b49d8bb0b0d58efa9880e638feba23fbd45286efb4f04df378199c6305"},"id":"973ca201-6102-4de5-b433-4d1b3783db8e","version":3}`,
	`{"address":"43df32c7a899efd337824497f69e1d0b194b25af","crypto":{"cipher":"aes-128-ctr","ciphertext":"d41315d43d7a9b042bb30c37e527a811fd6d20220bb94e11f69643f5a56916ab","cipherparams":{"iv":"7141cf3afbdd350b9c95bff7f358529d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"eb937bb421ef61bdcc23508fa1d5defb18678f8b7d0284151ba708c951b9dfdd"},"mac":"10fa566f2c574530c26679400dfed6da92d41cabca3bfb88309a5ededd9b22db"},"id":"ed622f7f-e2c3-4fee-a274-74bbadd1d30c","version":3}`,
	`{"address":"d8477f81e038ddee3ca39747cfc581d17053b31b","crypto":{"cipher":"aes-128-ctr","ciphertext":"8fbcc9e29bb45760a08b537f58444180a258669fe20f20b698c388aee4a6fe83","cipherparams":{"iv":"7e0d553d89f9836a29f4423e1c32b49b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5015cd797bd5c91855fa639bedc474681e47c5da8857d21a906cf1b4f4c63d78"},"mac":"d8915780ece6b28200111bd4c3dcc93dafbeb12380f6f465158a014e7cbdc6e6"},"id":"9d14fdc5-d6b8-4b57-a379-38795c6f588d","version":3}`,
	`{"address":"179576b1355ad891a2e71215aca32fb145c6342f","crypto":{"cipher":"aes-128-ctr","ciphertext":"7270af5f323f1fce8ddf9306737fa5b5569dfe555b828af786b3fd058f31fd4c","cipherparams":{"iv":"e0782027b27e777c1e1e9d1f0549d6c5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"e35c702f98ea737131c523008b82614664a3ad9f4e1b8935578463d549f67328"},"mac":"e540e89bbfed3e245a944e301d3182ea0507d583c16422a06c86deb2e0bd806b"},"id":"de48c66a-b11a-43ff-81e8-66833e285cae","version":3}`,
	`{"address":"bb19556ffb699a26e87c69374579338fe79949ac","crypto":{"cipher":"aes-128-ctr","ciphertext":"735cb69d13646b609ff868d0492fa0530059ba6842e41ed1c8265acfc6e989c6","cipherparams":{"iv":"a071be41fed3fe518cc027dceb10fad5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8d866822e0c175f2202bb7a9209a5ad7e300e7915d28e8ade017b42bbf8cd92d"},"mac":"8d7df9ae5babfd972eb0ff4f20e760d197293948080faa53856afbb3ca6ddd0d"},"id":"6e134fcb-3293-430c-b56d-9ef6e5f779b5","version":3}`,
}

var AccountOpts = InitOpts(keys)

func InitOpts(keys []string) []*bind.TransactOpts {
	opts := make([]*bind.TransactOpts, 0, len(keys))
	for _, key := range keys {
		opt, err := bind.NewTransactorWithChainID(strings.NewReader(key), "fuzamei123456", big.NewInt(666))
		if err != nil {
			fmt.Printf("Failed to connect to the Ethereum client: %v", err)
			continue
		}
		opts = append(opts, opt)
	}
	return opts
}

var (
	Address1 = common.HexToAddress("0xcfdeedacf829ca9ea4b2c0012f72bd3832acc381")
	Address2 = common.HexToAddress("0d3a7b7323c40a8b47bce79a81c74d3347f7037c")
)

func PollSend(poolSize, nftStep, operationType int, chainAddr, contractAddr string, wg *sync.WaitGroup) {
	nftId := 0

	for i := 0; i < poolSize; i++ {
		switch operationType {
		case 1:
			y := i % len(AccountOpts)
			go goodsMint(nftId, nftId+nftStep, chainAddr, contractAddr, AccountOpts[y], wg)
		case 3:
			y := i % len(AccountOpts)
			go goodsTransfer(nftId, nftId+nftStep, chainAddr, contractAddr, AccountOpts[y], wg)
		default:
			fmt.Println("error operationType:", operationType)
		}
		nftId += nftStep
	}
}

func goodsMint(nftId, nftMax int, chainAddr, contractAddr string, opt *bind.TransactOpts, wg *sync.WaitGroup) {
	conn, err := ethclient.Dial(chainAddr)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	contract, err := goods.NewGoods(common.HexToAddress(contractAddr), conn)
	if err != nil {
		fmt.Printf("Failed to deploy new token contract: %v", err)
		return
	}
	no, err := conn.PendingNonceAt(context.Background(), opt.From)
	log.Info("bind.TransactOpts", "nonce", no, "address", opt.From.String())
	nonce := int64(no) - 1
	for j := nftId; j < nftMax; j++ {
		nftId++
		nonce++
		opt.Nonce = big.NewInt(nonce)
		opt.GasLimit = 100000
		mintTx, err := contract.Mint(opt, Address1, big.NewInt(int64(nftId)))
		if err != nil {
			log.Info("mint_failed ", "nftId", nftId, "nonce", nonce, "err", err)
			continue
		}

		log.Info("mint_success ", "nftId", nftId, "nonce", nonce, "txHash", mintTx.Hash())
	}
	wg.Done()
}

func goodsTransfer(nftId, nftMax int, chainAddr, contractAddr string, opt *bind.TransactOpts, wg *sync.WaitGroup) {
	conn, err := ethclient.Dial(chainAddr)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	contract, err := goods.NewGoods(common.HexToAddress(contractAddr), conn)
	if err != nil {
		fmt.Printf("Failed to deploy new token contract: %v", err)
	}

	for j := nftId; j < nftMax; j++ {
		mintTx, err := contract.Transfer(opt, Address1, Address2, big.NewInt(int64(nftId)))
		if err != nil {
			fmt.Printf("Failed to Mint  contract: %v", err)
		}
		log.Info("transfer_success ", "time", time.Now().String(), "txHash", mintTx.Hash())
	}
	wg.Done()
}

func GoodsSuccessNum(chainAddr, contractAddr string, opt *bind.CallOpts) int64 {
	conn, err := ethclient.Dial(chainAddr)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return 0
	}

	contract, err := goods.NewGoods(common.HexToAddress(contractAddr), conn)
	if err != nil {
		fmt.Printf("Failed to deploy new token contract: %v", err)
		return 0
	}
	num, err := contract.GetSuccessNum(opt)
	if err != nil {
		fmt.Printf("Failed to contract.GetSuccessNum: %v", err)
		return 0
	}
	log.Info("GetSuccessNum", "num", num)
	return num.Int64()
}
