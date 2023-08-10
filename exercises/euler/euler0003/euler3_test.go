package main

import (
	"testing"
)

func TestEuler3WithInput13_195(t *testing.T) {
	result := Euler3(13_195)
	if result != 29 {
		t.Fatalf(`Euler #3 returns invalid value = %d`, result)
	}
}

func TestEuler3WithInput600_851_475_143(t *testing.T) {
	result := Euler3(600_851_475_143)
	if result != 6_857 {
		t.Fatalf(`Euler #3 returns invalid value = %d`, result)
	}
}
