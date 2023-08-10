package main

import (
	"testing"
)

func TestEuler2WithInput4_000_000(t *testing.T) {
	result := Euler2(4_000_000)
	if result != 4_613_732 {
		t.Fatalf(`Euler #2 returns invalid value = %d`, result)
	}
}
