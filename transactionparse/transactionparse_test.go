package transactionparse

import (
	"log"
	"staterecovery/common/util"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBitsToHex(t *testing.T) {
	dataHex := "1523a34"
	dataBits := util.Bytes2Bits(common.FromHex(dataHex))
	dataHexTwo := util.BitsToHex(dataBits)
	log.Println("dataHexTwo:", dataHexTwo)

	bitsStr := util.HexToBitsStr(dataHex)
	log.Println("bitsStr:", bitsStr)
	dataHexThree := util.BitsStrToHex(bitsStr)
	log.Println("dataHexThree:", dataHexThree)

	bitsStr = "10010"
	log.Println("bitsHex:", util.BitsStrToHex(bitsStr))
}

func TestSplitArray(t *testing.T) {
	testArray := make([]int, 5)
	testArray[0] = 0
	testArray[1] = 1
	testArray[2] = 2
	testArray[3] = 3
	testArray[4] = 4
	depositSize := 1
	accountUpdateSize := 1
	withdrawSize := 1

	depositArray := testArray[:depositSize]
	accountUpdateArray := testArray[depositSize : depositSize+accountUpdateSize]
	otherArray := testArray[depositSize+accountUpdateSize : len(testArray)-withdrawSize]
	withdrawArray := testArray[len(testArray)-withdrawSize:]

	for _, dp := range depositArray {
		log.Println("deposit:", strconv.Itoa(dp))
	}
	for _, dp := range accountUpdateArray {
		log.Println("accountUpdateArray:", strconv.Itoa(dp))
	}
	for _, dp := range otherArray {
		log.Println("otherArray:", strconv.Itoa(dp))
	}
	for _, dp := range withdrawArray {
		log.Println("withdrawArray:", strconv.Itoa(dp))
	}
}

func TestPublicKey(t *testing.T) {
	compressedKey := "2160ceb725341c107a1cb4214ee5201a094dcd8820d6345425569a27f76ba3bf"
	x, y := decompressPublicKey(compressedKey)
	log.Println("x:", x.String())
	log.Println("y:", y.String())
}
