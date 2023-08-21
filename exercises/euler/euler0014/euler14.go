package main

func Euler14() int {
	max := 0
	start := 0
	for x := 0; x <= 1_000_000; x++ {
		c := countTerms(x)
		if c > max {
			max = c
			start = x
		}
	}

	return start
}

func countTerms(n int) int {
	if n <= 0 {
		return 0
	}

	c := 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		c += 1
	}

	return c
}
