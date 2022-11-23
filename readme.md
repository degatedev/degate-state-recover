# About
`StateRecovery` is an important program for withdrawing assets on DeGate exodus mode. Any user can restore the Asset Merkle Tree by executing the `StateRecovery` program after DeGate enters the exodus mode without relying on any DeGate services and obtaining the key parameters required by the exodus mode. Execution reference of exodus mode [document](https://docs.degate.com/testnet/how-to-withdraw-assets-in-exodus-mode)
# Build
## Ordinary users, executable programs compiled with DeGate
Goto release page of this github repository, download the latest executable program suitable for your machine OS.

For example, mac-OS downloads mac.zip, linux-OS downloads linux.zip
## Developers, use source code to compile
### Install the golang environment
Please refer to golang [official website]('https://go.dev/')
### Dependency pull
Excute the command:
```
go mod download
```
### program compilation
1. Go to `cmd` folder
```
cd cmd
```
2. Execute compile
```
go build -o staterecovery ./main.go
```
generate executable program`staterecovery`
## Configuration instructions
The configuration file needs to be prepared before the program is executed, and all  parameters can be obtained on the Ethereum chain.

The configuration file is `config.toml`, which is in the same directory as `start.sh` and `staterecovery` executable file.

Configuration that can be used directly without modification:
```
# The storage path of the Asset Merkle Tree built by crawling DeGate-Block. If the program execution is interrupted, the next startup will start loading from the file, that is, continue processing from the last interrupted position
stateSavedFile = "state.json"
# The storage path of all the crawled DeGate-Block information, all DeGate-Block files are stored in the `blockfile` folder under the `blockFilePath` configuration
blockFilePath="./"

# Asset Merkle Tree build interval speed, 0:0s seconds, no interval
loopInterval = 0 # second
# Asset Merkle Tree file storage interval, 1000: perform storage every 1000 DeGate blocks
saveStateBlockInterval = 1000

# The output product of the `StateRecovery` program is used to extract the user's assets in escape mode. The `merkleProof.json` file contains the input parameters of the contract call method. The extracted account and token information are specified by the configuration `withdrawModeAccount` and `withdrawModeToken`
withdrawModeFilePath = "merkleProof.json"
```

Configuration that needs to be modified:
```
# The wallet address that needs to be withdrawn
withdrawModeAccount="0xf8fedcc8855e569b17c3f69ec96547ae45c7684d"
# The token address that needs to be withdrawn
withdrawModeToken="0x466595626333c55fa7d7Ad6265D46bA5fDbBDd99"
# The blockID of the first zkRollup initiated by DeGate, which is the blockID of Ethereum
firstL1BlockID=7919204
# DeGate's exchange contract address
exchangeContract="0xbcd394f0579db46f2b94d0490cee09bd34288c08"
# Ethereum rpc node
chainNode="https://mainnet.infura.io"
```
### Exchange contract address acquisition
DeGate will pre-prepare the addresses of the exchange contract in the configuration file. Anyone can verify the addresses easily through explorers.
### chainNode get
Register [Infura](https://www.infura.io/) account and use the Ethereum rpc node provided by Infura
## program execution
Give `start.sh` and `staterecovery` execution programs executable permissions.
```
chmod +x start.sh
chmod +x staterecovery
```
Execute the command:
```
./start.sh
```
It needs to crawl all the DeGate-Block, and the execution time is long. You can check the program execution process by observing the file created in the `blockfile` folder or check the status in `StateRecovery.log`.
## Output
Program will stop automatically when its done, and output the constructed Asset Merkle Tree and input parameters of the contract method `withdrawFromMerkleTree`.
### Constructed Asset Merkle Tree
The constructed Asset Merkle Tree is in the configuration file of `stateSavedFile`, which is a Merkle tree containing all DeGate user accounts and asset information, which will only be generated once.
### Input parameters
The input parameters of the contract method `withdrawFromMerkleTree` are in the configuration file of `withdrawModeFilePath`.
`withdrawFromMerkleTree` can only extract one token of one user at a time. If you need to extract multiple tokens, you need to modify the configuration of `withdrawModeToken` and execute `start.sh` again to get new input parameters.
