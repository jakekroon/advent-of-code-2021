package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalDepthIncreases := 0
	lastDepth := -1

	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())

		if lastDepth != -1 && currentDepth > lastDepth {
			totalDepthIncreases++
		}

		lastDepth = currentDepth
	}

	fmt.Printf("Total depth increases %v \n", totalDepthIncreases)
}
