package main

import (
	"testing"
)

func TestFindSecretNumber(t *testing.T) {
	expected := 1048970
	input := "pqrstuv"
	if number := findSecretNumber(input, "00000"); number != expected {
		t.Errorf("Error: expected %v, got %v", expected, number)
	}
}
