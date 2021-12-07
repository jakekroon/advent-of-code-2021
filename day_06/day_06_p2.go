package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

func getSortedKeys(m map[int]int) []int {
	keys := make([]int, len(m))

	for i, k := range m {
		keys[i] = k
	}

	sort.Ints(keys)

	return keys
}

func getTotalFish(fishDays map[int]int) (totalFish int) {
	for _, total := range fishDays {
		totalFish += total
	}

	return
}

func main() {
	start := time.Now()

	file, _ := os.ReadFile("./puzzle_input.txt")
	fishInitStates, _ := sliceAtoi(strings.Split(strings.TrimSpace(string(file)), ","))

	fishDays := make(map[int]int)

	for i := 0; i < 9; i++ {
		fishDays[i] = 0
	}

	days := getSortedKeys(fishDays)

	for _, num := range fishInitStates {
		fishDays[num]++
	}

	for i := 0; i < 256; i++ {
		newFish := fishDays[0]
		fishDays[0] = 0

		for day := range days {
			if day != 0 {
				fishDays[day-1] = fishDays[day]
				fishDays[day] = 0
			}
		}

		fishDays[6] += newFish
		fishDays[8] += newFish
	}

	elapsed := time.Since(start)

	fmt.Printf("elapsed time %v\n", elapsed)
	fmt.Printf("Total fish spawned: %v\n", getTotalFish(fishDays))
}
