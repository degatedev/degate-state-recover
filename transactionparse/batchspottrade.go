package transactionparse

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type BatchBalance struct {
	firstTokenPositiveAmount decimal.Decimal
	secondTokenPositiveAmount decimal.Decimal
	thirdTokenPositiveAmount decimal.Decimal

	firstTokenNegativeAmount decimal.Decimal
	secondTokenNegativeAmount decimal.Decimal
	thirdTokenNegativeAmount decimal.Decimal
}

// - txType: 3bits
// - bindToken: 5bits
// - FirstToken: 32bits
// - SecondToken: 32bits
// - User1TokenType: 2bits
// - User2TokenType: 2bits
// - User3TokenType: 2bits
// - User4TokenType: 2bits
// - User5TokenType: 2bits

// - User1AccountID: 32bits
// - User1FirstTokenAmountExchange: 30bits
// - User1SecondTokenAmountExchange: 30bits

// - User2AccountID: 32bits
// - User2FirstTokenAmountExchange: 30bits
// - User2SecondTokenAmountExchange: 30bits

// - User3AccountID: 32bits
// - User3FirstTokenAmountExchange: 30bits
// - User3SecondTokenAmountExchange: 30bits

// - User4AccountID: 32bits
// - User4FirstTokenAmountExchange: 30bits
// - User4SecondTokenAmountExchange: 30bits

// - User5AccountID: 32bits
// - User5FirstTokenAmountExchange: 30bits
// - User5SecondTokenAmountExchange: 30bits

// - User0AccountID: 32bits
// - User0FirstTokenAmountExchange: 30bits
// - User0SecondTokenAmountExchange: 30bits
// - User0ThirdTokenAmountExchange: 30bits
func parseBatchSpotTrade(index int, batchSpotTradeData []byte, accountUpdatePack *entity.AccountUpdatePack, operatorAccountUpdate *entity.AccountUpdate, state *entity.State) {
	batchSpotTradeDataBits := util.HexToBitsStr(common.Bytes2Hex(batchSpotTradeData))
	// txType := batchSpotTradeDataBits[:3]
	bindToken := util.BitsStrToIntStr(batchSpotTradeDataBits[3:8])
	firstToken := util.BitsStrToIntStr(batchSpotTradeDataBits[8:40])
	secondToken := util.BitsStrToIntStr(batchSpotTradeDataBits[40:72])
	userOneTokenType := batchSpotTradeDataBits[72:74]
	userTwoTokenType := batchSpotTradeDataBits[74:76]
	userThreeTokenType := batchSpotTradeDataBits[76:78]
	userFourTokenType := batchSpotTradeDataBits[78:80]
	userFiveTokenType := batchSpotTradeDataBits[80:82]

	userOneAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[82:114])
	userOneFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[114:144], 30)
	userOneSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[144:174], 30)

	userTwoAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[174:206])
	userTwoFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[206:236], 30)
	userTwoSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[236:266], 30)

	userThreeAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[266:298])
	userThreeFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[298:328], 30)
	userThreeSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[328:358], 30)

	userFourAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[358:390])
	userFourFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[390:420], 30)
	userFourSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[420:450], 30)

	userFiveAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[450:482])
	userFiveFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[482:512], 30)
	userFiveSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[512:542], 30)

	userZeroAccountID := util.BitsStrToIntStr(batchSpotTradeDataBits[542:574])
	userZeroFirstTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[574:604], 30)
	userZeroSecondTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[604:634], 30)
	userZeroThirdTokenAmountExchange := util.BitsStrToFixedBits(batchSpotTradeDataBits[634:664], 30)

	// 1: update user0
	// 2: update user1
	// 3: update user2
	// 4: update user3
	// 5: update user4
	// 6: update user5
	// 7: update operator

	batchBalance := &BatchBalance{decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0), decimal.NewFromInt(0)}

	accountZeroUpdate := getAccountUpdate(userZeroAccountID, accountUpdatePack, state)
	updateZeroAccountBalance(accountZeroUpdate, userZeroFirstTokenAmountExchange, userZeroSecondTokenAmountExchange, userZeroThirdTokenAmountExchange, firstToken, secondToken, bindToken, batchBalance)
	
	accountOneUpdate := getAccountUpdate(userOneAccountID, accountUpdatePack, state)
	updateBatchAccountBalanceByTokenType(accountOneUpdate, userOneFirstTokenAmountExchange, userOneSecondTokenAmountExchange, userOneTokenType, firstToken, secondToken, bindToken, batchBalance)

	accountTwoUpdate := getAccountUpdate(userTwoAccountID, accountUpdatePack, state)
	updateBatchAccountBalanceByTokenType(accountTwoUpdate, userTwoFirstTokenAmountExchange, userTwoSecondTokenAmountExchange, userTwoTokenType, firstToken, secondToken, bindToken, batchBalance)

	accountThreeUpdate := getAccountUpdate(userThreeAccountID, accountUpdatePack, state)
	updateBatchAccountBalanceByTokenType(accountThreeUpdate, userThreeFirstTokenAmountExchange, userThreeSecondTokenAmountExchange, userThreeTokenType, firstToken, secondToken, bindToken, batchBalance)

	accountFourUpdate := getAccountUpdate(userFourAccountID, accountUpdatePack, state)
	updateBatchAccountBalanceByTokenType(accountFourUpdate, userFourFirstTokenAmountExchange, userFourSecondTokenAmountExchange, userFourTokenType, firstToken, secondToken, bindToken, batchBalance)

	accountFiveUpdate := getAccountUpdate(userFiveAccountID, accountUpdatePack, state)
	updateBatchAccountBalanceByTokenType(accountFiveUpdate, userFiveFirstTokenAmountExchange, userFiveSecondTokenAmountExchange, userFiveTokenType, firstToken, secondToken, bindToken, batchBalance)

	// calculate operator
	operatorFirstAmount := batchBalance.firstTokenNegativeAmount.Sub(batchBalance.firstTokenPositiveAmount)
	operatorSecondAmount := batchBalance.secondTokenNegativeAmount.Sub(batchBalance.secondTokenPositiveAmount)
	operatorThirdAmount := batchBalance.thirdTokenNegativeAmount.Sub(batchBalance.thirdTokenPositiveAmount)

	if operatorFirstAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		panic("in parseBatchSpotTrade operatorFirstAmount < 0:" + operatorFirstAmount.String())
	}
	if operatorSecondAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		panic("in parseBatchSpotTrade operatorSecondAmount < 0:" + operatorSecondAmount.String())
	}
	if operatorThirdAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		panic("in parseBatchSpotTrade operatorThirdAmount < 0:" + operatorThirdAmount.String())
	}
	statemanager.AddBalance(operatorAccountUpdate, operatorFirstAmount.String(), firstToken)
	statemanager.AddBalance(operatorAccountUpdate, operatorSecondAmount.String(), secondToken)
	statemanager.AddBalance(operatorAccountUpdate, operatorThirdAmount.String(), bindToken)

	statemanager.UpdateAccountAsset(accountZeroUpdate)
	statemanager.UpdateAccountAsset(accountOneUpdate)
	statemanager.UpdateAccountAsset(accountTwoUpdate)
	statemanager.UpdateAccountAsset(accountThreeUpdate)
	statemanager.UpdateAccountAsset(accountFourUpdate)
	statemanager.UpdateAccountAsset(accountFiveUpdate)
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
}

