package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(f)
	inventories := strings.Split(file, "\n\n")
	var caloriesSum []int
	for _, inv := range inventories {
		caloriesList := strings.Split(inv, "\n")
		var sum int
		for _, calories := range caloriesList {
			caloriesInt, _ := strconv.Atoi(calories)
			sum += caloriesInt
		}
		caloriesSum = append(caloriesSum, sum)
	}
	sort.Ints(caloriesSum[:])

	var top3 int
	for i := 1; i <= 3; i++ {
		top3 += caloriesSum[len(caloriesSum)-i]
	}

	fmt.Println("Puzzle 1:", caloriesSum[len(caloriesSum)-1])
	fmt.Println("Puzzle 2:", top3)
}
