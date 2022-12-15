package main

import (
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {

	var ret []int
	var numAsString string
	for _, item := range digits {
		numAsString = numAsString + strconv.Itoa(item)
	}

	bigNum, ok := new(big.Int).SetString(numAsString, 0)
	if !ok {
		return ret
	}
	bigNum = bigNum.Add(bigNum, big.NewInt(int64(1)))

	numAsString = bigNum.String()

	// numAsString = strconv.Itoa(int(newVal))
	for i := 0; i < len(numAsString); i++ {
		val, err := strconv.ParseInt(string(numAsString[i]), 0, 0)
		ret = append(ret, int(val))
		if err != nil {
			break
		}
	}

	return ret
}
