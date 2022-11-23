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
// - StorageID: 4 bytes
// - Fee token ID: 4 bytes
// - Fee amount: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
func parseOrderCancel(index int, orderCancelData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	orderCancelDataBits := util.HexToBitsStr(common.Bytes2Hex(orderCancelData))
	// txType := orderCancelDataBits[:3]
	accountID := util.BitsStrToIntStr(orderCancelDataBits[4:36])
	// storageID := util.BitsStrToIntStr(orderCancelDataBits[36:68])
	feeTokenID := util.BitsStrToIntStr(orderCancelDataBits[68:100])
	feeAmount := util.ParseFloat(util.BitsStrToFixedBits(orderCancelDataBits[100:116], 16), floatEncoding16)

	accountUpdate := getAccountUpdate(accountID, accountUpdatePack, state)

	// 1: update user fee
	// 2: update operator fee
	// 3: update user account
	statemanager.SubBalance(accountUpdate, feeAmount.String(), feeTokenID)
	statemanager.AddBalance(operatorAccountUpdate, feeAmount.String(), feeTokenID)

	statemanager.UpdateAccountAsset(accountUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}