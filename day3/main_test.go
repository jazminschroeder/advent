package main

import "testing"

func TestMoveRight(t *testing.T) {
	starting := [2]int{0, 0}
	expected := [2]int{1, 0}
	if position := moveRight(starting); position != expected {
		t.Errorf("Error: expected %v, got %v", expected, position)
	}
}

func TestMoveLeft(t *testing.T) {
	starting := [2]int{0, 0}
	expected := [2]int{-1, 0}
	if position := moveLeft(starting); position != expected {
		t.Errorf("Error: expected %v, got %v", expected, position)
	}
}

func TestMoveUp(t *testing.T) {
	starting := [2]int{0, 0}
	expected := [2]int{0, 1}
	if position := moveUp(starting); position != expected {
		t.Errorf("Error: expected %v, got %v", expected, position)
	}
}

func TestMoveDown(t *testing.T) {
	starting := [2]int{0, 0}
	expected := [2]int{0, -1}
	if position := moveDown(starting); position != expected {
		t.Errorf("Error: expected %v, got %v", expected, position)
	}
}

func TestCountVisitedHouses(t *testing.T) {
	input := []byte("^>v<")
	expected := 4
	if hv := countVisitedHouses(input); hv != expected {
		t.Errorf("Error: expected %d, got %d", expected, hv)
	}
}

func TestCountVisitedHousesWithRobo(t *testing.T) {
	input := []byte("^v^v^v^v^v")
	expected := 11
	if hv := countVisitedHousesWithRobo(input); hv != expected {
		t.Errorf("Error: expected %d, got %d", expected, hv)
	}
}
