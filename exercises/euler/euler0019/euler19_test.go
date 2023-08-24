package main

import (
	"testing"
)

func TestEuler19(t *testing.T) {
	result := Euler19()
	if result != 171 {
		t.Fatalf(`Euler #19 returns invalid value = %d`, result)
	}
}
