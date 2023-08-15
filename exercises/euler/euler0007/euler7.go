package main

import (
	"fmt"
	"math"
)

func Euler7(n int) int {
	i := 0
	number := 2

	for {
		if isPrime(number) {
			fmt.Println(number)
			i += 1
		}

		if i == n {
			break
		}

		number += 1
	}

	return number
}

func isPrime(number int) bool {
	sq := int(math.Sqrt(float64(number)))

	for factor := 2; factor <= sq; factor++ {
		if number%factor == 0 {
			return false
		}
	}

	return true
}
