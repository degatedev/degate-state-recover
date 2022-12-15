# 关于（About）
StateRecovery是DeGate逃离模式提现的关键程序，任何用户均可在DeGate进入逃离模式后，在不依赖任何DeGate服务的情况下，通过执行StateRecovery程序来恢复资产树，并获得逃离模式需要的关键参数。逃离模式的执行参考[说明文档](https://docs.degate.com/testnet/how-to-withdraw-assets-in-exodus-mode)
# 编译（build）
## 使用DeGate编译好的可执行程序
release列表页面，下载最新的，适合自己机器环境的可执行程序。

例如，mac系统下载mac.zip，linux系统下载linux.zip
## 开发者，使用Docker进行编译
### 使用Dockerfile构建docker镜像
执行命令
```
cd docker
docker build -t zkp:1.0 .
```
### 使用镜像创建容器
执行命令
```
docker run --name=zkp -it zkp:1.0 bash
```
创建容器后会自动进入容器内
### 程序编译
执行命令
```
cd /go/src/degate-state-recover/cmd
./build_linux.sh
```
## 使用源码进行编译
### 安装golang环境
请参考golang[官网]('https://go.dev/')
### 依赖拉取
执行命令
```
go mod download
```
### 程序编译
1. 跳转cmd文件夹
```
cd cmd
```
2. 执行编译
```
go build -o staterecovery ./main.go
```
生成可执行程序`staterecovery`
## 配置说明
程序执行之前需要先准备配置文件，所有配置信息均可在以太坊链上获取。

配置文件为`config.toml`，与`start.sh`、`staterecovery`可执行程序在一个目录下。

无需修改，可直接使用的配置：
```
# 爬块构建的资产树的存储路径，如果程序执行中断，下次启动会从文件开始加载，即从上次中断位置开始继续处理
stateSavedFile = "state.json"
# 爬取的DeGate所有出块信息的存储路径，所有DeGate出块文件均存储在`blockFilePath`配置下的`blockfile`文件夹内
blockFilePath="./"

# 资产树构建间隔速度，0：0s秒，无间隔
loopInterval = 0 # second
# 资产树文件存储间隔，1000：每1000个DeGate出块执行一次存储
saveStateBlockInterval = 1000

# StateRecovery程序的输出产物，用于逃离模式下提取用户的资产，`merkleProof.json`文件内包含合约调用方法的输入参数，提取的账户与token信息由配置`withdrawModeAccount`与`withdrawModeToken`指定
withdrawModeFilePath = "merkleProof.json"
```

需要修改的配置：
```
# 需要提现的钱包地址
withdrawModeAccount="0xf8fedcc8855e569b17c3f69ec96547ae45c7684d"
# 需要提现的token地址
withdrawModeToken="0x466595626333c55fa7d7Ad6265D46bA5fDbBDd99"
# DeGate发起的第一个zkRollup的blockID，此blockID为以太坊的blockID
firstL1BlockID=7919204
# DeGate的exchange合约地址
exchangeContract="0xbcd394f0579db46f2b94d0490cee09bd34288c08"
# 以太坊rpc节点
chainNode="https://mainnet.infura.io"
```
### exchange合约地址获取
DeGate会预先将exchange合约地址写入到配置文件中，同时，DeGate会在公开的github仓库内公布所有DeGate部署的合约地址。
### chainNode获取
注册[Infura](https://www.infura.io/)账户，使用Infura提供的以太坊rpc节点
## 程序执行
赋予start.sh与staterecovery执行程序可执行权限。
```
chmod +x start.sh
chmod +x staterecovery
```
开始执行命令:
```
./start.sh
```
需要爬取所有DeGate的出块信息，执行时间较长，可以通过观察`blockfile`文件夹内的变化来查看程序执行的过程或查看`StateRecovery.log`日志文件来判定程序是否出现异常。
## 程序输出
程序执行完毕会自动停止，并输出构建好的资产树与合约提现输入参数，分别在`stateSavedFile`与`withdrawModeFilePath`配置的文件中。
用户获取其他资产提现信息时需修改配置文件中`withdrawModeToken`信息，然后再次执行`start.sh`脚本，此时无需再次爬取所有DeGate的出块信息，输出速度较快。