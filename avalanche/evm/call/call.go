package call

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	l "github.com/33cn/chain33/common/log/log15"
	"github.com/chendehai/pressure-test/eth/solidity/goods"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var log = l.New("module", "call")

var keys9652 = []string{
	"3dcc4ae1d4b660354423cf36833d536dfda77f59f098574d371eb198e197f1fd",
	"a03450617e834879192b67c7fd88e092097f83c8e0fcc2b158b3b3ae062e7b3c",
	"abcb0a1a1d8767007d89aefa739a3b7f2761cff4ced1c1b5c6a6155176d12e42",
	"dbe12ccc68cf874c1588cb0f952b6012e8c928745d99b0e0ac6549311f873cfd",
	"3ce3b13a5e4014ad44d9f8a91849c518f5b63c565eaacd93685b91110fcbf143",
	"5da6a1c2a772135e1cf50d4ad2c18e55edfd3bb459a43725fc30ee029c7d5568",
	"3c1aaeea06511e84df69f95032c51c92d0d337ec28c364a97df5fb16f184bf0f",
	"04c83aa3da9c32887cd7b4ed193c8fe9a437ff40fff27b766c48dcc2b3e6a9fd",
	"fde39355eb7c15d1ad4ed8deb2b7612f1e751a598ba639185bc93be641c6c4ac",
	"ca889f3746955492d940414e1f3f148fdbb671f5151886b8f2fd78263e019039",
	"90e7334663dd01b270deeb247487613dd376b2103111b637d036e73ae73716f6",
	"ddb92e818e3f103c298c41e123b25325f56fef88d5d6281e23cde220d00e6415",
	"4c68a7454f6075191b98a378228cbb415ae423981ab964794babeca280fca859",
	"8ea8f139d8f7bb04c77994f46baf1ba5302a26ca3a57e2fbfc772c2908c8b0e4",
	"1f3fc4161c6db748a3ed6b1df0c9becc6fb93576e3a7a320737b8f2897bcd973",
	"304eb4941a2cefd6a1bef08af6f3a092330803e20964b729076f212d2048d6d5",
	"c3b0bc9465daf0f7f7c93d2016b07026a3e1578d4fcec054576a421623aca984",
	"c9be961a2352b90a3ccca7f445ec5b52e2a39e66d509a886c3f49c03ac7d6f9c",
	"48f0748cf368f634efa4b9cb4125f3f047d82226dbff23c0fea9689820e8945d",
	"7b05873a52da04b285f1c994167744ccdd90917e3e3847701f6cadce60fbf4b8",
	"e40937f678964641f225eb79aaefa25c93e17a67907fa7d565e28fa7a4de4a18",
	"a8f8ee98221a058037f4d884c69ec5486c545496ff2b44ca4c32b7b4bdf913ec",
	"2d25b7103c3c5285f6d361ba440aa2f261abf62458a789b66f6ed7b11ee2223d",
	"9867ba261a52dfa68cac5f923973fc14ff5a842000a46988331c8d3d63488fea",
	"6594c17fd87f701fcd0524a690cbd9428e08a18beb7770c6199731609f155543",
	"55f1ad21c1cd47a2033b4cdbb07c66296ea8acc2d1f7f08d6f0fb54f4ae5f238",
	"5c0824fbc332dda81b97edac35059c6a0b831992cd0e15282c04799cb34a3ebf",
	"02355393dce935ee762f88fb9dfcb32af59c0cae78b01e412d9a6a67bad607cf",
	"a20d8d6e1926e5ef2ea16cc331aa674e52c3de5721f68d9192f9ce1f959acd7f",
	"4b611fa669182cc68b98659c93e85f3c43245047e10e6075ddcdb79e3736a652",
	"166a5d2af4b056ecdba276446055b6cfc40728c2959175384f73f78fbb023be4",
	"6ccc4a154e3f4a903be7d34a75a069b6453036be88bf34a5574a0a3d177c8f02",
	"7b574140f02ade17b3844aa938fe8885087c1dd1ebd87479977ab9480c5985a1",
	"d599944d6a1f0b8661e8da349d91c05d3281d63cf1567f6f4eb4e80288e6eb61",
	"bd762040ba62fbb5686d6f3b8da24337416065ac576b2b4c4b9de27cc68ca534",
	"b8d8fc95ce60fa7a525bdc7d5031071f49ac3bf1a75cb2a7cf25b018337cce2c",
	"7e4971f42d457d75611b7733dff9238bc86b0e1f30c9b799191bef78c515ba99",
	"f5b454837d637323460f7bc6348f36ff5dcbea7b85c79da4ccc8724a67d9273f",
	"bcf93aa6cd856bc92930ddeba300257c7893ed79d1424ddf0508ab62194ab922",
	"57953f30af25610b79c73a8bc85a2fbe5d5fcfe3f05996d0cedc82482fb82def",
	"a5beb864e41348672e820dd64324b5c42c74c2c77e93bc0a8c4594a03c2c3f83",
	"9222d1cee10fcda5e35f5743178145d801a90f2fb0d6d48966d42495db030407",
	"cf5ec87e17f5d9b445342c2edf52ddc449e48176fa2b3ae7ca53a7f9afaf2852",
	"4aa7fc24073288c2bcb7d4704237bbb6c042f0e37992cbffcd07229e1cc6e37c",
	"6d02586d5c46de7421ee88ec3ed762a358c78d0ad5961ed9ea6e229cd7fec689",
	"034ef046048556197fc0abbce61e888194027cf291970cbec7096b8cb98c3cca",
	"d3b767a77c12129e186644dff92c8c3feff60753f6a86971596f3caa54e62125",
	"8c5be4e3ddacde8525f4c5008038fb0a9ffa3a1950a5dd3c56804e13519cf5a5",
	"1d89c30d2956a7f7f149cc45c900cb93e7741797601e2f06f3a0c46f894c7f75",
	"63ebf414772ffe416af251096e8914b9c7e187d1c1a21de0c19c70c65e9c366e",
	"78e128b8eccc4c89c486aac45a0277fb8d82b8b9ad5168997f68d1a15c0705ec",
	"049c3ed54eaf8dd7e6d57c30659ec708561314831647f8b53b87d9f53b2e1fc2",
	"be03b5311e654785a8ed18eaca6c90a811699479310d3b173789d382629e721b",
	"36afc89e67ea8540576ae15fd4f5efddcb645ef4aadcbdb06929c13d1f08b4fe",
	"df334643d1004cd06c21d0f52ff7c9fb93943cd2500cf509c54dbc732d5ff32b",
	"b67f388f692ea0c3e426febacc13b2aea4a2932fc1e12e3a531018b8df6bb68a",
	"f414146baf8ea89fe854e97138b74b1453545618b6e270e2c835dbb52e241cda",
	"15bddb4aaa6be5291efa7e01418dbbc21bcf0a99a6100919074916fda7bc1ec9",
	"2b0f92007b989822b213cfb35cc1990fe6bf5c658de9bffe79b4fdf9ef5db145",
	"826220d79e3782bf1825ee0a3b37a6bc3f1a3023f6279fa299d265c38a7bfb39",
	"dbf058efad55cfb449022ac61950c2e2628a7cbe56630c5c4fd6b3ae8012eac5",
	"62bf4ffc4c9fa202f9eedda2eb70f2e218fb9bde8a7dc07fd12e873c528a15eb",
	"8607b811e035cc209b07c31941c250db1b2d86e784d2fc34b91621c09c85cd1e",
	"514b7ef89488338a7950a58354d7a3adaa9a19daa230c6c24c128bdad2925c91",
	"da4c50592b1da35a1ffecf1254f944480389e7e814c6feb10d4178af852399c4",
	"bbd71b3b00ef5656eaa00e5dfdbedca5cd66b2c0c71e8e61231e21f547e01342",
	"831ddc8cb904b46bd198289081e77adacbd674ff67b5c72785a26e8312076b3b",
	"b8c664e7849f5aaed5bd2749c7c3ce735864fb9aa6913c7e6a026018cf34d366",
	"e539447f4b86db40a3a5911973f8d2c96f519c69db2747ccf21809f03ef52e9b",
	"3e7bf2c7014dd9d736ac85999a82b6fdd77135c121aca9c24436c8fe0fe46301",
}

