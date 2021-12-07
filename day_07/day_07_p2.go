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

// see: https://nrich.maths.org/2478
func gauss(num int) int {
	return (num * (1 + num)) / 2
}

func main() {
	file, _ := os.ReadFile("./puzzle_input.txt")
	crabs, _ := sliceAtoi(strings.Split(strings.TrimSpace(string(file)), ","))

	sort.Ints(crabs)

	max := crabs[len(crabs)-1]

	minFuel := math.MaxInt
	for i := 0; i <= max; i++ {
		currFuel := 0
		for _, crab := range crabs {
			seqCnt := int(math.Abs(float64(i - crab)))
			currFuel += gauss(seqCnt)
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	fmt.Printf("Minimum fuel %v\n", int(minFuel))
}
