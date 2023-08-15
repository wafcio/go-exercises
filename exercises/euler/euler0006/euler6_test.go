package main

import (
	"testing"
)

func TestEuler6WithInput10(t *testing.T) {
	result := Euler6(10)
	if result != 2_640 {
		t.Fatalf(`Euler #6 returns invalid value = %d`, result)
	}
}

func TestEuler6WithInput20(t *testing.T) {
	result := Euler6(100)
	if result != 25_164_150 {
		t.Fatalf(`Euler #6 returns invalid value = %d`, result)
	}
}
