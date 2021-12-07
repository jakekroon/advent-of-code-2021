package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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
	crabs, _ := sliceAtoi(strings.Split(strings.TrimSpace(string(file)), ","))

	sort.Ints(crabs)

	midPoint := crabs[len(crabs)/2]

	var totalMoves = 0.0
	for _, crab := range crabs {
		totalMoves += math.Abs(float64(crab) - float64(midPoint))
	}

	fmt.Printf("total moves %v\n", int(totalMoves))
}
