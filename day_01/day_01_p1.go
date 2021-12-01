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
	lastDepth := -1

	for scanner.Scan() {
		currentDepth, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		if lastDepth != -1 && currentDepth > lastDepth {
			totalDepthIncreases++
		}

		lastDepth = currentDepth
	}

	fmt.Printf("Total depth increases %v \n", totalDepthIncreases)
}
