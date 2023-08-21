package main

import (
	"bufio"
	"math/big"
	"os"
)

func Euler13() string {
	bigNumbers := readFile()

	total := new(big.Int)
	total.SetInt64(0)
	for i := 0; i < len(bigNumbers); i++ {
		total.Add(total, &(bigNumbers[i]))
	}

	return total.String()[0:10]
}

func readFile() []big.Int {
	var output []big.Int

	readFile, _ := os.Open("big_numbers.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		intValue := new(big.Int)
		intValue.SetString(fileScanner.Text(), 10)
		output = append(output, *intValue)
	}
	readFile.Close()

	return output
}
