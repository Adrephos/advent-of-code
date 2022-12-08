package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findWrongPlaced(rucksack []string) int {
	total := 0

	for _, items := range rucksack {
		if len(items) != 0 {
			half := (len(items) / 2)
			compartment1, compartment2 := items[0:half], items[half:]

			priority := 0

			index := strings.IndexAny(compartment1, compartment2)
			ascii := int(compartment1[index])

			if ascii >= 65 && ascii <= 90 {
				priority = ascii - 38
			} else {
				priority = ascii - 96
			}
			total += priority
		}

	}
	return total
}

func getCommon(str, str2 string) string {
	common := ""
	for _, rune1 := range str {
		for _, rune2 := range str2 {
			if rune1 == rune2 {
				common += string(rune1)
			}
		}
	}
	return common
}

func findBadges(rucksack []string) int {
	lngt := len(rucksack)
	total := 0
	for i := 0; i <= lngt-3; i += 3 {
		elf1, elf2, elf3 := rucksack[i], rucksack[i+1], rucksack[i+2]

		common := getCommon(elf1, elf2)
		index := strings.IndexAny(common, elf3)
		priority := 0
		ascii := int(common[index])

		if ascii >= 65 && ascii <= 90 {
			priority = ascii - 38
		} else {
			priority = ascii - 96
		}
		total += priority

	}

	return total
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	file := string(f)
	rucksack := strings.Split(file, "\n")

	fmt.Println("Puzzle 1:", findWrongPlaced(rucksack))
	fmt.Println("Puzzle 2:", findBadges(rucksack))
}
