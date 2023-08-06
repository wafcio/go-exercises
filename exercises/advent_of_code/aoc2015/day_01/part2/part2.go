package main

import (
	"strings"
)

func Aoc2015Day1Part2(input string) int {
	level := 0

	for i := 0; i < len(input); i++ {
		if strings.Compare(string(input[i]), "(") == 0 {
			level += 1
		} else if strings.Compare(string(input[i]), ")") == 0 {
			level -= 1
		}

		if level == -1 {
			return i + 1
		}
	}

	return -1
}
