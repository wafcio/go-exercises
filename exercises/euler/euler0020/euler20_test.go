package main

import (
	"testing"
)

func TestEuler20WithInput10(t *testing.T) {
	result := Euler20(10)
	if result != 27 {
		t.Fatalf(`Euler #20 returns invalid value = %d`, result)
	}
}

func TestEuler20WithInput100(t *testing.T) {
	result := Euler20(100)
	if result != 648 {
		t.Fatalf(`Euler #20 returns invalid value = %d`, result)
	}
}
