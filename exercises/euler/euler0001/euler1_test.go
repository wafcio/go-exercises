package main

import (
	"testing"
)

func TestEuler1WithInput10(t *testing.T) {
	result := Euler1(10)
	if result != 23 {
		t.Fatalf(`Euler #1 returns invalid value = %d`, result)
	}
}

func TestEuler1WithInput1_000(t *testing.T) {
	result := Euler1(1_000)
	if result != 233_168 {
		t.Fatalf(`Euler #1 returns invalid value = %d`, result)
	}
}
