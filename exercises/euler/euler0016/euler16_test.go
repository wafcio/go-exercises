package main

import (
	"testing"
)

func TestEuler16WithInput15(t *testing.T) {
	result := Euler16(15)
	if result != 26 {
		t.Fatalf(`Euler #16 returns invalid value = %d`, result)
	}
}

func TestEuler16WithInput1_000(t *testing.T) {
	result := Euler16(1_000)
	if result != 1_366 {
		t.Fatalf(`Euler #16 returns invalid value = %d`, result)
	}
}
