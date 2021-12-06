package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(ss []string) ([]int, error) {
	is := make([]int, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return is, err
		}
		is = append(is, i)
	}

	return is, nil
}

func main() {
	file, _ := os.ReadFile("./puzzle_input.txt")
	fishStates, _ := sliceAtoi(strings.Split(strings.TrimSpace(string(file)), ","))

	for i := 0; i < 80; i++ {
		var newFishStates []int
		for j := range fishStates {
			if fishStates[j] == 0 {
				fishStates[j] = 6
				fishStates = append(fishStates, 8)
				continue
			}
			fishStates[j]--
		}

		fishStates = append(fishStates, newFishStates...)
		newFishStates = nil
	}

	fmt.Printf("Total fish spawned: %v\n", len(fishStates))
}
