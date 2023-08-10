package main

import (
	"strconv"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i, j := 0, size-1; i < size>>1; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Euler4(n int) int {
	aStr := "1" + strings.Repeat("0", n-1)
	a, _ := strconv.Atoi(aStr)
	b := 0
	largestPalindrom := 0

	maxStr := strings.Repeat("9", n)
	max, _ := strconv.Atoi(maxStr)

	for a <= max {
		b = a

		for b <= max {
			result := a * b

			resultStr := strconv.Itoa(result)
			if resultStr == reverseString(resultStr) && result > largestPalindrom {
				largestPalindrom = result
			}

			b += 1
		}

		a += 1
	}

	return largestPalindrom
}
