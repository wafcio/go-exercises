package main

import (
	"testing"
)

func TestEuler17WithInput5(t *testing.T) {
	result := Euler17(5)
	if result != 19 {
		t.Fatalf(`Euler #17 returns invalid value = %d`, result)
	}
}

func TestEuler17WithInput1_000(t *testing.T) {
	result := Euler17(1_000)
	if result != 21_124 {
		t.Fatalf(`Euler #17 returns invalid value = %d`, result)
	}
}
