package main

import (
	"io/ioutil"
	"strings"
)

func Euler8(n int) int {
	bigNumberStr := readFile()
	product := 0
	limit := len(bigNumberStr) - n

	for position := 0; position < limit; position++ {
		temp_product := 1
		for i := position; i <= position+n-1; i++ {
			num := int(bigNumberStr[i]) - 48
			temp_product = temp_product * num
		}

		if temp_product > product {
			product = temp_product
		}
	}

	return product
}

func readFile() string {
	content, _ := ioutil.ReadFile("big_number.txt")
	return strings.Replace(string(content), "\n", "", -1)
}
