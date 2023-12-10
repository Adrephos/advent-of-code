package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(path string) [][]int {
	var sequences [][]int
	data, _ := os.ReadFile(path)

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		line, sequence := strings.TrimSpace(line), []int{}
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			sequence = append(sequence, num)
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func getSubSequence(seq []int) ([]int, bool) {
	subseq, allZeroes := []int{}, 0
	for i := 0; i < len(seq)-1; i++ {
		diff := seq[i+1] - seq[i] 
		if diff == 0 { allZeroes++ }
		subseq = append(subseq, diff)
	}
	return subseq, allZeroes == len(subseq)
}

func nextValue(seq []int) int {
	subseq, allZeroes := getSubSequence(seq)
	if allZeroes { return seq[0] }

	return seq[len(seq)-1] + nextValue(subseq)
}

func prevValue(seq []int) int {
	subseq, allZeroes := getSubSequence(seq)
	if allZeroes { return seq[0] }

	return seq[0] - prevValue(subseq)
}

func firstPuzzle(sequences [][]int) (int, int) {
	first, second := 0, 0
	for _, sequence := range sequences {
		first += nextValue(sequence)
		second += prevValue(sequence)
	}
	return first, second
}

func main() {
	sequences := getData("./input.txt")
	first, second := firstPuzzle(sequences)
	fmt.Println("First Puzzle:", first)
	fmt.Println("Second Puzzle:", second)
}
