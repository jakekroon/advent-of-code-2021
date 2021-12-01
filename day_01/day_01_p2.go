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
	lastSlidingWindowSum := -1
	var depthSlidingWindow []int

	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())

		depthSlidingWindow = append(depthSlidingWindow, currentDepth)

		if len(depthSlidingWindow) == 3 {
			currentSlidingWindowSum := sum(depthSlidingWindow)

			if lastSlidingWindowSum != -1 && currentSlidingWindowSum > lastSlidingWindowSum {
				totalDepthIncreases++
			}

			lastSlidingWindowSum = currentSlidingWindowSum
			depthSlidingWindow = depthSlidingWindow[1:]
		}
	}

	fmt.Printf("Total depth increases %v\n", totalDepthIncreases)
}

func sum(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}

	return sum
}
