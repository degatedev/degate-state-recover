package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"

	"github.com/ethereum/go-ethereum/common"
)

// Noop = 0,
// Transfer,
// SpotTrade,
// OrderCancel,
// AppKeyUpdate,
// BatchSpotTrade,
// Deposit,
// AccountUpdate,
// Withdrawal,
func ParseOtherTransactions(otherTransactionDatas [][]byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	if len(otherTransactionDatas) == 0 {
		return
	}
	for index, otherTransactionData := range otherTransactionDatas {
		otherTransactionDataBits := util.HexToBitsStr(common.Bytes2Hex(otherTransactionData))
		txTypeHex := util.BitsStrToHex(otherTransactionDataBits[:3])
		if txTypeHex == "0" {
			parseNoop(index, otherTransactionData)
		} else if txTypeHex == "1" {
			parseTransfer(index, otherTransactionData, accountUpdatePack, operatorAccountUpdate, state)
		} else if txTypeHex == "2" {
			parseSpotTrade(index, otherTransactionData, accountUpdatePack, operatorAccountUpdate, state)
		} else if txTypeHex == "3" {
			parseOrderCancel(index, otherTransactionData, accountUpdatePack, operatorAccountUpdate, state)
		} else if txTypeHex == "4" {
			parseAppKeyUpdate(index, otherTransactionData, accountUpdatePack, operatorAccountUpdate, state)
		} else if txTypeHex == "5" {
			parseBatchSpotTrade(index, otherTransactionData, accountUpdatePack, operatorAccountUpdate, state)
		}
	}
}

func ParseTransferToAccount(otherTransactionDatas [][]byte) ([]*entity.TransferToAccount) {
	transferToAccountList := make([]*entity.TransferToAccount, 0)
	if len(otherTransactionDatas) == 0 {
		return nil
	}
	for index, otherTransactionData := range otherTransactionDatas {
		otherTransactionDataBits := util.HexToBitsStr(common.Bytes2Hex(otherTransactionData))
		txTypeHex := util.BitsStrToHex(otherTransactionDataBits[:3])
		if txTypeHex == "1" {
			transferToAccount := parseTransferToAccount(index, otherTransactionData)
			if transferToAccount != nil {
				transferToAccountList = append(transferToAccountList, transferToAccount)
			}
		}
	}
	return transferToAccountList
}