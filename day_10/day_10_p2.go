package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func reverseSlice(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func getIncompleteStack(chars []string) (stack []string) {
	for _, char := range chars {
		if char == "(" || char == "[" || char == "{" || char == "<" {
			stack = append(stack, char)
			continue
		}

		if (char == ")" || char == "]" || char == "}" || char == ">") && len(stack) == 0 {
			return []string{}
		}

		switch char {
		case ")":
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				continue
			}
			return []string{}
		case "]":
			if stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
				continue
			}
			return []string{}
		case "}":
			if stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
				continue
			}
			return []string{}
		case ">":
			if stack[len(stack)-1] == "<" {
				stack = stack[:len(stack)-1]
				continue
			}
			return []string{}
		}
	}

	return
}

func getScoreForStack(stack []string) (score int) {
	for _, char := range stack {
		score *= 5
		switch char {
		case "(":
			score += 1
		case "[":
			score += 2
		case "{":
			score += 3
		case "<":
			score += 4
		}
	}

	return
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scores []int
	for scanner.Scan() {
		incompleteStack := getIncompleteStack(strings.Split(scanner.Text(), ""))

		if len(incompleteStack) == 0 {
			continue
		}

		reverseSlice(incompleteStack)
		scores = append(scores, getScoreForStack(incompleteStack))
	}

	sort.Ints(scores)

	finalScore := scores[len(scores)/2]

	fmt.Printf("Total autocomplete score %v\n", finalScore)
}
