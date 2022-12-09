package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := string(f)

	split := strings.Split(file, "\n\n")

	split[0] = strings.ReplaceAll(split[0], "[", " ")
	split[0] = strings.ReplaceAll(split[0], "]", " ")

	drawing := strings.Split(split[0], "\n")
	drawing = reverseArray(drawing)
	num := len(strings.ReplaceAll(drawing[0], " ", ""))
	size := len(drawing[0])

	var stacks []Stack
	for i := 0; i < num; i++ {
		var stack Stack
		stacks = append(stacks, stack)

	}

	for k := 1; k < len(drawing); k++ {
		line := drawing[k]
		for i, j := 1, 0; i < size; i, j = i+4, j+1 {
			rowStr := string(line[i])
			if rowStr != " " {
				stacks[j].Push(rowStr)
			}
		}
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Println(i+1, "=>", stacks[i])
	}

}
