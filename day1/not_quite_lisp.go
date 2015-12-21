package main

import (
	"fmt"
	"io/ioutil"
)

func calculateFloor(input string) int {
	floor := 0
	for _, i := range input {
		s := string(i)
		if s == "(" {
			floor += 1
		} else if s == ")" {
			floor -= 1
		}
	}
	return floor
}

func calculateBasementEntrance(input string) int {
	floor := 0
	for pos, i := range input {
		s := string(i)
		if s == "(" {
			floor += 1
		} else if s == ")" {
			floor -= 1
		}
		if floor < 0 {
			return pos + 1
		}
	}
	return 0
}

func main() {
	dat, err := ioutil.ReadFile("santa_instructions.txt")
	if err != nil {
		panic(err)
	}

	//To what floor do the instructions take Santa?
	floor := calculateFloor(string(dat))
	fmt.Println("Floor: ", floor)

	//What is the position of the character that causes Santa to first enter the basement?
	position := calculateBasementEntrance(string(dat))
	fmt.Println("Position", position)

}
