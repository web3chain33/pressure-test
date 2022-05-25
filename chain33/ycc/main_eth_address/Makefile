run:
    @chmod 777 ycc;chmod 777 ycc-cli
	@echo "ycc solo starting..."
	@nohup ./ycc -f ycc.toml > /dev/null 2>&1 &
	@echo "done"

stop:
	@./ycc-cli --rpc_laddr="http://127.0.0.1:7901" close

unlock:
	@./ycc-cli wallet unlock -p 1314fuzamei --rpc_laddr="http://127.0.0.1:7901"
clear:
	@rm -rf datadir && rm -rf wallet && rm -rf logs
init:
	@bash wallet-init.sh
	@echo "don't forget bind after sync"
bind:
	@./ycc-cli --rpc_laddr=http://localhost:7901 pos33 blsbind  #在init后同步完再执行，为了映射 bls 地址和挖矿地址，矿机使用bls聚合签名来投票
