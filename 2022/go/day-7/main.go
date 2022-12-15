package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func createDirMap(commands string) map[string][]string {
	var m map[string][]string
	commandsLines := strings.Split(strings.TrimSpace(commands), "\n")

	var actualDir []string
	for _, line := range commandsLines {
		if strings.HasPrefix(line, "$") {
			commands := strings.Split(strings.TrimSpace(line), " ")
			if commands[1] == "cd" {
				if commands[2] == "/" {
					actualDir = append(actualDir, "/")
				}
				if commands[2] == ".." {
					lngt := len(actualDir)
					actualDir = actualDir[:lngt-2]
				} else if commands[2] != "/" {
					actualDir = append(actualDir, commands[2])
					actualDir = append(actualDir, "/")
				}
				actualDirStr := strings.Join(actualDir, "")
				fmt.Println(actualDir)
				fmt.Println(actualDirStr)
			}
			if commands[1] == "ls" {
			}
		}
	}

	return m
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	file := string(f)
	createDirMap(file)
}
