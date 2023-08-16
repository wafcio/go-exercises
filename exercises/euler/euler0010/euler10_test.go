package main

import (
	"testing"
)

func TestEuler10WithInput10(t *testing.T) {
	result := Euler10(10)
	if result != 17 {
		t.Fatalf(`Euler #10 returns invalid value = %d`, result)
	}
}

func TestEuler10WithInput2_000_000(t *testing.T) {
	result := Euler10(2_000_000)
	if result != 142_913_828_922 {
		t.Fatalf(`Euler #10 returns invalid value = %d`, result)
	}
}
