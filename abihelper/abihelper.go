package abihelper

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func AbiDecode(abiJson, data string) (resArray []interface{}, err error) {
	abiObj, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		log.Println("in AbiDecode abiJson err:", err.Error())
		return nil, err
	}
	// decode txInput method signature
	dataByte := common.FromHex(data)
	// recover Method from signature and ABI
	method, err := abiObj.MethodById(dataByte)
	if err != nil {
		return nil, err
	}
	// unpack method inputs
	resArray, err = method.Inputs.Unpack(dataByte[4:])
	if err != nil {
		log.Println("in AbiDecode Unpack err:", err.Error())
		return nil, err
	}
	return resArray, nil
}

func AbiEncode(abiJson, method string, args ...interface{}) (string, error) {
	abiObj, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		log.Println("in AbiEncode abiJson err:", err.Error())
		return "", err
	}
	data, err := abiObj.Pack(method, args)
	if err != nil {
		log.Println("in AbiEncode pack err:", err.Error())
		return "", err
	}
	return common.Bytes2Hex(data), nil
}