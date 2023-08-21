package main

import (
	"testing"
)

func TestEuler14(t *testing.T) {
	result := Euler14()
	if result != 837_799 {
		t.Fatalf(`Euler #14 returns invalid value = %d`, result)
	}
}
