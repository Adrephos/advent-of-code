package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToIntArr(str string) []int {
	num, arr, isNum := []byte{}, []int{}, false

	for i := 0; i < len(str); i++ {
		char := str[i]
		if '0' <= char && char <= '9' { num, isNum = append(num, char), true } 
		if (!('0' <= char && char <= '9') || i == len(str)-1) && isNum {
			intValue, _ := strconv.Atoi(string(num))
			arr, num, isNum = append(arr, intValue), []byte{}, false
		}
	}
	return arr
}

func getTimeAndDistance(path string) (string, string) {
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))
	time, _ := strings.CutPrefix(strings.Split(input, "\n")[0], "Time:")
	distance, _ := strings.CutPrefix(strings.Split(input, "\n")[1], "Distance:")

	return time, distance
}

func findAlternatives(time int, distance int) int {
	num := 0
	for i := 0; i <= time; i++ {
		if i * (time - i) > distance { num++ }
	}
	return num
}

func firstPuzzle(times []int, distances []int) int {
	mult := 1
	for i, time := range times {
		distance := distances[i]
		mult *= findAlternatives(time, distance)
	}
	return mult
}

func getNumber(str string) int {
	var num []rune
	for _, char := range str {
		if '0' <= char && char <= '9' { num = append(num, char) }
	}
	intVal, _ := strconv.Atoi(string(num))
	return intVal
}

func secondPuzzle(time string, distance string) int {
	return findAlternatives(getNumber(time), getNumber(distance))
}

func main() {
	times, distances :=	getTimeAndDistance("./input.txt")
	ans := firstPuzzle(strToIntArr(times), strToIntArr(distances))
	fmt.Println("First Puzzle:", ans)
	ans = secondPuzzle(times, distances)
	fmt.Println("Second Puzzle:", ans)
}
