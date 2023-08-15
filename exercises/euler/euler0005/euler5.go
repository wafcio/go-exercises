package main

func Euler5(n int) int {
	factors := []int{}

	for number := 2; number <= n; number++ {
		currentFactors := primeFactors(number)
		missingFactors := sliceDifference(currentFactors, factors)
		factors = append(factors, missingFactors...)
	}

	return multiply(factors)
}

func multiply(factors []int) int {
	result := 1

	for i := 0; i < len(factors); i++ {
		result *= factors[i]
	}

	return result
}

func sliceDifference(a, b []int) []int {
	diff := make([]int, len(a))
	copy(diff, a)

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(diff); j++ {
			if diff[j] == b[i] {
				diff[j] = diff[len(diff)-1]
				diff = diff[:len(diff)-1]
				break
			}
		}
	}

	return diff
}

func primeFactors(number int) []int {
	m := 2
	numbers := []int{}

	for number > 1 {
		if number%m == 0 {
			number = number / m
			numbers = append(numbers, m)
		} else {
			m += 1
		}
	}

	return numbers
}
