package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")

	file := string(f)

	rucksack := strings.Split(file, "\n")

	total := 0

	for _, items := range rucksack {
		if len(items) != 0 {
			half := (len(items) / 2)
			compartment1 := items[0:half]
			compartment2 := items[half:]
			priority := 0
			for _, item1 := range compartment1 {
				for _, item2 := range compartment2 {
					if item1 == item2 {
						ascii := int(item1)
						if ascii >= 65 && ascii <= 90 {
							priority = ascii - 38
						} else {
							priority = ascii - 96
						}
					}
				}
			}

			total += priority
		}

	}
	fmt.Println(total)
}
