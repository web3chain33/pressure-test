YCC私链配置和脚本

代码仓库
https://github.com/yccproject/ycc/commits/master 8e0e84848b5b9bf2f463c7f9422c555582cc5d8c

节点运行指令
make run

节点清理指令
make clear

初始化初始节点指令
bash wallet-genesis.sh

初始化共识节点指令
bash wallet-init.sh

搭建步骤
1.拉取代码仓库,将本目录下的ycc.go替换代码仓库的ycc.go文件,make build编译代码
2.将编译好的ycc和ycc-cli和本目录下的脚本都拷贝到服务器上
3.搭建初始化节点,make run运行，bash wallet-genesis.sh初始化挖矿
4.搭建共识节点，修改 配置文件下p2p.sub.dht的seeds和consensus.sub.pos33下的bootPeers,
将ip和端口改为初始节点的ip和对应的端口,make run运行,bash wallet-init.sh初始化共识节点,
等高度同步后再执行make bind进行绑定挖矿