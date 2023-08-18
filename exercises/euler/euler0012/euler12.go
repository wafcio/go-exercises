package main

import "math"

func Euler12(limitDivisorCount int) int {
	triangleNumberP := 0
	triangleNumber := 0
	divisorCount := 0

	for divisorCount <= limitDivisorCount {
		triangleNumberP += 1
		triangleNumber += triangleNumberP
		var divisors []int
		for factor := 1; factor <= int(math.Sqrt(float64(triangleNumber))); factor++ {
			if triangleNumber%factor == 0 {
				divisors = append(divisors, factor)
			}
		}
		divisorCount = len(divisors) * 2
	}

	return triangleNumber
}
