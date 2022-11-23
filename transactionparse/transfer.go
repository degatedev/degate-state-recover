package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
)

// - Type Tx: 3 bits
// - Type Tx Pad: 1 bit
// - Type: 1 bytes (must be 0)
// - From account ID: 4 bytes
// - To account ID: 4 bytes
// - Token ID: 4 bytes
// - Amount: 4 bytes （32 bits, 25 bits for the mantissa part and 7 for the exponent part)
// - Fee token ID: 4 bytes
// - Fee amount: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
// - StorageID: 4 bytes
// - To: 20 bytes （only set when transferring to a new account)
// - From: 20 bytes （only set for conditional transfers)
func parseTransfer(index int, transferData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	transferDataBits := util.HexToBitsStr(common.Bytes2Hex(transferData))
	// txType := transferDataBits[:3]
	// transferType := util.BitsStrToIntStr(transferDataBits[4:12])
	fromAccountID := util.BitsStrToIntStr(transferDataBits[12:44])
	toAccountID := util.BitsStrToIntStr(transferDataBits[44:76])
	tokenID := util.BitsStrToIntStr(transferDataBits[76:108])
	amount := util.ParseFloat(util.BitsStrToFixedBits(transferDataBits[108:140], 32), floatEncoding32)
	feeTokenID := util.BitsStrToIntStr(transferDataBits[140:172])
	feeAmount := util.ParseFloat(util.BitsStrToFixedBits(transferDataBits[172:188], 16), floatEncoding16)
	// storageID := util.BitsStrToIntStr(transferDataBits[188:220])
	to := "0"
	if len(transferDataBits) > 220 {
		to = util.BitsStrToIntStr(transferDataBits[220:380])
	}
	// from := ""
	// if len(transferDataBits) > 380 {
	// 	from = util.BitsStrToIntStr(transferDataBits[380:540])
	// }

	fromAccountUpdate := getAccountUpdate(fromAccountID, accountUpdatePack, state)
	statemanager.SubBalance(fromAccountUpdate, amount.String(), tokenID)
	statemanager.SubBalance(fromAccountUpdate, feeAmount.String(), feeTokenID)

	toAccountUpdate := getAccountUpdate(toAccountID, accountUpdatePack, state)
	if to != "0" {
		toAccountUpdate.AccountValue.Owner = to
	}

	statemanager.AddBalance(toAccountUpdate, amount.String(), tokenID)
	statemanager.AddBalance(operatorAccountUpdate, feeAmount.String(), feeTokenID)

	statemanager.UpdateAccountAsset(fromAccountUpdate)
	statemanager.UpdateAccountAsset(toAccountUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}

func parseTransferToAccount(index int, transferData []byte) *entity.TransferToAccount {
	transferDataBits := util.HexToBitsStr(common.Bytes2Hex(transferData))
	toAccountID := util.BitsStrToIntStr(transferDataBits[44:76])
	toAddress := "0"
	if len(transferDataBits) > 220 {
		toAddress = util.BitsStrToIntStr(transferDataBits[220:380])
	}
	if toAddress == "0" {
		return nil
	}
	return &entity.TransferToAccount{
		AccountID: toAccountID,
		Owner: toAddress,
	}
}