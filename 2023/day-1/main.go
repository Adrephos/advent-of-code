package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCalibrationDoc(path string) []string {
	data, _ := os.ReadFile(path)
	txt := strings.Trim(string(data), "\n")
	calibrationDoc := strings.Split(txt, "\n")
	return calibrationDoc
}

func getCalibrationValue(str string) int {
	i, j, num := 0, len(str)-1, make([]byte, 2)

	for i <= len(str)-1 && j >= 0 {
		lVal, rVal := str[i], str[j]

		if lVal >= '0' && lVal <= '9' {
			if num[0] == 0 { num[0] = lVal }
		}
		if rVal >= '0' && rVal <= '9' {
			if num[1] == 0 { num[1] = rVal }
		}
		i++; j--
	}
	numInt, _ := strconv.Atoi(string(num))
	return numInt
}

func stringFixNums(str string) string {
	bytes := []byte(str)
	nums := map[string]byte{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for key, value := range nums {
		index := strings.Index(str, key)
		if index >= 0 {
			bytes[index] = value
		}
		index = strings.LastIndex(str, key)
		if index >= 0 {
			bytes[index] = value
		}
	}
	return string(bytes)
}

// Part one
func calibrateSum(calibrationDoc []string) int {
	sum := 0
	for _, calibrationLine := range calibrationDoc {
		calibrationValue := getCalibrationValue(calibrationLine)
		sum += calibrationValue
	}

	return sum
}

// Part two
func calibrateSumFixed(calibrationDoc []string) int {
	sum := 0
	for _, calibrationLine := range calibrationDoc {
		calibrationLine = stringFixNums(calibrationLine)
		calibrationValue := getCalibrationValue(calibrationLine)
		sum += calibrationValue
	}

	return sum
}

func main() {
	calibrationDoc := getCalibrationDoc("./input.txt")
	fmt.Println(calibrateSum(calibrationDoc))
	fmt.Println(calibrateSumFixed(calibrationDoc))
}
