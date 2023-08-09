package main

import (
	"testing"
)

func TestEuler4FirstExample(t *testing.T) {
	result := Euler4(2)
	if result != 9009 {
		t.Fatalf(`Euler #4 returns invalid value = %d`, result)
	}
}

func TestEuler4SecodExample(t *testing.T) {
	result := Euler4(3)
	if result != 906609 {
		t.Fatalf(`Euler #4 returns invalid value = %d`, result)
	}
}
