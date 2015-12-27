package main

import "testing"

func TestCalulateWrappingPaper(t *testing.T) {
	expected := 58
	if paper := calculateWrappingPaper(2, 3, 4); paper != expected {
		t.Errorf("Error, expected %d, got %d", expected, paper)
	}
}

func TestCalculateRibbon(t *testing.T) {
	expected := 14
	if ribbon := calculateRibbon(1, 1, 10); ribbon != expected {
		t.Errorf("Error, expected %d, got %d", expected, ribbon)
	}
}
