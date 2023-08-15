package main

import "math"

func Euler6(n int) int {
	num1 := 0
	num2 := 0

	for m := 1; m <= n; m++ {
		num1 += int(math.Pow(float64(m), 2))
		num2 += m
	}

	num2 = int(math.Pow(float64(num2), 2))

	return num2 - num1
}
