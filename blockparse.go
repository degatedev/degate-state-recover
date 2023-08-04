package staterecovery

import (
	"encoding/json"
	"errors"
	"log"
	"staterecovery/abihelper"
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/data"
	"staterecovery/transactionparse"

	"github.com/ethereum/go-ethereum/common"
)

func parseBlock(blockData *data.BlockData, state *entity.State) (err error) {
	// parse submitted block
	submittedBlocks, err := parseSubmittedBlock(blockData)
	if err != nil {
		return err
	}
	for _, submittedBlock := range submittedBlocks {
		// parse data
		parseData(submittedBlock, state)
	}
	
	return nil
}

func parseData(submittedBlock *entity.SubmittedBlock, state *entity.State) (err error) {
	// verify length of data
	if submittedBlock.Data == nil || len(submittedBlock.Data) != 167+int(submittedBlock.BlockSize)*83 {
		return errors.New("in parseData invalid data")
	}
	// exchange := submittedBlock.Data[:20]
	// merkleRootBefore := submittedBlock.Data[20:52]
	// merkleRootAfter := submittedBlock.Data[52:84]
	// assetMerkleRootBefore := submittedBlock.Data[84:116]
	// assetMerkleRootAfter := submittedBlock.Data[116:148]
	// timestamp := submittedBlock.Data[148:152]
	// protocolFeeBips := submittedBlock.Data[152:153]
	// numConditionalTransactions := submittedBlock.Data[153:157]
	operatorAccountID := util.BytesToIntStr(submittedBlock.Data[157:161])
	depositSize := util.BytesToInt(submittedBlock.Data[161:163])
	accountUpdateSize := util.BytesToInt(submittedBlock.Data[163:165])
	withdrawSize := util.BytesToInt(submittedBlock.Data[165:167])

	// log.Println("in parseData print data info")

	// // parse aux data
	withdrawAuxDataList := parseAuxData(submittedBlock.AuxiliaryData, depositSize, accountUpdateSize, withdrawSize)

	transactionparse.ParseTransaction(int(submittedBlock.BlockSize), depositSize, accountUpdateSize, withdrawSize, submittedBlock.Data[167:], withdrawAuxDataList, operatorAccountID, state)
	return nil
}

func parseSubmittedBlock(blockData *data.BlockData) (submittedBlock []*entity.SubmittedBlock, err error) {
	// decode binary
	// method: submitBlocksWithCallbacks or submitBlocks(latest contract)
	resArray, err := abihelper.AbiDecode(blockAbi, common.Bytes2Hex(blockData.Data))
	if err != nil {
		log.Println("in parseBlock first decode err:", err.Error())
		return nil, err
	}
	if resArray == nil || len(resArray) < 2 {
		log.Println("in parseBlock first decode resArray length invalid")
		return nil, errors.New("in parseBlock first decode resArray length invalid")
	}
	// decode binary
	// method: submitBlocks
	resArray, err = abihelper.AbiDecode(blockAbi, common.Bytes2Hex(resArray[1].([]byte)))
	if err != nil {
		log.Println("in parseBlock second decode err:", err.Error())
		return nil, err
	}
	if resArray == nil || len(resArray) != 1 {
		log.Println("in parseBlock second decode resArray length invalid")
		return nil, errors.New("in parseBlock second decode resArray length invalid")
	}
	blocksJson, err := json.Marshal(resArray[0])
	if err != nil {
		log.Println("in parseBlock submittedBlock Marshal err:", err.Error())
		return nil, err
	}
	var submittedBlocks []*entity.SubmittedBlock
	err = json.Unmarshal(blocksJson, &submittedBlocks)
	if err != nil {
		log.Println("in parseBlock submittedBlock Unmarshal err:", err.Error())
		return nil, err
	}
	return submittedBlocks, nil
}
