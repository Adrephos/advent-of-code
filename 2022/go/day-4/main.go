package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func toIntArray(str string) []int {
	strArr := strings.Split(str, "-")
	intArr := []int{}
	for _, ranges := range strArr {
		num, _ := strconv.Atoi(ranges)
		intArr = append(intArr, num)
	}
	return intArr
}

func getOverlaps(str string) (puzz1, puzz2 int) {
	total1 := 0
	total2 := 0

	assignmentPairs := strings.Split(str, "\n")
	for _, pair := range assignmentPairs {
		if len(pair) != 0 {
			sections := strings.Split(pair, ",")
			elf1 := toIntArray(sections[0])
			elf2 := toIntArray(sections[1])

			if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] {
				total1 += 1
			} else if elf2[0] <= elf1[0] && elf2[1] >= elf1[1] {
				total1 += 1
			}
			if elf1[1] >= elf2[0] && elf1[0] <= elf2[0] {
				total2 += 1
			} else if elf2[1] >= elf1[0] && elf2[0] <= elf1[0] {
				total2 += 1
			}

		}
	}
	return total1, total2
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := string(f)

	ans1, ans2 := getOverlaps(file)
	fmt.Println("Puzzle 1:", ans1)
	fmt.Println("Puzzle 2:", ans2)
}
