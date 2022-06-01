bin:=ycc
http_port:="http://127.0.0.1:8901"

run:
	@chmod 777 ycc;chmod 777 ycc-cli
	@echo "ycc para starting..."
	@nohup ./$(bin) -f $(bin).toml > /dev/null 2>&1 &
	@echo "done"

show:
	@./$(bin) -f $(bin).toml

stop:
	@./ycc-cli --rpc_laddr=$(http_port) close

unlock:
	@./ycc-cli --rpc_laddr=$(http_port) wallet unlock -p fzm12345

clear:
	@rm -rf para*
	@rm -rf logs

init:
	@./ycc-cli --rpc_laddr=$(http_port) seed save -s "huge position mixed captain spatial adapt visual ankle push upon fly top crop group fun"  -p fzm12345
	@make unlock
	@./ycc-cli --rpc_laddr=$(http_port) account import_key -k "0x9d4f8ab11361be596468b265cb66946c87873d4a119713fd0c3d8302eae0a8e4" -l "genness"