var keys9672 = []string{
	"0d40510fbae754aea1c05c3a6c71e4961736ddc35cce2215b59afaeb8d99953a",
	"85c9399205383f7901f1041395d2d4fcbf391316fadf24f867c3a9b62d6c6e89",
	"f3c333f1cf97a5c48fcc73b4a19627c030b2c0d3852354598198b9a551651db9",
	"ef4c05d922f00e33688553c5da0fb105c6bb6a3e74d5007b776439a476acc57b",
	"e5d4e8eb73a86c8878712a7e1744f3619e10e924330dfaad9338a9d60d79be12",
	"1a811b656cfd5455df30fa60c2fd9d459e632d36fe3eeb6e61e241f9c8446f45",
	"cf578aa126488301860857ef6722cbc244fea2ae3c61d1b6bb94850b66a550d0",
	"fb2c56b81ac77e22e710aabc091d1e999ad709c7c0e62e87ba9a924b0411e521",
	"ed08757cf029ca4b2eb4751d8aeb18ac1a8a743034494ad73cccf211b87dbb3f",
	"a566184848532b398abfffdae7ec3a72af3b37bc441dbafb6de5ba5ed9913b5d",
	"a7c91f7b45271325f914f45459e4cd2ee43b5605729b60a30b445f3417c53c62",
	"b2b16437bcf8762bf389794e186354aab2393f00feb27cc17062a1542085f000",
	"d167effd1af8293c5aa620f1ebe4ce52ebfb18003eb687a77c46a1936e618e2e",
	"64716f0e158ae2486ca80d2188d3ae3273362f0ec782c90e3c20182931eca284",
	"41012a848733ff2c322d275fc904084d7191ae09e6e8f9acd1bd2e9d540904be",
	"e592aa797b476bb1ccc13b60b5d7ad476b6c594cd60abd0c55708446f8762c93",
	"3203bb1bfc2b04fa597c41b83f71ed71092db64d5df7c0adc4e71f937e2a33ac",
	"fc0327c3f56aabb019d84c02a4ad9d7575c7d07ab5f959ab9f4466d21eee79e4",
	"9e82e5a28b7823c31a663798670a55536053b09b8e6964594326197c632e4981",
	"992f494eb67f0464e59df4de72e58284da74d421c91d3c35668471cef0ce9962",
	"2c947ea09b99fa9ec3bad25101b0f062001a4b323d9ec5cb5d13eaca82c2ea50",
	"fc67ad512a6ead695b0e550370c6514bdc60fb7fc9263bc2dbc94282909d2aa5",
	"f4b7cc3ebd53139bed5556946c7de0e4a7ab11b2f6ac96821c7da234553083f0",
	"c1dfe4e6700dc2a51b0d8f2ca2e90d21eab89bb02ea98aa3bf06a1051ce369cf",
	"99ff8713b1f2e7ff94326ba0026c34805847f6ed021416b0b529bcb0b278c621",
	"94ca82d9bf341cd32f4cbf51cd9216eca1f997efbb256f4f9752e739073028e2",
	"376737dada3c6ebee1c72cca8b36f288c29515064586db27b198c1fe96f9c634",
	"d99973842cead10c3d7f16eb07e22cd025142de5ef06bb23cd3c368686ac05f6",
	"fe0b3f7c80ba0ca14e0d6bf3156b51e74774b3febe9745e17ef8ed215915e0f1",
	"174efc11448d1948995e6fd4ddba2f8532786d2bff0d2e2520bb2e7b3b380499",
	"75bfa2314e412205ce48dac840c5cfcae656bf2d135cc3181c652cf609a37c56",
	"0924dde0368534ab6e217642b2e97a541bf75c69e11dc8b6f8a40df1a3d3a397",
	"c4c79a5c5f36ce5663f8fadf246a6b019a138c59d4348726d513c5f6516f96c3",
	"d8819a89a666547feb3c82ae137a2e63d81bb982bdf206f0f89258fbc68f653b",
}

