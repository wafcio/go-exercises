package main

func Euler3(n int) int {
	m := 2

	for n > 1 {
		if n%m == 0 {
			n = n / m
		} else {
			m += 1
		}
	}

	return m
}
