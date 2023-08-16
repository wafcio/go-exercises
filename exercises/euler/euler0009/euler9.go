package main

import (
	"math"
)

func Euler9(n int) int {
	a := 1
	b := 2

	for {
		powA := int(math.Pow(float64(a), 2))
		powB := int(math.Pow(float64(b), 2))
		c := math.Sqrt(float64(powA + powB))
		cInt := int(c)
		difference := c - float64(cInt)
		sum := a + b + cInt

		if a+b > n {
			break
		}

		if sum < n {
			b += 1
		}

		if sum > n || (sum == n && difference > 0) {
			a += 1
			b = a + 1
		}

		if sum == n && difference == 0 {
			return a * b * cInt
		}
	}

	return 0
}
