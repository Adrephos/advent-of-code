package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func (s *Stack) PopMany(num int) (Stack, bool) {
	if s.IsEmpty() {
		return *s, false
	} else {
		index := len(*s) - num  // Get the index of the top most element.
		element := (*s)[index:] // Index into the slice and obtain the elements.
		*s = (*s)[:index]       // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) AppendMany(stack Stack) {
	for _, item := range stack {
		*s = append(*s, item) // Simply append the new value to the end of the stack
	}
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getStacks(str string) []Stack {
	str = strings.ReplaceAll(str, "[", " ")
	str = strings.ReplaceAll(str, "]", " ")

	drawing := strings.Split(str, "\n")
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

	return stacks

}

func getTopCrates(comands []string, stacks []Stack) string {
	str := ""

	for _, command := range comands {
		commandParts := strings.Split(command, " ")
		amount, _ := strconv.Atoi(commandParts[1])
		origin, _ := strconv.Atoi(commandParts[3])
		destiny, _ := strconv.Atoi(commandParts[5])
		origin, destiny = origin-1, destiny-1

		for i := 0; i < amount; i++ {
			x, y := stacks[origin].Pop()
			if y == true {
				stacks[destiny].Push(x)
			}
		}

	}

	for _, stack := range stacks {
		z, k := stack.Pop()
		if k == true {
			str += z
		}
	}

	return str
}

func getTopCrates2(comands []string, stacks []Stack) string {
	str := ""

	for _, command := range comands {
		commandParts := strings.Split(command, " ")
		amount, _ := strconv.Atoi(commandParts[1])
		origin, _ := strconv.Atoi(commandParts[3])
		destiny, _ := strconv.Atoi(commandParts[5])
		origin, destiny = origin-1, destiny-1

		x, y := stacks[origin].PopMany(amount)
		if y == true {
			stacks[destiny].AppendMany(x)
		}

	}

	for _, stack := range stacks {
		z, k := stack.Pop()
		if k == true {
			str += z
		}
	}

	return str
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := string(f)

	split := strings.Split(file, "\n\n")

	commands := strings.TrimSpace(split[1])
	commandsArr := strings.Split(commands, "\n")

	cratesStacks := getStacks(split[0])
	cratesStacks2 := getStacks(split[0])

	topCrates := getTopCrates(commandsArr, cratesStacks)
	topCrates2 := getTopCrates2(commandsArr, cratesStacks2)

	fmt.Println("Puzzle 1:", topCrates)
	fmt.Println("Puzzle 2:", topCrates2)

}
