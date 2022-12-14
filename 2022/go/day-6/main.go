package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getStartOfPacket(file string, size int) int {
	str := ""
	counter := 0
	index := 0
	char1 := ""
	for i, char := range file {
		char1 = string(char)
		if strings.ContainsAny(str, char1) {
			idx := strings.Index(str, char1)
			str = str[idx+1:] + char1
			counter = len(str)
		} else {
			counter++
			str += char1
			if counter == size {
				index = i + 1
				break
			}
		}
	}

	return index

}

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := strings.TrimSpace(string(f))

	fmt.Println("Puzzle 1:", getStartOfPacket(file, 4))
	fmt.Println("Puzzle 2:", getStartOfPacket(file, 14))

}
