## 快速使用教程
1. 首先根据操作系统类型下载发送交易的对应可执行文件，下载地址：  
2. 运行下载来的可执行文件，对于习惯使用鼠标操作的用户，双击即可运行  
3. 发送交易程序运行后，默认会发送20万笔交易，在区块链浏览器界面，打开链接即可看到区块信息: https://www.yuan.org/block  
4. 通过计算一段时间的总的交易量与耗时，就可以得出区块链的交易速度  
注意: 发送会消耗手续费，上面的发送程序，默认扣费地址为 0x8efd65cacad0c0a7b3eace77eeaac04476943980 ，  
如果使用时，提示 NoBalance，表示该地址里面，已经没有余额了，解决办法一个是使用下面的带配置选项的交易发送程序，修改扣费地址为一个有余额的地址，
或者是联系工作人员，给这个默认地址转账，工作人员邮箱为： ycc@yuan.org


## 带配置选项的交易发送使用教程
1. 使用流程和公链发送工具基本一样，只是增加了配置文件，可以配置发送的各项参数，带配置选项发送工具的下载地址为：
2. 将交易发送工具和配置文件放置在同一个目录，发送交易的程序运行后，会自动加载当前目录下的 config.yaml 的配置文件，然后根据配置，发送交易
3. 交易发送完成后，按照上面的方法查看和计算区块链交易速度