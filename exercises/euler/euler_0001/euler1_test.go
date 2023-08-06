package main

import (
	"testing"
)

func TestEuler1(t *testing.T) {
	result := Euler1()
	if result != 233168 {
		t.Fatalf(`Euler #1 returns invalid value = %d`, result)
	}
}
