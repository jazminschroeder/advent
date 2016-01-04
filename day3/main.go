package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("How many houses received at least one present?", countVisitedHouses(input))
	fmt.Println("How many houses received at least one present with Santa-Robo's help", countVisitedHousesWithRobo(input))
}

func countVisitedHouses(input []byte) int {
	houses := make(map[[2]int]int)
	currentLocation := [2]int{0, 0}
	houses[currentLocation] = 1

	for _, s := range input {
		switch string(s) {
		case ">":
			currentLocation = moveRight(currentLocation)
		case "v":
			currentLocation = moveDown(currentLocation)
		case "<":
			currentLocation = moveLeft(currentLocation)
		case "^":
			currentLocation = moveUp(currentLocation)
		}
		houses[currentLocation] = houses[currentLocation] + 1
	}
	return len(houses)
}

func countVisitedHousesWithRobo(input []byte) int {
	houses := make(map[[2]int]int)
	distributors := make(map[string][2]int)
	var distributor string

	startingSantaLocation := [2]int{0, 0}
	startingRoboLocation := [2]int{0, 0}

	distributors["santa"] = startingSantaLocation
	distributors["robo"] = startingRoboLocation

	houses[startingSantaLocation] = 1
	houses[startingRoboLocation] = 1

	for index, s := range input {
		if index%2 == 0 {
			distributor = "santa"
		} else {
			distributor = "robo"
		}
		switch string(s) {
		case ">":
			distributors[distributor] = moveRight(distributors[distributor])
		case "v":
			distributors[distributor] = moveDown(distributors[distributor])
		case "<":
			distributors[distributor] = moveLeft(distributors[distributor])
		case "^":
			distributors[distributor] = moveUp(distributors[distributor])
		}
		houses[distributors[distributor]] = houses[distributors[distributor]] + 1
	}
	return len(houses)
}

func moveRight(start [2]int) [2]int {
	x := start[0] + 1
	y := start[1]
	return [2]int{x, y}
}

func moveLeft(start [2]int) [2]int {
	x := start[0] - 1
	y := start[1]
	return [2]int{x, y}
}
func moveUp(start [2]int) [2]int {
	x := start[0]
	y := start[1] + 1
	return [2]int{x, y}
}
func moveDown(start [2]int) [2]int {
	x := start[0]
	y := start[1] - 1
	return [2]int{x, y}
}
