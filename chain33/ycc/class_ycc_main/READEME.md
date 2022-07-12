类YCC主网配置和脚本

代码仓库
https://github.com/yccproject/ycc/commits/master e001a22a3e72b423e7e28a7fb39a13815dd07ccb

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

类ycc链的搭建：
需要根据需求修改ycc.go和ycc.toml文件
ycc.go里需要注意的三个配置
1.联盟链需要注释掉minTxFeeRate = 100000
2.修改evmChainID
3.修改superManager

初始节点:初始节点里一定要修改的是管理员地址,ycc.go里exec.sub.manage下的superManager,
建议使用钱包生成eth的私钥地址,填入地址即可
挖矿地址(consensus.sub.pos33.genesis下的minerAddr)和委托挖矿地址(returnAddr)可以修改(不是必须),
修改后需要将脚本里的私钥也换成对应的私钥
ycc.toml里需要将seeds和bootPeers的地址都注释掉
singleMode改为true
修改ChainID,需要未注册过的,可以在https://chainlist.org/zh上确认是否注册过
类ycc可以支持小狐狸钱包,如果要使用的话,需要先去https://github.com/ethereum-lists 提交新链信息的pr
参考ycc的pr https://github.com/ethereum-lists/chains/pull/1223/files

共识节点:编译出来的程序要和初始节点一致
ycc.toml里singleMode为false
seeds和bootPeers的地址需要填入初始节点和其他节点的p2p信息
本节点的p2p信息可以通过./ycc-cli --rpc_laddr=http://localhost:9901 net peer 指令查看
修改ChainID,和初始节点保持一致
