package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var paper, ribbon int

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l, w, h := parseDimensions(scanner.Text())
		paper += calculateWrappingPaper(l, w, h)
		ribbon += calculateRibbon(l, w, h)
	}
	fmt.Println("Wrapping paper needed: ", paper)
	fmt.Println("Ribbon needed: ", ribbon)
}

func parseDimensions(input string) (int, int, int) {
	dimensions := strings.Split(input, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])
	return l, w, h
}

func calculateWrappingPaper(l, w, h int) int {
	sides := calculateSidesArea(l, w, h)
	paper := 2*sides[0] + 2*sides[1] + 2*sides[2]
	slack := findSmallest(sides)
	return paper + slack
}

func calculateSidesArea(l, w, h int) []int {
	side1 := l * w
	side2 := h * w
	side3 := h * l
	return []int{side1, side2, side3}
}

func calculateSidesPerimeter(l, w, h int) []int {
	side1 := 2*l + 2*w
	side2 := 2*h + 2*w
	side3 := 2*h + 2*l
	return []int{side1, side2, side3}
}

func findSmallest(sides []int) int {
	smallestSide := sides[0]
	for _, v := range sides {
		if v < smallestSide {
			smallestSide = v
		}
	}
	return smallestSide
}

func calculateRibbon(l, w, h int) int {
	sides := calculateSidesPerimeter(l, w, h)
	ribbon := findSmallest(sides)
	bow := l * w * h
	return ribbon + bow
}
