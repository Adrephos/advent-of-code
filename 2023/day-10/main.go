package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type tile struct {
	u byte
	d byte
	r byte
	l byte
}

func (t tile) getShape() byte {
	upPossible := []byte{'|', '7', 'F'}
	downPossible := []byte{'|', 'J', 'L'}
	rightPossible := []byte{'J', '-', '7'}
	leftPossible := []byte{'F', '-', 'L'}

	if slices.Contains(upPossible, t.u) && slices.Contains(downPossible, t.d) {
		return '|'
	} else if slices.Contains(leftPossible, t.l) && slices.Contains(rightPossible, t.r) {
		return '-'
	} else if slices.Contains(upPossible, t.u) && slices.Contains(rightPossible, t.r) {
		return 'L'
	} else if slices.Contains(upPossible, t.u) && slices.Contains(leftPossible, t.l) {
		return 'J'
	} else if slices.Contains(downPossible, t.d) && slices.Contains(leftPossible, t.l) {
		return '7'
	} else if slices.Contains(downPossible, t.d) && slices.Contains(rightPossible, t.r) {
		return 'F'
	}
	return '.'
}

func NewTile(lines [][]byte, i, j int) tile {
	var adjacent []byte
	for i_ad := -1; i_ad <= 1; i_ad++ {
		for j_ad := -1; j_ad <= 1; j_ad++ {
			if !((i_ad == 0 && j_ad != 0) || (i_ad != 0 && j_ad == 0)) {
				continue
			}
			var symbol byte = '.'
			if i+i_ad >= 0 && j+j_ad >= 0 && i+i_ad < len(lines) && j+j_ad < len(lines[i]) {
				symbol = lines[i+i_ad][j+j_ad]
			}
			adjacent = append(adjacent, symbol)
		}
	}
	return tile{u: adjacent[0], d: adjacent[3], r: adjacent[2], l: adjacent[1]}
}

func getData(path string) [][]byte {
	data, _ := os.ReadFile(path)
	var lines [][]byte
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		lines = append(lines, []byte(line))
	}

	return lines
}

func findStart(lines [][]byte) (int, int, tile) {
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				return i, j, NewTile(lines, i, j)
			}
		}
	}
	return -1, -1, tile{}
}

func validNext(lines [][]byte, i, j, k, l int) bool {
	sy := lines[i][j]
	candidate := lines[k][l]

	if (sy == '|' || sy == 'J' || sy == 'L') && k == i-1 {
		return slices.Contains([]byte{'|', '7', 'F'}, candidate)
	} else if (sy == '|' || sy == 'F' || sy == '7') && k == i+1 {
		return slices.Contains([]byte{'|', 'J', 'L'}, candidate)
	} else if (sy == '-' || sy == 'F' || sy == 'L') && l == j+1 {
		return slices.Contains([]byte{'-', 'J', '7'}, candidate)
	} else if (sy == '-' || sy == 'J' || sy == '7') && l == j-1 {
		return slices.Contains([]byte{'-', 'L', 'F'}, candidate)
	}
	return false
}

func getNext(lines [][]byte, i, j int) (int, int, bool) {
	var k, l int
	for i_ad := -1; i_ad <= 1; i_ad++ {
		for j_ad := -1; j_ad <= 1; j_ad++ {
			if !((i_ad == 0 && j_ad != 0) || (i_ad != 0 && j_ad == 0)) {
				continue
			}
			k, l = i+i_ad, j+j_ad
			if k >= 0 && l >= 0 && k < len(lines) && l < len(lines[i]) {
				valid := validNext(lines, i, j, k, l)
				if valid { return k, l, true }
			}
		}
	}
	return k, l, false
}

func firstPuzzle(lines [][]byte) (int, [][]int) {
	i, j, tile := findStart(lines)
	lines[i][j] = tile.getShape()
	var X, Y []int

	distance := 0

	for {
		k, l, succ := getNext(lines, i, j)
		if !succ {
			distance++
			break
		}
		lines[i][j] = '+'
		X, Y = append(X, j), append(Y, i)
		distance++
		i, j = k, l
	}

	return distance, [][]int{X, Y}
}

func shoeLaceFormula(X []int, Y []int, n int) float64 {
	area := 0.0

	j := n - 1
	for i := 0; i < n; i++ {
		area += float64(X[j]+X[i]) * float64(Y[j]-Y[i])
		j = i 
	}
	return math.Abs(area / 2.0)
}

func main() {
	lines := getData("./input.txt")
	distance, vertices := firstPuzzle(lines)
	fmt.Println("First Puzzle:", distance/2)
	area := shoeLaceFormula(vertices[0], vertices[1], len(vertices[0]))
	fmt.Println("ShoeLace Formula:", area)
	fmt.Println("Second Puzzle:", math.Round(area - float64(distance)/2.0 + 1))
}
