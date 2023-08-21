package main

import (
	"math/big"
)

func Euler16(size int) int {
	total := 0

	number := new(big.Int)
	number.SetInt64(2)

	exponent := new(big.Int)
	exponent.SetInt64(int64(size))

	zero := new(big.Int)
	zero.SetInt64(0)

	number.Exp(number, exponent, zero)

	numberStr := number.String()

	for i := 0; i < len(numberStr); i++ {
		total += (int(numberStr[i]) - 48)
	}

	return total
}
