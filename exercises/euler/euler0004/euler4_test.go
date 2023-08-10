package main

import (
	"testing"
)

func TestEuler4WithInput2(t *testing.T) {
	result := Euler4(2)
	if result != 9_009 {
		t.Fatalf(`Euler #4 returns invalid value = %d`, result)
	}
}

func TestEuler4WithInput3(t *testing.T) {
	result := Euler4(3)
	if result != 906_609 {
		t.Fatalf(`Euler #4 returns invalid value = %d`, result)
	}
}
