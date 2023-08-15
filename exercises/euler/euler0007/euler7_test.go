package main

import (
	"testing"
)

func TestEuler7WithInput6(t *testing.T) {
	result := Euler7(6)
	if result != 13 {
		t.Fatalf(`Euler #7 returns invalid value = %d`, result)
	}
}

func TestEuler7WithInput10_001(t *testing.T) {
	result := Euler7(10_001)
	if result != 104_743 {
		t.Fatalf(`Euler #7 returns invalid value = %d`, result)
	}
}
