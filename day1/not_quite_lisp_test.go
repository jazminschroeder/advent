package main

import (
	"testing"
)

func TestGoUp(t *testing.T) {
	input := "(("
	expected := 2
	if floor := calculateFloor(input); floor != expected {
		t.Errorf("Error, expected %d got %d", expected, floor)
	}
}

func TestGoDown(t *testing.T) {
	input := ")))"
	expected := -3
	if floor := calculateFloor(input); floor != expected {
		t.Errorf("Expected %d got %d", expected, floor)
	}
}

func TestCalculateFloor(t *testing.T) {
	input := "(((())))(((("
	expected := 4
	if floor := calculateFloor(input); floor != expected {
		t.Errorf("Expected %d got %d", expected, floor)
	}
}

func TestFindBasementEntrance(t *testing.T) {
	input := "()())"
	expected := 5
	if position := calculateBasementEntrance(input); position != expected {
		t.Errorf("Expected %d got %d", expected, position)
	}
}
