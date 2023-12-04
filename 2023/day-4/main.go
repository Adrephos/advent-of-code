package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCards(path string) []string {
	data, _ := os.ReadFile(path)
	txt := strings.Trim(string(data), "\n")
	cards := strings.Split(txt, "\n")

	return cards
}

func winningNumbersToMap(winningNumbers []string) map[string]bool {
	winningNumbersMap := make(map[string]bool)

	for _, number := range winningNumbers {
		winningNumbersMap[number] = true
	}

	return winningNumbersMap
}

func getCardsWinningNumbers(card string) (int, map[string]bool, []string) {
	cardNumbers := strings.Split(card, ":")
	cardNumberStr := strings.TrimSpace(cardNumbers[0])
	cardNumberStr = strings.Split(cardNumberStr, " ")[len(strings.Split(cardNumberStr, " "))-1]
	cardNumbers = strings.Split(cardNumbers[1], "|")

	cardNumbers[0] = strings.TrimSpace(cardNumbers[0])
	winningNumbers := strings.Split(cardNumbers[0], " ")

	cardNumbers[1] = strings.TrimSpace(cardNumbers[1])
	obtainedNumbers := strings.Split(cardNumbers[1], " ")

	cardNumber, _ := strconv.Atoi(cardNumberStr)

	return cardNumber, winningNumbersToMap(winningNumbers), obtainedNumbers
}

func getPoints(winningNumbersMap map[string]bool, obtainedNumbers []string) (int, int) {
	points, copies := 0, 0
	for _, num := range obtainedNumbers {
		if _, ok := winningNumbersMap[num]; !ok || num == "" {
			continue
		}
		if points == 0 {
			points++
		} else {
			points *= 2
		}
		copies++
	}

	return points, copies
}

func getCopies(cardsCopies *map[int]int, cardNumber int) int {
	copies, ok := (*cardsCopies)[cardNumber]
	if !ok {
		(*cardsCopies)[cardNumber] = 1
		return 1
	}
	return copies
}

func addCopies(cardCopies *map[int]int, cardNumber int, copies int) {
	for i := cardNumber+1; i <= cardNumber+copies; i++ {
		(*cardCopies)[i] = getCopies(cardCopies, i) + getCopies(cardCopies, cardNumber)
	}
}

func sumCopies(cardCopies map[int]int) int {
	sum := 0
	for _, copies := range cardCopies {
		sum += copies
	}

	return sum
}

func getTotalPoints(cards []string) (int, int) {
	sum := 0
	cardCopies := make(map[int]int)
	for _, card := range cards {
		cardNumber, winningNumbersMap, obtainedNumbers := getCardsWinningNumbers(card)
		points, copies := getPoints(winningNumbersMap, obtainedNumbers)
		getCopies(&cardCopies, cardNumber)
		addCopies(&cardCopies, cardNumber, copies)
		sum += points
	}
	
	return sum, sumCopies(cardCopies)
}

func main() {
	cards := getCards("./input.txt")
	totalPoints, totalCards := getTotalPoints(cards)
	fmt.Println("First Puzzle:", totalPoints)
	fmt.Println("Second Puzzle:", totalCards)
}
