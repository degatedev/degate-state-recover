package transactionparse

import (
	"math/big"
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/crypto/babyjub"
	"staterecovery/statemanager"

	"github.com/ethereum/go-ethereum/common"
)

var (
	floatEncoding32 = &util.FloatEncoding{NumBitsExponent: 7, NumBitsMantissa: 25, ExponentBase: 10}
	floatEncoding29 = &util.FloatEncoding{NumBitsExponent: 5, NumBitsMantissa: 24, ExponentBase: 10}
	floatEncoding16 = &util.FloatEncoding{NumBitsExponent: 5, NumBitsMantissa: 11, ExponentBase: 10}
)

func ParseTransaction(blockSize, depositSize, accountUpdateSize, withdrawSize int, transactionData []byte, withdrawAuxDataList []*entity.WithdrawAuxData, operatorAccountID string, state *entity.State) {
	// transform
	// transaction bytes = 80 + 3
	// transactionData = 80 * blockSize + 3 * blockSize
	transactionDatas := make([][]byte, blockSize)
	firstPartDataSize := 80 * blockSize
	for i := 0; i < blockSize; i++ {
		transactionDatas[i] = append(transactionDatas[i], transactionData[80 * i: 80 * (i + 1)]...)
		transactionDatas[i] = append(transactionDatas[i], transactionData[firstPartDataSize + 3 * i: firstPartDataSize + 3 * (i + 1)]...)
	}
	accountUpdatePack := &entity.AccountUpdatePack{
		AccountUpdateList: make([]*entity.AccountUpdate, 0),
	}
	operatorAccountUpdate := getAccountUpdate(operatorAccountID, accountUpdatePack, state)
	// handle deposit
	ParseDeposits(transactionDatas[:depositSize], accountUpdatePack, state)
	// handle accountUpdate
	transferToAccountList := ParseTransferToAccount(transactionDatas[depositSize+accountUpdateSize:len(transactionDatas) - withdrawSize])
	ParseAccountUpdates(transactionDatas[depositSize:depositSize+accountUpdateSize], accountUpdatePack, operatorAccountUpdate, state, transferToAccountList)
	// handle other transaction
	ParseOtherTransactions(transactionDatas[depositSize+accountUpdateSize:len(transactionDatas) - withdrawSize], accountUpdatePack, operatorAccountUpdate, state)
	// handle withdraw
	ParseWithdraws(transactionDatas[len(transactionDatas)-withdrawSize:], accountUpdatePack, withdrawAuxDataList, operatorAccountUpdate, state)

	// update operator nonce 
	operatorAccountUpdate.AccountValue.Nonce++
	statemanager.UpdateAccountAsset(operatorAccountUpdate)
	statemanager.UpdateState(accountUpdatePack.AccountUpdateList, state)
}

func decompressPublicKey(compressedKey string) (publicKeyX, publicKeyY *big.Int) {
	if compressedKey == "" || compressedKey == "0" {
		return big.NewInt(0), big.NewInt(0)
	}
	// 32bytes
	for len(compressedKey) < 64 {
		compressedKey = "0" + compressedKey
	}
	if compressedKey == "0000000000000000000000000000000000000000000000000000000000000000" {
		return big.NewInt(0), big.NewInt(0)
	}
	compressedKeyBytes := common.FromHex(compressedKey)
	reverse := [32]byte{}
	for i := 0; i < 32; i ++ {
		reverse[31 - i] = compressedKeyBytes[i]
	}
	point := &babyjub.Point{}
	decompressedPoint, err := point.Decompress(reverse)
	if err != nil {
		panic("in DecompressPublicKey error:" + err.Error())
	}
	return decompressedPoint.X, decompressedPoint.Y
}

func getAccountUpdate(accountID string, accountUpdatePack *entity.AccountUpdatePack, state *entity.State) (selectAccountUpdate *entity.AccountUpdate) {
	for _, accountUpdate := range accountUpdatePack.AccountUpdateList {
		if accountID == accountUpdate.AccountID {
			return accountUpdate
		}
	}
	selectAccountUpdate = statemanager.GetAccount(state, accountID)
	accountUpdatePack.AccountUpdateList = append(accountUpdatePack.AccountUpdateList, selectAccountUpdate)
	return selectAccountUpdate
}