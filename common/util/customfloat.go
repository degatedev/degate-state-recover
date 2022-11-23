package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type FloatEncoding struct {
	NumBitsExponent int
	NumBitsMantissa int
	ExponentBase int
}

func ParseFloat(data []int, encoding *FloatEncoding) (result decimal.Decimal) {
	// verify length
	if len(data) != (encoding.NumBitsExponent + encoding.NumBitsMantissa) {
		panic("in parseFloat invalid data length or encoding length, len(data):" + strconv.Itoa(len(data)) + ";encoding length:" + strconv.Itoa((encoding.NumBitsExponent + encoding.NumBitsMantissa)))
	}
	numBitsExponentBits := data[:encoding.NumBitsExponent]
	numBitsMantissaBits := data[encoding.NumBitsExponent:]
	numBitsExponent := BitsToDecimal(numBitsExponentBits)
	numBitsMantissa := BitsToDecimal(numBitsMantissaBits)
	exponentBase := decimal.NewFromInt(int64(encoding.ExponentBase))
	return numBitsMantissa.Mul(exponentBase.Pow(numBitsExponent))
}

func ParseFloat30(data []int) (result decimal.Decimal, IsNegative bool) {
	float29 := &FloatEncoding{5, 24, 10}
	if len(data) != 30 {
		panic("in parseFloat invalid data length, len(data):" + strconv.Itoa(len(data)))
	}
	dataIntStr := BitsToIntStr(data)
	dataDecimal := NewDecimalFromString(dataIntStr)
	twoPow30 := decimal.NewFromInt(2).Pow(decimal.NewFromInt(30))
	twoPow29 := decimal.NewFromInt(2).Pow(decimal.NewFromInt(29))
	
	if dataDecimal.Cmp(twoPow29) <= 0 {
		// dataDecimal <= twoPow29: Positive number
		return ParseFloat(data[1:], float29), false
	} 
	// dataDecimal > twoPow29: Negative number
	negativeDataDecimal := twoPow30.Sub(dataDecimal)
	return ParseFloat(IntStrToFixedBits(negativeDataDecimal.String(), 29), float29), true
}