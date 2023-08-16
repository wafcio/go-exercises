package main

import (
	"testing"
)

func TestEuler9WithInput12(t *testing.T) {
	result := Euler9(12)
	if result != 60 {
		t.Fatalf(`Euler #9 returns invalid value = %d`, result)
	}
}

func TestEuler9WithInput1_000(t *testing.T) {
	result := Euler9(1_000)
	if result != 31_875_000 {
		t.Fatalf(`Euler #9 returns invalid value = %d`, result)
	}
}