func updateZeroAccountBalance(accountUpdate *entity.AccountUpdate, firstAmountBits, secondAmountBits, thirdAmountBits []int, firstToken, secondToken, thirdToken string, batchBalance *BatchBalance) {
	firstPositive, firstNegative := updateBatchAccountBalance(accountUpdate, firstAmountBits, firstToken)
	secondPositive, secondNegative := updateBatchAccountBalance(accountUpdate, secondAmountBits, secondToken)
	thirdPositive, thirdNegative := updateBatchAccountBalance(accountUpdate, thirdAmountBits, thirdToken)

	batchBalance.firstTokenPositiveAmount = batchBalance.firstTokenPositiveAmount.Add(firstPositive)
	batchBalance.firstTokenNegativeAmount = batchBalance.firstTokenNegativeAmount.Add(firstNegative)

	batchBalance.secondTokenPositiveAmount = batchBalance.secondTokenPositiveAmount.Add(secondPositive)
	batchBalance.secondTokenNegativeAmount = batchBalance.secondTokenNegativeAmount.Add(secondNegative)

	batchBalance.thirdTokenPositiveAmount = batchBalance.thirdTokenPositiveAmount.Add(thirdPositive)
	batchBalance.thirdTokenNegativeAmount = batchBalance.thirdTokenNegativeAmount.Add(thirdNegative)
}

func updateBatchAccountBalanceByTokenType(accountUpdate *entity.AccountUpdate, amountBitsA, amountBitsB []int, tokenType, firstToken, secondToken, thirdToken string, batchBalance *BatchBalance) {
	firstPositive := decimal.NewFromInt(0)
	firstNegative := decimal.NewFromInt(0)
	secondPositive := decimal.NewFromInt(0)
	secondNegative := decimal.NewFromInt(0)
	thirdPositive := decimal.NewFromInt(0)
	thirdNegative := decimal.NewFromInt(0)
	if tokenType == "00" {
		firstPositive, firstNegative = updateBatchAccountBalance(accountUpdate, amountBitsA, firstToken)
		secondPositive, secondNegative = updateBatchAccountBalance(accountUpdate, amountBitsB, secondToken)
	} else if tokenType == "01" {
		firstPositive, firstNegative = updateBatchAccountBalance(accountUpdate, amountBitsA, firstToken)
		thirdPositive, thirdNegative = updateBatchAccountBalance(accountUpdate, amountBitsB, thirdToken)
	} else if tokenType == "10" {
		secondPositive, secondNegative = updateBatchAccountBalance(accountUpdate, amountBitsA, secondToken)
		thirdPositive, thirdNegative = updateBatchAccountBalance(accountUpdate, amountBitsB, thirdToken)
	} else {
		panic("in updateBatchAccountBalanceByTokenType invalid tokenType")
	}
	batchBalance.firstTokenPositiveAmount = batchBalance.firstTokenPositiveAmount.Add(firstPositive)
	batchBalance.firstTokenNegativeAmount = batchBalance.firstTokenNegativeAmount.Add(firstNegative)

	batchBalance.secondTokenPositiveAmount = batchBalance.secondTokenPositiveAmount.Add(secondPositive)
	batchBalance.secondTokenNegativeAmount = batchBalance.secondTokenNegativeAmount.Add(secondNegative)

	batchBalance.thirdTokenPositiveAmount = batchBalance.thirdTokenPositiveAmount.Add(thirdPositive)
	batchBalance.thirdTokenNegativeAmount = batchBalance.thirdTokenNegativeAmount.Add(thirdNegative)
}

func updateBatchAccountBalance(accountUpdate *entity.AccountUpdate, amountBits []int, tokenID string) (positive, negative decimal.Decimal) {
	amountDecimal, isNeg := util.ParseFloat30(amountBits)
	if isNeg {
		statemanager.SubBalance(accountUpdate, amountDecimal.String(), tokenID)
		return decimal.NewFromInt(0), amountDecimal
	} else {
		statemanager.AddBalance(accountUpdate, amountDecimal.String(), tokenID)
		return amountDecimal, decimal.NewFromInt(0)
	}
}