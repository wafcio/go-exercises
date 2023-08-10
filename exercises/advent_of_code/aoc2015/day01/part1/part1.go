package main

import (
	"strings"
)

func Aoc2015Day1Part1(input string) int {
	level := 0

	for i := 0; i < len(input); i++ {
		if strings.Compare(string(input[i]), "(") == 0 {
			level += 1
		} else if strings.Compare(string(input[i]), ")") == 0 {
			level -= 1
		}
	}

	return level
}
