package main

import (
	"bufio"
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
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	var heightMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		heights, _ := sliceAtoi(strings.Split(scanner.Text(), ""))
		heightMap = append(heightMap, heights)
	}

	riskLevel := 0

	for i := range heightMap {
		for j := range heightMap[i] {
			if (i == 0 || heightMap[i][j] < heightMap[i-1][j]) &&
				(j == 0 || heightMap[i][j] < heightMap[i][j-1]) &&
				(i == len(heightMap)-1 || heightMap[i][j] < heightMap[i+1][j]) &&
				(j == len(heightMap[i])-1 || heightMap[i][j] < heightMap[i][j+1]) {
				riskLevel += heightMap[i][j] + 1
			}
		}
	}

	fmt.Printf("Risk Level %v\n", riskLevel)
}
