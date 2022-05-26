YCC的ETH地址版本配置和脚本

代码仓库
https://github.com/33cn/plugin/commits/master 7369d842ae803f2c5e35a79cc559ea58d12cd697

拉取后用make build指令编译chain33和chain33-cli可执行文件，改名为ycc和ycc-cli, 拷贝到服务器

单独测试环境的平行链需要使用plugin的版本进行编译，配置需要跟主链的配置对应上，修改plugin里的配置文件consensus.sub.para，consensus.sub.para需要比主链的ForkParacrossCommitTx大，也就是大于等于1，mainLoopCheckCommitTxDoneForkHeight需要比主链的ForkLoopCheckCommitTxDone大，也是大于等于1，startHeight=1，mainBlockHashForkHeight和主链的ForkBlockHash一致

配置和脚本仓库 https://github.com/chendehai/pressure-test/tree/master/chain33/ycc/para_eth_address ，里面的配置已经修改过，只用修改关于主链的节点配置，[consensus.sub.para]下的ParaRemoteGrpcClient里ip配置为主链的ip，端口配置为主链的GRPC端口

将该目录下的文件都拷贝到服务器

切换到服务器对应目录

执行 make run 指令启动ycc

执行 make init 指令导入钱包种子，挖矿私钥，解锁钱包



3.YCC平行链创建超级账户组

平行链节点跨链需要开启自共识，开启自共识=挖矿=创建超级账户组=支持跨链

一共需要2笔交易，均发送在主链上

参考 https://chain.33.cn/document/134

步骤1：

申请超级账户组

./ycc-cli --paraName user.p.para1. para nodegroup apply -a "0x862f23de306f0a35c122b11a6f0769570913d542" -c 0 --rpc_laddr "http://localhost:7901"

步骤2：

批准超级账户组

./ycc-cli --rpc_laddr "http://localhost:7901" --paraName user.p.para1. para nodegroup approve -i "0x2f60e8133ca38089190651134d20c59fd00adf201679b61355c9065aa3eccf81（替换为步骤1的最终交易哈希）" -c 0 -a "0x2f60e8133ca38089190651134d20c59fd00adf201679b61355c9065aa3eccf81（替换为步骤1的最终交易哈希）"

签名文档 https://chain.33.cn/document/93#1.2%20%E4%BA%A4%E6%98%93%E7%AD%BE%E5%90%8D%20SignRawTx
签名均采用 0x862f23de306f0a35c122b11a6f0769570913d542 的私钥签名,需要指定地址类型参数为ETH

发送完成后查看主链的交易结果和平行链的交易结果，正常情况下日志应该均为execok

4.发送跨链转账交易

参考https://chain.33.cn/document/266

汇总到postman里

https://www.getpostman.com/collections/698eb803898ea0a27fa8