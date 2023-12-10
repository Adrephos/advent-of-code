package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func seedsAndConversions(path string) ([]string, []string) {
	data, _ := os.ReadFile(path)
	txt := strings.TrimSpace(string(data))
	parts := strings.Split(txt, "\n\n")
	seeds := strings.Split(parts[0], ":")[1]
	seedsArr := strings.Split(strings.TrimSpace(seeds), " ")

	return seedsArr, parts[1:]
}

func isInRange(str string, sourceNum int, reversed bool) (int, bool) {
	str = strings.TrimSpace(str)
	rangeParts := strings.Split(str, " ")

	destStart, _ := strconv.Atoi(rangeParts[0])
	srcStart, _ := strconv.Atoi(rangeParts[1])
	rangeLen, _ := strconv.Atoi(rangeParts[2])

	if (srcStart <= sourceNum && sourceNum <= (srcStart+rangeLen-1)) && !reversed {
		return destStart + (sourceNum - srcStart), true
	}

	if (destStart <= sourceNum && sourceNum <= (destStart+rangeLen-1)) && reversed {
		return srcStart + (sourceNum - destStart), true
	}

	return 0, false
}

func convert(source int, convertionMap string, reversed bool) int {
	convertionMap = strings.Split(convertionMap, ":")[1]
	convertionMap = strings.TrimSpace(convertionMap)

	convertionRanges := strings.Split(convertionMap, "\n")

	convertion, inRange := 0, false

	for _, convertionRange := range convertionRanges {
		convertion, inRange = isInRange(convertionRange, source, reversed)
		if inRange {
			return convertion
		}
	}

	return source
}

func findLocation(seed string, convertionsMaps []string) int {
	source, _ := strconv.Atoi(seed)
	for _, convertionMap := range convertionsMaps {
		source = convert(source, convertionMap, false)
	}
	return source
}

func getMinLocation(seedsArr []string, convertionsMaps []string) int {
	minLocation := 9223372036854775807

	for _, seedStr := range seedsArr {
		location := findLocation(seedStr, convertionsMaps)
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func findSeed(location int, convertionsMaps []string) string {
	seed := strconv.Itoa(location)
	for i := len(convertionsMaps) - 1; i >= 0; i-- {
		convertionMap := convertionsMaps[i]
		source, _ := strconv.Atoi(seed)
		seed = strconv.Itoa(convert(source, convertionMap, true))
	}

	return seed
}

func inSeeds(seed string, seedArr []string) bool {
	seedNum, _ := strconv.Atoi(seed)

	for i := 0; i < len(seedArr)-1; i += 2 {
		rangeStartStr, rangeLenStr := seedArr[i], seedArr[i+1]
		rangeStart, _ := strconv.Atoi(rangeStartStr)
		rangeLen, _ := strconv.Atoi(rangeLenStr)
		if rangeStart <= seedNum && seedNum <= (rangeStart+rangeLen-1) {
			return true
		}
	}
	return false
}

func findMinLocation(seedArr []string, convertionsMaps []string) int {
	location := 0

	for {
		seed := findSeed(location, convertionsMaps)
		if inSeeds(seed, seedArr) { break }
		location++
	}

	return location
}

func main() {
	seedArr, convertionsMaps := seedsAndConversions("./input2.txt")
	minLoc := getMinLocation(seedArr, convertionsMaps)
	fmt.Println("First Puzzle: ", minLoc)
	minLoc = findMinLocation(seedArr, convertionsMaps)
	fmt.Println("Second Puzzle: ", minLoc)
}
