package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func puzzle1(file string) int {
	total := 0

	posibilities := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	rounds := strings.Split(file, "\n")

	for _, round := range rounds {
		total += posibilities[round]
	}

	return total

}

func puzzle2(file string) int {
	total := 0

	posibilities := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	rounds := strings.Split(file, "\n")

	for _, round := range rounds {
		total += posibilities[round]
	}

	return total

}

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := string(f)

	fmt.Println("Puzzle 1:", puzzle1(file))
	fmt.Println("Puzzle 2:", puzzle2(file))

}
