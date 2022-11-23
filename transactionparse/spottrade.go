package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// - txType: 3bits
// - txTypePad: 1bits
// - For both Orders:
//     - Account ID: 4 bytes
//     - TokenS: 4 bytes
//     - FillS: 4 bytes （24 bits, 19 bits for the mantissa part and 5 for the exponent part)
//     - FeeTokenID: 4 bytes
//     - Fee: 2 bytes （16 bits, 11 bits for the mantissa part and 5 for the exponent part)
//     - Order data (fillAmountBorS,feeBips): 1 byte
func parseSpotTrade(index int, spotTradeData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	spotTradeDataBits := util.HexToBitsStr(common.Bytes2Hex(spotTradeData))
	// txType := spotTradeDataBits[:3]
	orderAAccountID := util.BitsStrToIntStr(spotTradeDataBits[4:36])
	orderBAccountID := util.BitsStrToIntStr(spotTradeDataBits[36:68])
	orderATokenSID := util.BitsStrToIntStr(spotTradeDataBits[68:100])
	orderBTokenSID := util.BitsStrToIntStr(spotTradeDataBits[100:132])
	orderAFilledS := util.BitsStrToFixedBits(spotTradeDataBits[132:164], 32)
	orderBFilledS := util.BitsStrToFixedBits(spotTradeDataBits[164:196], 32)

	orderAFeeTokenID := util.BitsStrToIntStr(spotTradeDataBits[196:228])
	orderAFeeAmount := util.ParseFloat(util.BitsStrToFixedBits(spotTradeDataBits[228:244], 16), floatEncoding16)
	orderBFeeTokenID := util.BitsStrToIntStr(spotTradeDataBits[244:276])
	orderBFeeAmount := util.ParseFloat(util.BitsStrToFixedBits(spotTradeDataBits[276:292], 16), floatEncoding16)

	// orderAFillAmountBorS := spotTradeDataBits[292:293]
	orderAFeeBipsFlag := spotTradeDataBits[293:294]
	orderAFeeBipsData := spotTradeDataBits[294:300]
	// orderBFillAmountBorS := spotTradeDataBits[300:301]
	orderBFeeBipsFlag := spotTradeDataBits[301:302]
	orderBFeeBipsData := spotTradeDataBits[302:308]


	orderAFilledSDecimal := util.ParseFloat(orderAFilledS, floatEncoding32)
	orderBFilledSDecimal := util.ParseFloat(orderBFilledS, floatEncoding32)

	// trading fee
	orderATradingFee := calculateTradingFee(orderAFeeBipsData, orderAFeeBipsFlag, orderBFilledSDecimal)
	orderBTradingFee := calculateTradingFee(orderBFeeBipsData, orderBFeeBipsFlag, orderAFilledSDecimal)

	accountAUpdate := getAccountUpdate(orderAAccountID, accountUpdatePack, state)
	statemanager.AddBalance(accountAUpdate, orderBFilledSDecimal.String(), orderBTokenSID)
	statemanager.SubBalance(accountAUpdate, orderATradingFee.String(), orderBTokenSID)
	statemanager.SubBalance(accountAUpdate, orderAFilledSDecimal.String(), orderATokenSID)
	statemanager.SubBalance(accountAUpdate, orderAFeeAmount.String(), orderAFeeTokenID)

	accountBUpdate := getAccountUpdate(orderBAccountID, accountUpdatePack, state)
	statemanager.AddBalance(accountBUpdate, orderAFilledSDecimal.String(), orderATokenSID)
	statemanager.SubBalance(accountBUpdate, orderBTradingFee.String(), orderATokenSID)
	statemanager.SubBalance(accountBUpdate, orderBFilledSDecimal.String(), orderBTokenSID)
	statemanager.SubBalance(accountBUpdate, orderBFeeAmount.String(), orderBFeeTokenID)

	statemanager.AddBalance(operatorAccountUpdate, orderAFeeAmount.String(), orderAFeeTokenID)
	statemanager.AddBalance(operatorAccountUpdate, orderBFeeAmount.String(), orderBFeeTokenID)

	statemanager.AddBalance(operatorAccountUpdate, orderATradingFee.String(), orderBTokenSID)
	statemanager.AddBalance(operatorAccountUpdate, orderBTradingFee.String(), orderATokenSID)


	statemanager.UpdateAccountAsset(accountAUpdate)
	statemanager.UpdateAccountAsset(accountBUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}

func calculateTradingFee(feeBipsData, feeBipsFlag string, filledDecimal decimal.Decimal) decimal.Decimal {
	feeBipsInt := int64(0)
	if feeBipsFlag == "1" {
		feeBipsInt = int64(50) * util.BitsStrToInt(feeBipsData)
	} else {
		feeBipsInt = util.BitsStrToInt(feeBipsData)
	}
	feeBipsDecimal := decimal.NewFromInt(feeBipsInt)
	result := filledDecimal.Mul(feeBipsDecimal).Div(decimal.NewFromInt(10000))
	return decimal.NewFromBigInt(result.BigInt(), 0)
}