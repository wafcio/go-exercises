package main

import (
	"testing"
)

func TestEuler2(t *testing.T) {
	result := Euler2()
	if result != 4_613_732 {
		t.Fatalf(`Euler #2 returns invalid value = %d`, result)
	}
}
