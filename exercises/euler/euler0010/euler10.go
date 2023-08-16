package main

import (
	"math"
)

func Euler10(n int) int {
	sum := 0
	for number := 2; number < n; number++ {
		if isPrime(number) {
			sum += number
		}
	}

	return sum
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
