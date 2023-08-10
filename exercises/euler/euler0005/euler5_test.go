package main

import (
	"testing"
)

func TestEuler5WithInput10(t *testing.T) {
	result := Euler5(10)
	if result != 2_520 {
		t.Fatalf(`Euler #5 returns invalid value = %d`, result)
	}
}

func TestEuler5WithInput20(t *testing.T) {
	result := Euler5(20)
	if result != 232_792_560 {
		t.Fatalf(`Euler #5 returns invalid value = %d`, result)
	}
}
