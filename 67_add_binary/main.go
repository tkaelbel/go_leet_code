package main

import (
	"math/big"
)

func addBinary(a string, b string) string {
	bigA, ok := new(big.Int).SetString(a, 2)
	if !ok {
		return ""
	}

	bigB, ok := new(big.Int).SetString(b, 2)
	if !ok {
		return ""
	}

	sum := bigA.Add(bigA, bigB)

	return sum.Text(2)
}
