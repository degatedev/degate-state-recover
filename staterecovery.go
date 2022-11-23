package staterecovery

import (
	"encoding/json"
	"log"
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/conf"
	"staterecovery/data"
	"staterecovery/statemanager"
	"strconv"
	"time"
)

var (
	blockAbi = `[{"inputs":[{"components":[{"internalType":"uint8","name":"blockType","type":"uint8"},{"internalType":"uint16","name":"blockSize","type":"uint16"},{"internalType":"uint8","name":"blockVersion","type":"uint8"},{"internalType":"bytes","name":"data","type":"bytes"},{"internalType":"uint256[8]","name":"proof","type":"uint256[8]"},{"internalType":"bool","name":"storeBlockInfoOnchain","type":"bool"},{"internalType":"bytes","name":"auxiliaryData","type":"bytes"},{"internalType":"bytes","name":"offchainData","type":"bytes"}],"internalType":"struct ExchangeData.Block[]","name":"blocks","type":"tuple[]"}],"name":"submitBlocks","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"components":[{"internalType":"uint32","name":"accountID","type":"uint32"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"pubKeyX","type":"uint256"},{"internalType":"uint256","name":"pubKeyY","type":"uint256"},{"internalType":"uint32","name":"nonce","type":"uint32"}],"internalType":"struct ExchangeData.AccountLeaf","name":"accountLeaf","type":"tuple"},{"components":[{"internalType":"uint32","name":"tokenID","type":"uint32"},{"internalType":"uint248","name":"balance","type":"uint248"}],"internalType":"struct ExchangeData.BalanceLeaf","name":"balanceLeaf","type":"tuple"},{"internalType":"uint256[48]","name":"accountMerkleProof","type":"uint256[48]"},{"internalType":"uint256[48]","name":"balanceMerkleProof","type":"uint256[48]"}],"internalType":"struct ExchangeData.MerkleProof","name":"merkleProof","type":"tuple"}],"name":"withdrawFromMerkleTree","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bool","name":"isDataCompressed","type":"bool"},{"internalType":"bytes","name":"data","type":"bytes"},{"components":[{"components":[{"internalType":"uint16","name":"blockIdx","type":"uint16"},{"components":[{"internalType":"uint16","name":"txIdx","type":"uint16"},{"internalType":"uint16","name":"numTxs","type":"uint16"},{"internalType":"uint16","name":"receiverIdx","type":"uint16"},{"internalType":"bytes","name":"data","type":"bytes"}],"internalType":"structLoopringIOExchangeOwner.TxCallback[]","name":"txCallbacks","type":"tuple[]"}],"internalType":"structLoopringIOExchangeOwner.BlockCallback[]","name":"blockCallbacks","type":"tuple[]"},{"internalType":"address[]","name":"receivers","type":"address[]"}],"internalType":"structLoopringIOExchangeOwner.CallbackConfig","name":"config","type":"tuple"}],"name":"submitBlocksWithCallbacks","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bool","name":"isDataCompressed","type":"bool"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"submitBlocks","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	methodSubmitBlocks = "submitBlocks"
	dataResource data.DataResource
)

func initDataResource(config *conf.Config) {
	if dataResource != nil {
		return
	}
	// read data from db
	dataResource = data.NewDataResource(config)
}

func StartRecovery(config *conf.Config) (state *entity.State) {
	initDataResource(config)
	state = recoveryState(config)
	// create withdraw mode data
	createWithdrawModeMerkleProof(config, state)
	return state
}

func recoveryState(config *conf.Config) (state *entity.State) {
	// load state from file or create empty state
	state = loadState(config.StateSavedFile)
	if state.BlockId == config.LastL2BlockID {
		return state
	}
	if state.BlockId > config.LastL2BlockID {
		panic("in StartRecovery state.BlockId > config.LastL2BlockID state.BlockId:" + strconv.Itoa(state.BlockId) + ";config.LastL2BlockID:" + strconv.Itoa(config.LastL2BlockID))
	}
	// set some config into state
	setConfigIntoState(config, state)
	startBlockID := state.BlockId
	for true {
		// start next block parse
		state.BlockId++
		blockData, err := dataResource.GetDataByBlockId(state.BlockId)
		if err != nil {
			log.Println("in StartRecovery GetDataByBlockId:", err.Error(), "; maybe loop end.")
			state.BlockId--
			saveState(startBlockID, config.StateSavedFile, state)
			break
		}
		parseBlock(blockData, state)

		if state.BlockId == config.LastL2BlockID {
			saveState(startBlockID, config.StateSavedFile, state)
			break
		}
		// save state file
		if state.BlockId % config.SaveStateBlockInterval == 0 {
			saveState(startBlockID, config.StateSavedFile, state)
		}

		time.Sleep(time.Second * time.Duration(config.LoopInterval))
	}
	return state
}

func setConfigIntoState(config *conf.Config, state *entity.State) {
	state.ProtocolAccount = config.ProtocolAccount
	state.WhiteAccounts = config.WhiteAccounts
}

func loadState(statePath string) *entity.State {
	// load from file
	fileExist, _ := util.PathExists(statePath)
	if !fileExist {
		return statemanager.CreateEmptyState()
	}
	content, err := util.ReadTxtFile(statePath)
	if err != nil {
		panic("in loadState state file can not open err:" +  err.Error())
	}
	state := &entity.State{}
	err = json.Unmarshal([]byte(content), state)
	if err != nil {
		panic("in loadState state file can not json Unmarshal err:" +  err.Error())
	}
	return state
}

func saveState(startBlockID int, statePath string, state *entity.State) {
	if startBlockID == state.BlockId {
		// same blockID, not need to update state
		return
	}
	// save file to temp
	stateJson, err := json.Marshal(state)
	if err != nil {
		panic("in saveState state to json error:" + err.Error())
	}
	err = util.WriteTxtFile(string(stateJson), statePath)
	if err != nil {
		panic("in saveState write state tmp file error:" + err.Error())
	}
}