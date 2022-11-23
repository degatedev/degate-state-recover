package util

import (
	"encoding/binary"
	"errors"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var (
	ErrorToAnyHexadecimal = errors.New("in ToAnyHexadecimal error")
)

func BytesToInt(buf []byte) int {
	return int(binary.BigEndian.Uint16(buf))
}
func BytesToUInt32(buf []byte) uint64 {
	return uint64(binary.BigEndian.Uint32(buf))
}

func HexToUInt32(dataHex string) uint64 {
	dataBytes := common.FromHex(dataHex)
	return BytesToUInt32(dataBytes)
}

func Bytes2Bits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	return dst
}

func BitsToStr(data []int) string {
	if data == nil || len(data) == 0 {
		return ""
	}
	result := ""
	for _, item := range data {
		if item != 1 && item != 0 {
			panic("in BitsToStr data is not bits")
		}
		result += strconv.Itoa(item)
	}
	return result
}

func SupplyZero(length int, str string) string {
	if len(str) < length {
		strLength := len(str)
		for i := 0; i < length-strLength; i++ {
			str = "0" + str
		}
	}
	return str
}
func ToAnyHexadecimal(numberStr string, currentHexadecimal, toHexadecimal, toLength int) (string, error) {
	bigCurrent, ok := new(big.Int).SetString(numberStr, currentHexadecimal)
	if !ok {
		log.Println("in ToAnyHexadecimal SetString:", ok)
		return "", ErrorToAnyHexadecimal
	}
	toNumberStr := bigCurrent.Text(toHexadecimal)
	if toLength != 0 {
		toNumberStr = SupplyZero(toLength, toNumberStr)
	}
	return toNumberStr, nil
}

func BitsToHex(data []int) string {
	dataStr := BitsToStr(data)
	log.Println("in bits to hex:", dataStr)
	dataHex, err := ToAnyHexadecimal(dataStr, 2, 16, len(data) / 8)
	if err != nil {
		panic("in BitsToHex ToAnyHexadecimal error:" + err.Error())
	}
	return dataHex
}

func BitsToBytes(data []int) []byte {
	dataHex := BitsToHex(data)
	return common.FromHex(dataHex)
}

func BitsToIntStr(data []int) string {
	dataStr := BitsToStr(data)
	dataIntStr, err := ToAnyHexadecimal(dataStr, 2, 10, 0)
	if err != nil {
		panic("in BitsToIntStr ToAnyHexadecimal error:" + err.Error())
	}
	return dataIntStr
}

func BitsToDecimal(data []int) decimal.Decimal {
	dataIntStr := BitsToIntStr(data)
	return NewDecimalFromString(dataIntStr)
}

func HexToBitsStr(dataHex string) string {
	dataBytes := common.FromHex(dataHex)
	dataBits := Bytes2Bits(dataBytes)
	dataBitsStr := ""
	for _, bitInt := range dataBits {
		dataBitsStr += strconv.Itoa(bitInt)
	}
	return dataBitsStr
}

func HexToIntStr(dataHex string) string {
	if len(dataHex) >=2 && dataHex[:2] == "0x" {
		dataHex = dataHex[2:]
	}
	intStr, err := ToAnyHexadecimal(dataHex, 16, 10, 0)
	if err != nil {
		panic("in HexToIntStr ToAnyHexadecimal error:" + err.Error())
	}
	return intStr
}

func IntStrToHex(dataIntStr string) string {
	dataHex, err := ToAnyHexadecimal(dataIntStr, 10, 16, 0)
	if err != nil {
		panic("in HexToIntStr ToAnyHexadecimal error:" + err.Error())
	}
	return dataHex
}

func HexToBits(dataHex string) []int {
	dataBytes := common.FromHex(dataHex)
	return Bytes2Bits(dataBytes)
}

func BytesToIntStr(data []byte) string {
	dataHex := common.Bytes2Hex(data)
	return HexToIntStr(dataHex)
}

func IntStrToBitStr(dataStr string) string {
	bitStr, err := ToAnyHexadecimal(dataStr, 10, 2, 0)
	if err != nil {
		panic("in IntStrToBitStr ToAnyHexadecimal error:" + err.Error())
	}
	return bitStr
}

func BitsStrToHex(bitsStr string) string {
	dataHex, err := ToAnyHexadecimal(bitsStr, 2, 16, len(bitsStr) / 8)
	if err != nil {
		panic("in BitsStrToHex ToAnyHexadecimal error:" + err.Error())
	}
	return dataHex
}

func BitsStrToIntStr(bitsStr string) string {
	dataHex, err := ToAnyHexadecimal(bitsStr, 2, 10, 0)
	if err != nil {
		panic("in BitsStrToIntStr ToAnyHexadecimal error:" + err.Error())
	}
	return dataHex
}

func BitsStrToInt(bitsStr string) int64 {
	dataIntStr := BitsStrToIntStr(bitsStr)
	dataInt, err := strconv.ParseInt(dataIntStr, 10, 64)
	if err != nil {
		panic("in BitsStrToInt parse int64 error:" + err.Error())
	}
	return dataInt
}

func NewBigIntFromStr(str string) *big.Int {
	strDecimal, err := decimal.NewFromString(str)
	if err != nil {
		panic("in NewBigIntFromStr error:" + err.Error())
	}
	return strDecimal.BigInt()
}

func ConvertStringArrayToBigIntArray(strArray []string) []*big.Int {
	bigIntArray := make([]*big.Int, 0)
	if strArray == nil || len(strArray) == 0 {
		return bigIntArray
	}
	for _, str := range strArray {
		bigIntArray = append(bigIntArray, NewBigIntFromStr(str))
	}
	return bigIntArray
}

func ConvertDecimalArrayToBigIntArray(strArray []decimal.Decimal) []*big.Int {
	bigIntArray := make([]*big.Int, 0)
	if strArray == nil || len(strArray) == 0 {
		return bigIntArray
	}
	for _, str := range strArray {
		bigIntArray = append(bigIntArray, str.BigInt())
	}
	return bigIntArray
}

func NewDecimalFromString(str string) decimal.Decimal {
	strDecimal, err := decimal.NewFromString(str)
	if err != nil {
		panic("in NewDecimalFromString err:" + err.Error())
	}
	return strDecimal
}

func NewBigIntFromString(str string) *big.Int {
	strDecimal := NewDecimalFromString(str)
	return strDecimal.BigInt()
}

func IntStrToFixedBits(dataStr string, bitsNum int) ([]int) {
	dataBitsStr := IntStrToBitStr(dataStr)
	if len(dataBitsStr) > bitsNum {
		panic("in IntStrToFixedBits invalid dataStr or bitsNum dataStr:" + dataStr + ";bitsNum:" + strconv.Itoa(bitsNum))
	}
	if len(dataBitsStr) < bitsNum {
		dataBitsStr = SupplyZero(bitsNum, dataBitsStr)
	}
	dataBits := make([]int, bitsNum)
	var err error
	for i := 0; i < bitsNum; i++ {
		dataBits[i], err = strconv.Atoi(string(dataBitsStr[i]))
		if err != nil {
			panic("in IntStrToFixedBits invalid dataBitsStr")
		}
	}
	return dataBits
}

func BitsStrToFixedBits(bitsStr string, bitsNum int) []int {
	dataBits := make([]int, bitsNum)
	var err error
	for i := 0; i < bitsNum; i++ {
		dataBits[i], err = strconv.Atoi(string(bitsStr[i]))
		if err != nil {
			panic("in BitsStrToFixedBits invalid bitsStr")
		}
	}
	return dataBits
}