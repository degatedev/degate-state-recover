package util

import (
	"log"
	"testing"
)

func TestCustomFloat(t *testing.T) {
	data := []int{
		1, 1, 1, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		0, 0, 0, 0, 1,
	}
	result, isNeg := ParseFloat30(data)
	if isNeg {
		log.Println("is neg")
	}
	log.Println("result:", result.String())
	// 33554431000000000000000
	// 167772150000000000000000000000000000000
	// 16777215000000000000000
	// 167772150000000
	// 167767030000000

}
