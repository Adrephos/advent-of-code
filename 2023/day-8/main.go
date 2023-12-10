package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Node struct {
	left string
	right string
}

func GCD(a, b int) int {
      if b == 0 { return a }
      return GCD(b, a % b)
}

func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func getData(path string) (string, map[string]Node, []string) {
	data, _ := os.ReadFile(path)
	instructions, nodesStr, _ := strings.Cut(string(data), "\n\n")
	nodes := strings.Split(strings.TrimSpace(nodesStr), "\n")
	nodeMap := make(map[string]Node)
	startingNodes := []string{}

	for _, nodeLine := range nodes {
		nodeLine = strings.TrimSpace(nodeLine)
		nodeName, nodeConections, _ := strings.Cut(nodeLine, " = ")
		left, right, _ := strings.Cut(nodeConections, ", ")
		node := Node{ left[1:], right[:len(right)-1] }
		nodeMap[nodeName] = node

		if nodeName[2] == 'A' { startingNodes = append(startingNodes, nodeName) }
	}
	
	return strings.TrimSpace(instructions), nodeMap, startingNodes
}

func countSteps(instructions string, nodes map[string]Node, startNode string) int {
	steps, i := 0, 0
	for startNode[2] != 'Z' {
		instruction := instructions[i]

		node := nodes[startNode]

		if instruction == 'L' { startNode = node.left
		} else { startNode = node.right }

		if i == len(instructions)-1 { i = 0 
		} else { i++ }
		steps++
	}
	return steps
}

func secondPuzzle(instructions string, nodes map[string]Node, startingNodes []string) int {
	steps := []int{}
	for _, startNode := range startingNodes {
		steps = append(steps,countSteps(instructions, nodes, startNode))
	}
	return LCM(steps[0], steps[1], steps[2:]...)
}

func main() {
	startTime := time.Now()
	instructions, nodes, startingNodes := getData("./input.txt")
	fmt.Println("First Puzzle:", countSteps(instructions, nodes, "AAA"))
	fmt.Printf("Time: %vµs\n", time.Since(startTime).Microseconds())

	startTime = time.Now()
	fmt.Println("Second Puzzle:", secondPuzzle(instructions, nodes, startingNodes))
	fmt.Printf("Time: %vµs\n", time.Since(startTime).Microseconds())
}
