run:
	@echo "ycc para starting..."
	@nohup ./ycc-para1 -f ycc.para1.toml > /dev/null 2>&1 &
	@echo "done"

show:
	@./ycc-para1 -f ycc.para1.toml

stop:
	@./ycc-cli --rpc_laddr="http://127.0.0.1:8911" close

unlock:
	@./ycc-cli --rpc_laddr="http://127.0.0.1:8911" wallet unlock -p fzm12345
clear:
	@rm -rf para*
	@rm -rf logs
init:
	@./ycc-cli --rpc_laddr="http://127.0.0.1:8911" seed save -s "huge position mixed captain spatial adapt visual ankle push upon fly top crop group fun"  -p fzm12345
	@make unlock
	@./ycc-cli --rpc_laddr="http://127.0.0.1:8911" account import_key -k "0x9d4f8ab11361be596468b265cb66946c87873d4a119713fd0c3d8302eae0a8e4" -l "genness"

