package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGames(path string) []string {
	data, _ := os.ReadFile(path)
	txt := strings.Trim(string(data), "\n")
	games := strings.Split(txt, "\n")

	return games
}

// get game id and set of cubes
func getGameContent(game string) (int, []string) {
	setsNId := strings.Split(game, ":")
	gameId, _ := strconv.Atoi(strings.Split(setsNId[0], " ")[1])
	sets := strings.Split(setsNId[1], ";")

	return gameId, sets
}

func getNumberAndColor(cube string) (int, string) {
	cube = strings.TrimSpace(cube)
	cubeArr := strings.Split(cube, " ")
	num, _ := strconv.Atoi(cubeArr[0])

	return num, cubeArr[1]
}

func validateSet(set string, available map[string]int) bool {
	set = strings.TrimSpace(set)
	cubes := strings.Split(set, ",")

	for _, cube := range cubes {
		num, color := getNumberAndColor(cube)

		if num > available[color] {
			return false
		}
	}

	return true
}

func validateSets(sets []string, available map[string]int) bool {
	for _, set := range sets {
		if !validateSet(set, available) {
			return false
		}
	}
	return true
}

func getPossibleSum(games []string, available map[string]int) int {
	sum := 0
	for _, game := range games {
		gameId, sets := getGameContent(game)
		if validateSets(sets, available) {
			sum += gameId
		}
	}
	return sum
}

func getMinFreqs(sets []string) map[string]int {
	freqs := make(map[string]int)

	for _, set := range sets {
		set = strings.TrimSpace(set)
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			num, color := getNumberAndColor(cube)

			if freq, ok := freqs[color]; ok{
				if num > freq { freqs[color] = num }
			} else if !ok {
				freqs[color] = num
			}
		}
	}

	return freqs
}

func getPower(freqs map[string]int) int {
	power := 1
	for _, freq := range freqs {
		power *= freq
	}
	return power
}

func getPowerSum(games []string) int {
	sum := 0
	for _, game := range games {
		_, sets := getGameContent(game)
		freqs := getMinFreqs(sets)
		sum += getPower(freqs)
	}
	return sum
}

func main() {
	games := getGames("./input.txt")
	available := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	fmt.Println(getPossibleSum(games, available))

	fmt.Println(getPowerSum(games))
}
