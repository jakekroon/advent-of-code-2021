package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ONE   int = 2
	SEVEN int = 3
	FOUR  int = 4
	EIGHT int = 7
)

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digitCount := 0
	for scanner.Scan() {
		signals := strings.Split(strings.Split(scanner.Text(), "|")[1], " ")

		for _, signal := range signals {
			signalLen := len(signal)
			if signalLen == ONE || signalLen == FOUR ||
				signalLen == SEVEN || signalLen == EIGHT {
				digitCount++
			}
		}
	}

	fmt.Printf("Digit count: %v\n", digitCount)
}
