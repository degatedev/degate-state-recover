package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
)

func ParseWithdraws(withdrawDatas [][]byte, accountUpdatePack *entity.AccountUpdatePack, withdrawAuxDataList []*entity.WithdrawAuxData, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	if len(withdrawDatas) == 0 {
		return
	}
	for index, withdrawData := range withdrawDatas {
		parseWithdraw(index, withdrawData, accountUpdatePack, withdrawAuxDataList[index], operatorAccountUpdate, state)
	}
}

// - Type: 1 bytes (type > 0 for conditional withdrawals, type == 2 for a valid forced withdrawal, type == 3 when invalid)
// - Owner: 20 bytes
// - Account ID: 4 bytes
// - Token ID: 4 bytes
// - Fee token ID: 4 bytes
// - Fee amount: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
// - StorageID: 4 bytes
// - OnchainDataHash: 20 bytes
func parseWithdraw(index int, withdrawData []byte, accountUpdatePack *entity.AccountUpdatePack, withdrawAuxData *entity.WithdrawAuxData, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	// withdrawType := withdrawData[:1]
	// Owner := util.BytesToIntStr(withdrawData[1:21])
	accountID := util.BytesToIntStr(withdrawData[21:25])
	tokenID := util.BytesToIntStr(withdrawData[25:29])
	feeTokenID := util.BytesToIntStr(withdrawData[29:33])
	feeAmount := util.ParseFloat(util.HexToBits(common.Bytes2Hex(withdrawData[33:35])), floatEncoding16)
	// storageID := util.BytesToIntStr(withdrawData[35:39])
	// onchainDataHash := util.BytesToIntStr(withdrawData[39:59])

	accountUpdate := getAccountUpdate(accountID, accountUpdatePack, state)
	statemanager.SubBalance(accountUpdate, withdrawAuxData.Amount.String(), tokenID)
	statemanager.SubBalance(accountUpdate, feeAmount.String(), feeTokenID)

	statemanager.AddBalance(operatorAccountUpdate, feeAmount.String(), feeTokenID)

	statemanager.UpdateAccountAsset(accountUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}