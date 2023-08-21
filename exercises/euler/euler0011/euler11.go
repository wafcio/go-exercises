package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Euler11() int {
	content := readFile()
	product := 0

	product = checkHorizontallyAndVertically(content, product)
	product = checkDiagonally(content, product)

	return product
}

func readFile() [][]string {
	var output [][]string

	readFile, _ := os.Open("grid.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		output = append(output, strings.Split(fileScanner.Text(), " "))
	}
	readFile.Close()

	return output
}

func checkHorizontallyAndVertically(input_data [][]string, product int) int {
	for num1 := 0; num1 < 20; num1++ {
		for num2 := 0; num2 < 17; num2++ {
			newProduct := 1
			for index := 0; index < 4; index++ {
				value, _ := strconv.Atoi(input_data[num1][num2+index])
				newProduct *= value
			}
			product = pickGreater(product, newProduct)

			newProduct = 1
			for index := 0; index < 4; index++ {
				value, _ := strconv.Atoi(input_data[num2+index][num1])
				newProduct *= value
			}
			product = pickGreater(product, newProduct)
		}
	}

	return product
}

func checkDiagonally(input_data [][]string, product int) int {
	for num1 := 0; num1 < 17; num1++ {
		for num2 := 0; num2 < 17; num2++ {
			newProduct := 1
			for index := 0; index < 4; index++ {
				value, _ := strconv.Atoi(input_data[num1+index][num2+index])
				newProduct *= value
			}
			product = pickGreater(product, newProduct)

			newProduct = 1
			for index := 0; index < 4; index++ {
				value, _ := strconv.Atoi(input_data[num1+3-index][num2+index])
				newProduct *= value
			}
			product = pickGreater(product, newProduct)
		}
	}

	return product
}

func pickGreater(product int, newProduct int) int {
	if newProduct > product {
		return newProduct
	}
	return product
}
