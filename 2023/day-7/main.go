package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getHandsAndBids(path string) []string {
	data, _ := os.ReadFile(path)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}

func getCardStrenght(card rune, newRules bool) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if newRules { return 1 }
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func maxFreq(cardFreq map[byte]int, newRules bool) (byte, int) {
	card, max := byte(' '), 0
	for key, freq := range cardFreq {
		if key == 'J' && newRules { continue }
		if freq > max { card, max = key, freq }
	}
	return card, max
}

func getHandStrenght(hand string, newRules bool) int {
	cardFreq := make(map[byte]int)
	for _, card := range hand { cardFreq[byte(card)] += 1 }
	maxCard, maxFreq := maxFreq(cardFreq, newRules)

	if newRules {
		freq, ok := cardFreq['J']
		if maxCard != 'J' && ok  { 
			maxFreq += freq
			delete(cardFreq, 'J')
		}
	}

	if len(cardFreq) == 5 { return 1
	} else if maxFreq == 2 && len(cardFreq) == 4 { return 2
	} else if maxFreq == 2 && len(cardFreq) == 3 { return 3
	} else if maxFreq == 3 && len(cardFreq) == 3 { return 4
	} else if maxFreq == 3 && len(cardFreq) == 2 { return 5
	} else if maxFreq == 4{ return 6 }

	return 7
}

func rankHands(lines []string, newRules bool) []string {
	sort.Slice(lines, func(i, j int) bool {
		hand_i := strings.Split(lines[i], " ")[0]
		hand_j := strings.Split(lines[j], " ")[0]

		st_i := getHandStrenght(hand_i, newRules)
		st_j := getHandStrenght(hand_j, newRules)

		if st_i == st_j {
			for k, card_i := range hand_i {
				card_j := rune(hand_j[k])
				if card_i != card_j {
					return getCardStrenght(card_i, newRules) < getCardStrenght(card_j, newRules)
				}
			}
		}

		return st_i < st_j
	})
	return lines
}

func firstPuzzle(lines []string, newRules bool) int {
	lines, totalWinnigs := rankHands(lines, newRules), 0
	for rank, line := range lines {
		bidStr, rank := strings.Split(line, " ")[1], rank + 1
		bid, _ := strconv.Atoi(bidStr)
		totalWinnigs += bid * rank
	}
	return totalWinnigs
}

func main() {
	lines := getHandsAndBids("./input.txt")
	fmt.Println("First Puzzle:", firstPuzzle(lines, false))
	fmt.Println("Second Puzzle:", firstPuzzle(lines, true))
}