var keys = keys9652

var AccountOpts = InitOpts(keys)

func InitOpts(keys []string) []*bind.TransactOpts {
	opts := make([]*bind.TransactOpts, 0, len(keys))
	for _, key := range keys {
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			fmt.Printf("Failed to crypto.HexToECDSA: %v", err)
			continue
		}
		opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43112))
		if err != nil {
			fmt.Printf("Failed to connect to the Ethereum client: %v", err)
			continue
		}
		opts = append(opts, opt)
	}
	return opts
}

func CreateAccountOpts(n int, chainAddr string) []*bind.TransactOpts {
	conn, err := ethclient.Dial(chainAddr)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return nil
	}
	opts := make([]*bind.TransactOpts, 0, n)
	for i := 0; i < n; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			fmt.Println("crypto.GenerateKey, err:", err)
			return nil
		}

		fmt.Println(privateKey)
		fmt.Println(conn)

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
		case 2:
			y := i % len(AccountOpts)
			go goodsBatchMint(nftId, nftId+nftStep, chainAddr, contractAddr, AccountOpts[y], wg)
		case 3:
			y := i % len(AccountOpts)
			go goodsTransfer(nftId, nftId+nftStep, chainAddr, contractAddr, AccountOpts[y], wg)
		case 5:
			y := i % len(AccountOpts)
			go avaxTransfer(chainAddr, AccountOpts[y], wg)
		default:
			fmt.Println("error operationType:", operationType)
		}
		nftId += nftStep
	}
}

