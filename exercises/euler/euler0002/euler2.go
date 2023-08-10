package main

func Euler2(n int) int {
	num1 := 1
	num2 := 2
	sum := 2
	value := -1

	for value < n {
		value = num1 + num2

		if value%2 == 0 {
			sum += value
		}

		num1 = num2
		num2 = value
	}

	return sum
}
