package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	i int
	j int
}

type Ratio struct {
	nums  int
	ratio int
}

func getLines(path string) []string {
	data, _ := os.ReadFile(path)
	txt := strings.TrimSpace(string(data))
	lines := strings.Split(txt, "\n")

	return lines
}

func checkAdjacent(lines []string, i int, j int) (bool, bool, int, int) {
	symbol, gear := false, false
	gearI, gearJ := -1, -1
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if i+di < 0 || i+di >= len(lines) || j+dj < 0 || j+dj >= len(lines[i]) {
				continue
			}
			if di != 0 || dj != 0 {
				char := lines[i+di][j+dj]
				if char == '*' {
					gear = true
					gearI, gearJ = i+di, j+dj
				}
				if char != '.' && !('0' <= char && char <= '9') {
					symbol = true
				}
			}
		}
	}
	return symbol, gear, gearI, gearJ
}

func sumGears(gears map[Coord]Ratio) int {
	sum := 0
	for _, gear := range gears {
		if gear.nums == 2 {
			sum += gear.ratio
		}
	}
	return sum
}

func findSum(lines []string) (int, int) {
	sum := 0
	gears := make(map[Coord]Ratio)
	for i := range lines {
		numChars, adjSy, adjGear := []byte{}, false, false
		gearI, gearJ := -1, -1
		for j := range lines[i] {
			char := lines[i][j]
			if '0' <= char && char <= '9' {
				numChars = append(numChars, char)
				symbol, gear, gI, gJ := checkAdjacent(lines, i, j)
				if symbol {
					adjSy = true
				}
				if gear {
					adjGear = true
					gearI, gearJ = gI, gJ
				}
			}
			if !('0' <= char && char <= '9') || j == len(lines[i])-1 {
				num, _ := strconv.Atoi(string(numChars))
				if adjSy {
					sum += num
				}
				if adjGear {
					entry, ok := gears[Coord{gearI, gearJ}]
					if !ok {
						entry.ratio = num
					} else {
						entry.ratio *= num
					}
					entry.nums++
					gears[Coord{gearI, gearJ}] = entry
				}

				numChars, adjSy, adjGear = []byte{}, false, false
			}
		}
	}
	return sum, sumGears(gears)
}

func main() {
	lines := getLines("input.txt")
	sum, sumGears := findSum(lines)
	fmt.Println(sum)
	fmt.Println(sumGears)
}
