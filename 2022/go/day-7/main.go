package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func createDirMap(commands string) map[string][]string {
	var m map[string][]string
	commandsLines := strings.Split(strings.TrimSpace(commands), "\n")

	for _, line := range commandsLines {
		fmt.Println(line)
	}

	return m
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	file := string(f)
	createDirMap(file)
}
