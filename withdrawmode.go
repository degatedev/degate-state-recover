package staterecovery

import (
	"errors"
	"log"
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"staterecovery/conf"
	"staterecovery/statemanager"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func createWithdrawModeMerkleProof(config *conf.Config, state *entity.State) {
	if config.WithdrawModeAccount == "" || config.WithdrawModeToken == "" || config.WithdrawModeFilePath == "" {
		log.Println("in createWithdrawModeMerkleProof invalid withdraw mode param")
		return
	}
	// 获取对应的accountID与tokenID
	withdrawModeAccountID, err := findAccountID(config.WithdrawModeAccount, state)
	if err != nil {
		log.Println("in createWithdrawModeMerkleProof findAccountID err:", err.Error)
		return
	}
	withdrawModeTokenID, err := dataResource.GetTokenIDByAddress(config.WithdrawModeToken)
	if err != nil {
		log.Println("in createWithdrawModeMerkleProof GetTokenIDByAddress err:", err.Error)
		return
	}

	log.Println("in createWithdrawModeMerkleProof withdrawModeTokenID:", withdrawModeTokenID)
	statemanager.CalculateStateTree(state)
	merkleProof := statemanager.CreateWithdrawModeMerkleProof(state, withdrawModeAccountID, withdrawModeTokenID)
	merkleInput := convertMerkleProofToInput(merkleProof)
	// encode to binary
	merkleProofBinary, err := withdrawModeDataEncodeToBinary(merkleProof)
	if err != nil {
		log.Println("in createWithdrawModeMerkleProof withdrawModeDataEncodeToBinary error:", err.Error())
		return
	}
	log.Println("in createWithdrawModeMerkleProof merkleProofBinary:", merkleProofBinary)
	// merkleProofJson, err := json.Marshal(merkleProof)
	// if err != nil {
	// 	log.Println("in createWithdrawModeMerkleProof json Marshal error:", err.Error())
	// 	return
	// }

	log.Println("in createWithdrawModeMerkleProof merkleProofJson:", merkleInput)
	err = util.WriteTxtFile(merkleInput, config.WithdrawModeFilePath)
	if err != nil {
		log.Println("in createWithdrawModeMerkleProof WriteTxtFile error:", err.Error())
		return
	}
}

func findAccountID(accountAddr string, state *entity.State) (accountID string, err error) {
	accountAddrIntStr := util.HexToIntStr(accountAddr)
	log.Println("in findAccountID accountAddrIntStr:", accountAddrIntStr)
	for accountID, account := range state.AccountsValues {
		if accountAddrIntStr == account.Owner {
			return accountID, nil
		}
	}
	return "0", errors.New("in findAccountID can not find accountID by address:" + accountAddr)
}

func convertMerkleProofToInput(withdrawModeMerkleProof *entity.WithdrawModeMerkleProof) string {
	input := "["
	// accountLeaf
	input += "["
	input += "\"" + strconv.FormatUint(uint64(withdrawModeMerkleProof.AccountLeaf.AccountID), 10) + "\","
	input += "\"" + withdrawModeMerkleProof.AccountLeaf.Owner.String() + "\","
	input += "\"" + withdrawModeMerkleProof.AccountLeaf.PubKeyX.String() + "\","
	input += "\"" + withdrawModeMerkleProof.AccountLeaf.PubKeyY.String() + "\","
	input += strconv.FormatUint(uint64(withdrawModeMerkleProof.AccountLeaf.Nonce), 10)
	input += "],"
	// balanceLeaf
	input += "["
	input += "\"" + strconv.FormatUint(uint64(withdrawModeMerkleProof.BalanceLeaf.TokenID), 10) + "\","
	input += "\"" + withdrawModeMerkleProof.BalanceLeaf.Balance.String() + "\""
	input += "],"
	// accountMerkleProof
	input += "["
	for _, proof := range withdrawModeMerkleProof.AccountMerkleProof {
		input += "\"" + proof.String() + "\","
	}
	input = input[:len(input)-1]
	input += "],"
	// balanceMerkleProof
	input += "["
	for _, proof := range withdrawModeMerkleProof.BalanceMerkleProof {
		input += "\"" + proof.String() + "\","
	}
	input = input[:len(input)-1]
	input += "]"
	input += "]"
	return input
}

func withdrawModeDataEncodeToBinary(withdrawModeMerkleProof *entity.WithdrawModeMerkleProof) (string, error) {
	abiObj, err := abi.JSON(strings.NewReader(entity.WithdrawModeAbi))
	if err != nil {
		log.Println("in AbiEncode abiJson err:", err.Error())
		return "", err
	}
	data, err := abiObj.Pack(entity.WithdrawModeMethod, withdrawModeMerkleProof)
	if err != nil {
		log.Println("in AbiEncode pack err:", err.Error())
		return "", err
	}
	return common.Bytes2Hex(data), nil
}