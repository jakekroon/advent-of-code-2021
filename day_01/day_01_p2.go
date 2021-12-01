package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./puzzle_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalDepthIncreases := 0
	lastSlidingWindowSum := -1
	var depthSlidingWindow []int

	for scanner.Scan() {
		currentDepth, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

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
