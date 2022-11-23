package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
)

// - Type Tx: 3 bits
// - Type Tx Pad: 1 bit
// - Account ID: 4 bytes
// - Fee token ID: 4 bytes
// - Fee amount: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
// - Nonce: 4 bytes
func parseAppKeyUpdate(index int, appKeyUpdateData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	appKeyUpdateDataHex := common.Bytes2Hex(appKeyUpdateData)
	// txType := appKeyUpdateDataHex[:1]
	accountID := util.HexToIntStr(appKeyUpdateDataHex[1:9])
	feeTokenID := util.HexToIntStr(appKeyUpdateDataHex[9:17])
	feeAmount := util.ParseFloat(util.HexToBits(appKeyUpdateDataHex[17:21]), floatEncoding16)
	nonce := util.HexToUInt32(appKeyUpdateDataHex[21:29])
	nonce = nonce + 1

	// updadte account
	accountUpdate := getAccountUpdate(accountID, accountUpdatePack, state)
	accountUpdate.AccountValue.Nonce = nonce
	if accountID == "0" {
		panic("in app key")
	}
	// 1: update user fee amount
	// 2: update operator fee amount
	// 3: update user account
	statemanager.SubBalance(accountUpdate, feeAmount.String(), feeTokenID)
	statemanager.AddBalance(operatorAccountUpdate, feeAmount.String(), feeTokenID)
	statemanager.UpdateAccountAsset(accountUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}