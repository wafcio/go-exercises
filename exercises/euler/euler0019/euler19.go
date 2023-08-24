package main

import (
	"strconv"
	"time"
)

func Euler19() int {
	total := 0

	for year := 1901; year <= 2000; year += 1 {
		for month := 1; month <= 12; month += 1 {
			monthStr := strconv.Itoa(month)
			if len(monthStr) == 1 {
				monthStr = "0" + monthStr
			}

			str := strconv.Itoa(year) + "-" + monthStr + "-01T00:00:00.000Z"
			t, _ := time.Parse(time.RFC3339, str)

			if t.Weekday() == 0 {
				total += 1
			}
		}
	}

	return total
}
