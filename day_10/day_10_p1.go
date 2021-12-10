package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findCorruptedCharacter(chars []string) string {
	var stack []string

	for _, char := range chars {
		if char == "(" || char == "[" || char == "{" || char == "<" {
			stack = append(stack, char)
			continue
		}

		if (char == ")" || char == "]" || char == "}" || char == ">") && len(stack) == 0 {
			return char
		}

		switch char {
		case ")":
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				continue
			}
			return char
		case "]":
			if stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
				continue
			}
			return char
		case "}":
			if stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
				continue
			}
			return char
		case ">":
			if stack[len(stack)-1] == "<" {
				stack = stack[:len(stack)-1]
				continue
			}
			return char
		}
	}

	return ""
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		cc := findCorruptedCharacter(strings.Split(scanner.Text(), ""))

		switch cc {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}

	fmt.Printf("Total syntax error score %v\n", score)
}
