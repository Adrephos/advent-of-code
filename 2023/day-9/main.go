package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
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

func extraValueTR(seq []int, acc int) int {
	subseq, allZeroes := getSubSequence(seq)
	if allZeroes { return seq[0] + acc }

	return extraValueTR(subseq, seq[len(seq)-1] + acc)
}

func extraValue(seq []int) int { return extraValueTR(seq, 0) }

func firstPuzzle(sequences [][]int) (int, int) {
	first, second := 0, 0
	for _, sequence := range sequences {
		first += extraValue(sequence)
		slices.Reverse(sequence)
		second += extraValue(sequence)
	}
	return first, second
}

func main() {
	sequences := getData("./input.txt")
	first, second := firstPuzzle(sequences)
	fmt.Println("First Puzzle:", first)
	fmt.Println("Second Puzzle:", second)
}
