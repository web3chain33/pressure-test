YCC公网配置和脚本

代码仓库
https://github.com/yccproject/ycc/commits/master 1c3539519ef6d5daa44ef740004dc06541fa02b7

节点运行指令
make run

节点清理指令
make clear

初始化初始节点指令
bash wallet-genesis.sh

初始化共识节点指令
bash wallet-init.sh

搭建步骤
1.拉取代码仓库 make build编译代码
2.将编译好的ycc和ycc-cli和本目录下的脚本都拷贝到服务器上
3.搭建初始化节点,make run运行，bash wallet-genesis.sh初始化挖矿
4.搭建共识节点，修改 配置文件下p2p.sub.dht的seeds和consensus.sub.pos33下的bootPeers,
将ip和端口改为初始节点的ip和对应的端口,make run运行,bash wallet-init.sh初始化共识节点,
等高度同步后再执行make bind进行绑定挖矿

2.共识节点:编译出来的程序要和初始节点一致
ycc.toml里singleMode为false
seeds和bootPeers的地址需要填入初始节点和其他节点的p2p信息
本节点的p2p信息可以通过./ycc-cli --rpc_laddr=http://localhost:9901 net peer 指令查看
修改ChainID,和初始节点保持一致
