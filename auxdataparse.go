package staterecovery

import (
	"encoding/json"
	"log"
	"staterecovery/common/entity"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func parseAuxData(blockAuxData []byte, depositSize, accountUpdateSize, withdrawSize int) []*entity.WithdrawAuxData  {
	transactionAuxData := getTransactionAuxData(blockAuxData)
	withdrawAuxDataList := make([]*entity.WithdrawAuxData, 0)
	for index, transaction := range transactionAuxData {
		if index < depositSize {
			// deposit
		} else if index < depositSize + accountUpdateSize {
			// account update
		} else if index < depositSize + accountUpdateSize + withdrawSize {
			// withdraw
			withdrawAuxDataList = append(withdrawAuxDataList, parseWithdrawAux(transaction.Data))
		}
	}
	return withdrawAuxDataList
}

func getTransactionAuxData(data []byte) (aux []entity.Auxiliary) {
	var arguments []abi.Argument
	t, err := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "data", Type: "bytes"},
	})
	if err != nil {
		log.Println("in parseAuxiliaryData abi NewType error:", err.Error())
		return
	}
	arguments = append(arguments, abi.Argument{
		Name: "data",
		Type: t,
	})
	input := abi.Arguments(arguments)
	unpacked, err := input.Unpack(data)
	if err != nil {
		log.Println("in parseAuxiliaryData input Unpackerror:", err.Error())
		return
	}
	if len(unpacked) == 0 {
		log.Println("in parseAuxiliaryData unpacked length is 0")
		return
	}
	ad := &entity.AuxiliaryData{}
	err = input.Copy(ad, unpacked)
	if err != nil {
		log.Println("in parseAuxiliaryData input Copy:", err.Error())
		return
	}
	return ad.Data
}

func parseWithdrawAux(data []byte) (withdrawAuxData *entity.WithdrawAuxData) {
	if data == nil || len(data) == 0 {
		log.Println("in parseWithdrawAux invalid data")
	   return
	}

	// log.Println("in parseWithdrawAux data hex:", common.Bytes2Hex(data))
	var arguments []abi.Argument
	var t abi.Type
 
	t, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
	   {Name: "storeRecipient", Type: "bool"},
	   {Name: "gasLimit", Type: "uint256"},
	   {Name: "signature", Type: "bytes"},
	   {Name: "minGas", Type: "uint248"},
	   {Name: "to", Type: "address"},
	   {Name: "maxFee", Type: "uint96"},
	   {Name: "validUntil", Type: "uint32"},
	   {Name: "amount", Type: "uint248"},
	})
	if err != nil {
	   log.Println("in parseWithdrawAux abi NewType error:", err.Error())
	   return
	}
	arguments = append(arguments, abi.Argument{
	   Name: "data",
	   Type: t,
	})
 
	input := abi.Arguments(arguments)
	unpacked, err := input.Unpack(data)
	if err != nil {
	   log.Println("in parseWithdrawAux input Unpack error:", err.Error())
	   return
	}
	if len(unpacked) == 0 {
	   log.Println("in parseWithdrawAux unpacked length is 0")
	   return
	}
	unpackedJson, err := json.Marshal(unpacked)
	if len(unpacked) == 0 {
	   log.Println("in parseWithdrawAux json Marshal error")
	   return
	}
	withdrawAuxDataList := make([]*entity.WithdrawAuxData, 0)
	err = json.Unmarshal(unpackedJson, &withdrawAuxDataList)
	if err != nil {
	   log.Println("in parseWithdrawAux input Copy error:", err.Error())
	   return
	}
	return withdrawAuxDataList[0]
 }