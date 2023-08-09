package main

import (
	"testing"
)

func TestEuler5FirstExample(t *testing.T) {
	result := Euler5(10)
	if result != 2520 {
		t.Fatalf(`Euler #5 returns invalid value = %d`, result)
	}
}

func TestEuler5SecodExample(t *testing.T) {
	result := Euler5(20)
	if result != 232792560 {
		t.Fatalf(`Euler #5 returns invalid value = %d`, result)
	}
}
