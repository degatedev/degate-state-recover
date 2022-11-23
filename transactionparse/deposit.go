package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"
)

func ParseDeposits(depositDatas [][]byte, accountUpdatePack *entity.AccountUpdatePack, state *entity.State) {
	if len(depositDatas) == 0 {
		return
	}
	for index, depositData := range depositDatas {
		parseDeposit(index, depositData, accountUpdatePack, state)
	}
}

// - Type: 1 byte (0 for calling `deposit`, 1 for gas saving deposit)
// - Owner: 20 bytes
// - Account ID: 4 bytes
// - Token ID: 4 bytes
// - Amount: 31 bytes
func parseDeposit(index int, depositData []byte, accountUpdatePack *entity.AccountUpdatePack, state *entity.State) {
	// depositType := util.BytesToIntStr(depositData[:1])
	owner := util.BytesToIntStr(depositData[1:21])
	accountID := util.BytesToIntStr(depositData[21:25])
	tokenID := util.BytesToIntStr(depositData[25:29])
	amount := util.BytesToIntStr(depositData[29:60])

	accountUpdate := getAccountUpdate(accountID, accountUpdatePack, state)
	if amount != "" && amount != "0" {
		// update balance
		statemanager.AddBalance(accountUpdate, amount, tokenID)
	}
	accountUpdate.AccountValue.Owner = owner
	statemanager.UpdateAccountAsset(accountUpdate)
}