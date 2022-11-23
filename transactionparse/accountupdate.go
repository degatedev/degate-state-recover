package transactionparse

import (
	"log"
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
)

func ParseAccountUpdates(accountUpdateDatas [][]byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State, transferToAccountList []*entity.TransferToAccount) {
	if len(accountUpdateDatas) == 0 {
		return
	}
	for index, accountUpdateData := range accountUpdateDatas {
		parseAccountUpdate(index, accountUpdateData, accountUpdatePack, operatorAccountUpdate, state, transferToAccountList)
	}
}

// - Type: 1 byte (type > 0 for a conditional transaction)
// - Account owner: 20 bytes
// - Account ID: 4 bytes
// - Fee token ID: 4 bytes
// - Fee amount: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
// - Public key: 32 bytes
// - Nonce: 4 bytes
func parseAccountUpdate(index int, accountUpdateData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State, transferToAccountList []*entity.TransferToAccount) {
	// accountUpdateType := common.Bytes2Hex(accountUpdateData[:1])
	owner := util.BytesToIntStr(accountUpdateData[1:21])
	accountID := util.BytesToIntStr(accountUpdateData[21:25])
	feeTokenID := util.BytesToIntStr(accountUpdateData[25:29])
	feeAmount := util.ParseFloat(util.Bytes2Bits(accountUpdateData[29:31]), floatEncoding16)
	publicKey := common.Bytes2Hex(accountUpdateData[31:63])
	nonce := util.BytesToUInt32(accountUpdateData[63:67])
	accountIDBackup := util.BytesToIntStr(accountUpdateData[67:71])
	nonce = nonce + 1

	accountID = selectAccountID(accountID, accountIDBackup, state, owner, transferToAccountList)


	// updadte account
	publicKeyX, publicKeyY := decompressPublicKey(publicKey)
	accountUpdate := getAccountUpdate(accountID, accountUpdatePack, state)
	accountUpdate.AccountValue.PublicKeyX = publicKeyX.String()
	accountUpdate.AccountValue.PublicKeyY = publicKeyY.String()
	accountUpdate.AccountValue.Owner = owner
	accountUpdate.AccountValue.Nonce = nonce
	
	// 1: update user fee amount
	// 2: update operator fee amount
	// 3: update user account
	statemanager.SubBalance(accountUpdate, feeAmount.String(), feeTokenID)
	statemanager.AddBalance(operatorAccountUpdate, feeAmount.String(), feeTokenID)
	statemanager.UpdateAccountAsset(accountUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}

func selectAccountID(accountID, accountIDBackup string, state *entity.State, owner string, transferToAccountList []*entity.TransferToAccount) (selectedAccountID string) {
	if accountID == "0" && owner == state.ProtocolAccount || accountID != "0" {
		return accountID
	}
	if accountIDBackup != "0" {
		return accountIDBackup
	}
	return findAccountID(state, owner, transferToAccountList)
	// if accountID != "0" {
	// 	return accountID
	// }
	// return accountIDBackup
}

func findAccountID(state *entity.State, owner string, transferToAccountList []*entity.TransferToAccount) (foundAccountID string) {
	// find account from white account
	whiteAccount := getWhiteAccount(state, owner)
	if whiteAccount != nil {
		return whiteAccount.AccountID
	}
	// find account from tree where nonce == 0 and owner == accountUpdate.owner
	foundAccount := statemanager.GetAccountByOwnerAndNonce(state, owner, 0)
	if foundAccount != nil {
		return foundAccount.AccountID
	}
	// find account from transfer to account
	transferAccount := getTransferToAccount(transferToAccountList, owner)
	if transferAccount != nil {
		return transferAccount.AccountID
	}
	panic("in findAccountID can not find owner:" + owner)
}

func getWhiteAccount(state *entity.State, currentOwner string) *entity.WhiteAccount {
	if state.WhiteAccounts == nil || len(state.WhiteAccounts) == 0 {
		return nil
	}
	for _, whiteAccount := range state.WhiteAccounts {
		log.Println("whiteAccount owner:", whiteAccount.Owner, ";currentOwner:", currentOwner)
		if whiteAccount.Owner == currentOwner {
			return whiteAccount
		}
	}
	return nil
}

func getTransferToAccount(transferToAccountList []*entity.TransferToAccount, currentOwner string) *entity.TransferToAccount {
	if transferToAccountList == nil || len(transferToAccountList) == 0 {
		return nil
	}
	for _, account := range transferToAccountList {
		if account.Owner == currentOwner {
			return account
		}
	}
	return nil
}