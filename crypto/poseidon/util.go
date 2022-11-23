package poseidon

import "math/big"

// NewElement returns a new empty *Element
func NewElement() *Element {
	return &Element{}
}

// BigIntArrayToElementArray converts an array of *big.Int into an array of *ff.Element
func BigIntArrayToElementArray(bi []*big.Int) []*Element {
	var o []*Element
	for i := range bi {
		o = append(o, NewElement().SetBigInt(bi[i]))
	}
	return o
}
