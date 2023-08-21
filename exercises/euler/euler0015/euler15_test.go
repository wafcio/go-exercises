package main

import (
	"testing"
)

func TestEuler15WithInput2(t *testing.T) {
	result := Euler15(2)
	if result != 6 {
		t.Fatalf(`Euler #15 returns invalid value = %d`, result)
	}
}

func TestEuler15WithInput100(t *testing.T) {
	result := Euler15(20)
	if result != 137_846_528_820 {
		t.Fatalf(`Euler #15 returns invalid value = %d`, result)
	}
}
