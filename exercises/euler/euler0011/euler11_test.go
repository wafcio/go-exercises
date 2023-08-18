package main

import (
	"testing"
)

func TestEuler11WithBigInput(t *testing.T) {
	result := Euler11()
	if result != 70_600_674 {
		t.Fatalf(`Euler #11 returns invalid value = %d`, result)
	}
}
