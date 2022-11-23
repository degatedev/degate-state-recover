package poseidon

import (
	"math/big"
)


const NROUNDSF = 6 //nolint:golint

// For loopring only.
var NROUNDSP = 53

// if input is in this map, use corresponding value as NROUNDSP
var NROUNDSPMAP = map[int]int{1:52, 2:52, 3:52, 4:52, 5:52, 6: 52, 7: 52}


func zero() *Element {
	return NewElement()
}

// ark computes Add-Round Key, from the paper https://eprint.iacr.org/2019/458.pdf
func ark(state []*Element, c []*Element, it int) {
	for i := 0; i < len(state); i++ {
		state[i].Add(state[i], c[it])
	}
}

// exp5 performs x^5 mod p
// https://eprint.iacr.org/2019/458.pdf page 8
func exp5(a *Element) {
	a.Exp(*a, 5)
}

// sbox https://eprint.iacr.org/2019/458.pdf page 6
func sbox(nRoundsF, nRoundsP int, state []*Element, i int) {
	if (i < nRoundsF/2) || (i >= nRoundsF/2+nRoundsP) {
		for j := 0; j < len(state); j++ {
			exp5(state[j])
		}
	} else {
		exp5(state[0])
	}
}

// mix returns [[matrix]] * [vector]
func mix(state []*Element, newState []*Element, m [][]*Element) {
	mul := zero()
	for i := 0; i < len(state); i++ {
		newState[i].SetUint64(0)
		for j := 0; j < len(state); j++ {
			mul.Mul(m[i][j], state[j])
			newState[i].Add(newState[i], mul)
		}
	}
}

// Hash computes the Poseidon hash for the given inputs
func Hash(inpBI []*big.Int) (*big.Int, error) {
	t := len(inpBI) + 1
	// if len(inpBI) == 0 || len(inpBI) >= len(NROUNDSP)-1 {
	// 	return nil, fmt.Errorf("invalid inputs length %d, max %d", len(inpBI), len(NROUNDSP)-1)
	// }
	// if !utils.CheckBigIntArrayInField(inpBI[:]) {
	// 	return nil, errors.New("inputs values not inside Finite Field")
	// }
	inp := BigIntArrayToElementArray(inpBI[:])
	state := make([]*Element, t)
	state[t-1] = zero()
	copy(state[:t], inp[:])

	nRoundsF := NROUNDSF
	var nRoundsP = NROUNDSP
	nSp, ok := NROUNDSPMAP[t]
	if ok {
		nRoundsP = nSp
	}

	newState := make([]*Element, t)
	for i := 0; i < t; i++ {
		newState[i] = zero()
	}

	// ARK --> SBox --> M, https://eprint.iacr.org/2019/458.pdf pag.5
	for i := 0; i < nRoundsF+nRoundsP; i++ {
		ark(state, c.c, i)
		sbox(nRoundsF, nRoundsP, state, i)
		mix(state, newState, c.m[t])
		state, newState = newState, state
	}
	rE := state[0]
	r := big.NewInt(0)
	rE.ToBigIntRegular(r)
	return r, nil
}

func HashAccurate(inpBI []*big.Int, t int) (*big.Int, error) {
	inp := BigIntArrayToElementArray(inpBI[:])
	state := make([]*Element, t)
	for i := 0; i < t; i++ {
		state[i] = zero()
	}

	copy(state[:t], inp[:])

	nRoundsF := NROUNDSF
	var nRoundsP = NROUNDSP
	nSp, ok := NROUNDSPMAP[t]
	if ok {
		nRoundsP = nSp
	}

	newState := make([]*Element, t)
	for i := 0; i < t; i++ {
		newState[i] = zero()
	}

	// ARK --> SBox --> M, https://eprint.iacr.org/2019/458.pdf pag.5
	for i := 0; i < nRoundsF+nRoundsP; i++ {
		ark(state, c.c, i)
		sbox(nRoundsF, nRoundsP, state, i)
		mix(state, newState, c.m[t])
		state, newState = newState, state
	}
	rE := state[0]
	r := big.NewInt(0)
	rE.ToBigIntRegular(r)
	return r, nil
}
