package main

import "strings"

var Words = []string{
	"",
	"one ",
	"two ",
	"three ",
	"four ",
	"five ",
	"six ",
	"seven ",
	"eight ",
	"nine ",
	"ten ",
	"eleven ",
	"twelve ",
	"thirteen ",
	"fourteen ",
	"fifteen ",
	"sixteen ",
	"seventeen ",
	"eighteen ",
	"nineteen ",
}

var DecimalWords = []string{
	"",
	"",
	"twenty ",
	"thirty ",
	"forty ",
	"fifty ",
	"sixty ",
	"seventy ",
	"eighty ",
	"ninety ",
}

func Euler17(max int) int {
	total := 0

	for number := 1; number <= max; number += 1 {
		numberStr := strings.Replace(inWords(number), " ", "", -1)
		total += len(numberStr)
	}

	return total
}

func inWords(number int) string {
	if number >= 1_000 {
		return bigNumberInWords(number)
	}

	if number >= 100 {
		return houndredsInWord(number)
	}

	if number >= 20 {
		return DecimalWords[number/10] + " " + inWords(number%10)
	}

	if number > 0 {
		return Words[number]
	}

	return ""
}

func bigNumberInWords(number int) string {
	if number >= 1_000_000_000 {
		return inWords(number/1_000_000_000) + "billion " + inWords(number%1_000_000_000)
	}

	if number >= 1_000_000 {
		return inWords(number/1_000_000) + "million " + inWords(number%1_000_000)
	}

	return inWords(number/1_000) + "thousand " + inWords(number%1_000)
}

func houndredsInWord(number int) string {
	if number >= 100 && number%100 == 0 {
		return inWords(number/100) + "hundred "
	}

	return inWords(number/100) + "hundred and " + inWords(number%100)
}
