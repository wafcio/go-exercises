package main

import (
	"testing"
)

func TestEuler8WithInput4(t *testing.T) {
	result := Euler8(4)
	if result != 5_832 {
		t.Fatalf(`Euler #8 returns invalid value = %d`, result)
	}
}

func TestEuler8WithInput13(t *testing.T) {
	result := Euler8(13)
	if result != 23_514_624_000 {
		t.Fatalf(`Euler #8 returns invalid value = %d`, result)
	}
}