func GoodsSuccessNum(chainAddr, contractAddr string, opt *bind.CallOpts) {
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
	num, err := contract.GetSuccessNum(opt)
	if err != nil {
		fmt.Printf("Failed to contract.GetSuccessNum: %v", err)
		return
	}
	fmt.Println("GetSuccessNum num=", num)
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
	mintTx := &types.Transaction{}
	for j := nftId; j < nftMax; j++ {
		nftId++
		nonce++
		opt.Nonce = big.NewInt(nonce)
		opt.GasLimit = 100000
		mintTx, err = contract.Mint(opt, opt.From, big.NewInt(int64(nftId)))
		if err != nil {
			log.Info("mint_failed ", "nftId", nftId, "nonce", nonce, "err", err)
			continue
		}

		log.Info("mint_success ", "nftId", nftId, "nonce", nonce, "txHash", mintTx.Hash())
	}
	log.Info("mint_finish", "nftId", nftId, "nonce", nonce, "txHash", mintTx.Hash())
	ctx, concel := context.WithTimeout(context.Background(), 40*time.Second)
	defer concel()

	WaitTransaction(ctx, conn, mintTx, nftId)
	wg.Done()
}

func goodsBatchMint(nftId, nftMax int, chainAddr, contractAddr string, opt *bind.TransactOpts, wg *sync.WaitGroup) {
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
	mintTx := &types.Transaction{}

	for j := nftId; j < nftMax; j += 5 {
		nftId++
		nonce++
		opt.Nonce = big.NewInt(nonce)
		opt.GasLimit = 100000

		ids := []*big.Int{
			big.NewInt(int64(nftId)),
			big.NewInt(int64(nftId + 1)),
			big.NewInt(int64(nftId + 2)),
			big.NewInt(int64(nftId + 3)),
			big.NewInt(int64(nftId + 4)),
		}
		nftId += 4
		mintTx, err = contract.BatchMint(opt, opt.From, ids)
		if err != nil {
			log.Info("mint_failed ", "nftId", nftId, "nonce", nonce, "err", err)
			continue
		}

		log.Info("mint_success ", "nftId", nftId, "nonce", nonce, "txHash", mintTx.Hash())
	}
	log.Info("mint_finish", "nftId", nftId, "nonce", nonce, "txHash", mintTx.Hash())
	ctx, concel := context.WithTimeout(context.Background(), 40*time.Second)
	defer concel()

	WaitTransaction(ctx, conn, mintTx, nftId)
	wg.Done()
}

func WaitTransaction(ctx context.Context, conn *ethclient.Client, tx *types.Transaction, nftId int) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, _ := conn.TransactionReceipt(ctx, tx.Hash())
		if receipt != nil {
			fmt.Println("nftId=", nftId, "receipt.log", receipt.Logs)
			header, _ := conn.HeaderByHash(context.Background(), receipt.BlockHash)
			if header != nil {
				fmt.Println("nftId=", nftId, "header", header.Time)
			}
			return
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			fmt.Println("nftId=", nftId, "WaitTransaction timeout")
			return
		case <-queryTicker.C:
		}
	}
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
			continue
		}
		log.Info("transfer_success ", "time", time.Now().String(), "txHash", mintTx.Hash())
	}
	wg.Done()
}

func avaxTransfer(chainAddr string, opt *bind.TransactOpts, wg *sync.WaitGroup) {
	conn, err := ethclient.Dial(chainAddr)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	privateKey, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	if err != nil {
		fmt.Printf("Failed to crypto.HexToECDSA: %v", err)
		return
	}

	toAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	no, err := conn.PendingNonceAt(context.Background(), opt.From)
	log.Info("bind.TransactOpts", "nonce", no, "address", opt.From.String())
	nonce := int64(no) - 1

	signedTx := &types.Transaction{}
	for i := 0; i < 1000; i++ {
		nonce++
		rawTx := types.NewTx(&types.LegacyTx{
			To:       &toAddr,
			Nonce:    uint64(nonce),
			GasPrice: big.NewInt(27000000000),
			Gas:      22000,
			Value:    big.NewInt(100),
			Data:     nil,
		})

		signedTx, err = opt.Signer(opt.From, rawTx)
		if err != nil {
			fmt.Printf("Failed to signedTx: %v \n", err)
			continue
		}
		err = conn.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Printf("Failed to SendTransaction: %v \n", err)
			continue
		}
		log.Info("mint_success ", "nonce", nonce, "txHash", signedTx.Hash())
	}
	wg.Done()
}
