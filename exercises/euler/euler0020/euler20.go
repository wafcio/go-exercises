package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func Euler20(n int) int {
	total := 0
	nBig := big.NewInt(int64(n))

	factorialStr := factorial(nBig).String()
	fmt.Println(factorialStr)
	for i := 0; i < len(factorialStr); i++ {
		number, _ := strconv.Atoi(string(factorialStr[i]))
		total += number
	}

	return total
}

func factorial(num *big.Int) *big.Int {
	result := big.NewInt(int64(1))
	m := big.NewInt(int64(1))

	for m.Cmp(num) == -1 || m.Cmp(num) == -1 {
		result.Mul(result, m)

		m.Add(m, big.NewInt(int64(1)))
	}

	return result
}
