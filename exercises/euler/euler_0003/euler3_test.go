package main

import (
	"testing"
)

func TestEuler3FirstExample(t *testing.T) {
	result := Euler3(13195)
	if result != 29 {
		t.Fatalf(`Euler #3 returns invalid value = %d`, result)
	}
}

func TestEuler3SecodExample(t *testing.T) {
	result := Euler3(600851475143)
	if result != 6857 {
		t.Fatalf(`Euler #3 returns invalid value = %d`, result)
	}
}
