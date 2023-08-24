package main

import (
	"sort"
	"strconv"
	"strings"
)

func Euler18(input string) int {
	triangle := parseData(input)
	result := calculateMaxSum(triangle, 0, 0)

	return result
}

func parseData(data string) [][]string {
	var output [][]string

	lines := strings.Split(data, "\n")
	for i := 0; i < len(lines); i++ {
		output = append(output, strings.Split(lines[i], " "))
	}
	output = output[:len(output)-1]

	return output
}

func calculateMaxSum(triangle [][]string, x, y int) int {
	if y == len(triangle)-1 {
		value, _ := strconv.Atoi(triangle[y][x])
		return value
	} else {
		var list []int
		list = append(list, calculateMaxSum(triangle, x, y+1))
		list = append(list, calculateMaxSum(triangle, x+1, y+1))
		sort.Ints(list)
		intValue, _ := strconv.Atoi(triangle[y][x])

		value := list[1] + intValue
		return value
	}
}
