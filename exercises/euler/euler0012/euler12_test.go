package main

import (
	"testing"
)

func TestEuler12WithInput5(t *testing.T) {
	result := Euler12(5)
	if result != 28 {
		t.Fatalf(`Euler #12 returns invalid value = %d`, result)
	}
}

func TestEuler12WithInput100(t *testing.T) {
	result := Euler12(100)
	if result != 73_920 {
		t.Fatalf(`Euler #12 returns invalid value = %d`, result)
	}
}

func TestEuler12WithInput500(t *testing.T) {
	result := Euler12(500)
	if result != 76_576_500 {
		t.Fatalf(`Euler #12 returns invalid value = %d`, result)
	}
}
