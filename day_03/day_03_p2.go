package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getBitComparison(bitStrings []string, index int) int {
	on := 0
	off := 0

	for _, bitString := range bitStrings {
		bit := strings.Split(bitString, "")[index]

		if bit == "1" {
			on++
		} else {
			off++
		}
	}

	if on > off {
		return 1
	} else if on == off {
		return 0
	}

	return -1
}

func filter(strings []string, index int, test func(string, int) bool) (filtered []string) {
	for _, string := range strings {
		if test(string, index) {
			filtered = append(filtered, string)
		}
	}

	return
}

func getOxygenRating(bitStrings []string) (oxygenRating int64) {
	count := 0

	for len(bitStrings) != 1 && count != len(bitStrings[0]) {
		comparator := getBitComparison(bitStrings, count)

		if comparator == 1 || comparator == 0 {
			bitStrings = filter(bitStrings, count, func(s string, i int) bool {
				return strings.Split(s, "")[i] == "1"
			})
		} else {
			bitStrings = filter(bitStrings, count, func(s string, i int) bool {
				return strings.Split(s, "")[i] == "0"
			})
		}

		count++
	}

	oxygenRating, _ = strconv.ParseInt(bitStrings[0], 2, 32)
	return
}

func getCO2Rating(bitStrings []string) (co2Rating int64) {
	count := 0

	for len(bitStrings) != 1 && count != len(bitStrings[0]) {
		comparator := getBitComparison(bitStrings, count)

		if comparator == 1 || comparator == 0 {
			bitStrings = filter(bitStrings, count, func(s string, i int) bool {
				return strings.Split(s, "")[i] == "0"
			})
		} else {
			bitStrings = filter(bitStrings, count, func(s string, i int) bool {
				return strings.Split(s, "")[i] == "1"
			})
		}

		count++
	}

	co2Rating, _ = strconv.ParseInt(bitStrings[0], 2, 32)
	return
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bitStrings []string
	for scanner.Scan() {
		bitStrings = append(bitStrings, scanner.Text())
	}

	oxygenRating := getOxygenRating(bitStrings)
	co2Rating := getCO2Rating(bitStrings)

	fmt.Printf("Life Support Rating %v\n", oxygenRating*co2Rating)
}
