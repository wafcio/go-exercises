package main

import (
	"testing"
)

func TestEuler13WithInputFromFile(t *testing.T) {
	result := Euler13()
	if result != "5537376230" {
		t.Fatalf(`Euler #13 returns invalid value = %s`, result)
	}
}
