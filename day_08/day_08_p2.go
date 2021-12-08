package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func mapKey(m map[string]string, value string) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return
		}
	}
	return ""
}

func containsSegments(s, substr string) bool {
	contains := true

	if substr == "" {
		return false
	}

	for _, r := range strings.Split(substr, "") {
		contains = contains && strings.Contains(s, r)
	}
	return contains
}

func preProcessPattern(signalPattern []string) map[string]string {
	sort.Slice(signalPattern, func(i, j int) bool {
		return len(signalPattern[i]) < len(signalPattern[j])
	})

	for i := range signalPattern {
		signalPattern[i] = SortString(signalPattern[i])
	}

	signalMap := map[string]string{
		signalPattern[0]: "1",
		signalPattern[1]: "7",
		signalPattern[2]: "4",
		signalPattern[9]: "8",
	}

	for len(signalMap) != 10 {
		for _, signal := range signalPattern {
			if len(signal) == 6 {
				if containsSegments(signal, mapKey(signalMap, "4")) && signal != mapKey(signalMap, "4") {
					signalMap[signal] = "9"
				} else if containsSegments(signal, mapKey(signalMap, "1")) && signal != mapKey(signalMap, "1") {
					signalMap[signal] = "0"
				} else if mapKey(signalMap, "9") != "" && mapKey(signalMap, "0") != "" {
					signalMap[signal] = "6"
				}
			} else if len(signal) == 5 {
				if containsSegments(signal, mapKey(signalMap, "1")) && signal != mapKey(signalMap, "1") {
					signalMap[signal] = "3"
				} else if containsSegments(mapKey(signalMap, "6"), signal) && signal != mapKey(signalMap, "6") {
					signalMap[signal] = "5"
				} else if mapKey(signalMap, "3") != "" && mapKey(signalMap, "5") != "" {
					signalMap[signal] = "2"
				}
			}
		}
	}

	return signalMap
}

func processOutputSignal(signalOutput []string, signalMap map[string]string) int {
	for i := range signalOutput {
		signalOutput[i] = SortString(signalOutput[i])
	}

	var output strings.Builder
	for _, signal := range signalOutput {
		output.WriteString(signalMap[signal])
	}

	outputNum, _ := strconv.Atoi(output.String())
	return outputNum
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	outputTotal := 0
	for scanner.Scan() {
		signalData := strings.Split(scanner.Text(), " | ")
		signalMap := preProcessPattern(strings.Split(signalData[0], " "))
		outputTotal += processOutputSignal(strings.Split(signalData[1], " "), signalMap)
	}

	fmt.Printf("Digit count: %v\n", outputTotal)
}
