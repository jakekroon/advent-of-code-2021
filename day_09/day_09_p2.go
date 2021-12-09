package main

import (
	"bufio"
	"fmt"
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

func findMinimumDepths(heightMap [][]int) (minDepthCoords [][]int) {
	for i := range heightMap {
		for j := range heightMap[i] {
			if (i == 0 || heightMap[i][j] < heightMap[i-1][j]) &&
				(j == 0 || heightMap[i][j] < heightMap[i][j-1]) &&
				(i == len(heightMap)-1 || heightMap[i][j] < heightMap[i+1][j]) &&
				(j == len(heightMap[i])-1 || heightMap[i][j] < heightMap[i][j+1]) {
				minDepthCoords = append(minDepthCoords, []int{i, j})
			}
		}
	}

	return
}

func floodFill(heightMap [][]int, x, y int) (size int) {
	if heightMap[x][y] == 9 {
		return 1
	}

	heightMap[x][y] = 9
	size = 1

	if x-1 >= 0 && heightMap[x-1][y] != 9 {
		size += floodFill(heightMap, x-1, y)
	}
	if x+1 < len(heightMap) && heightMap[x+1][y] != 9 {
		size += floodFill(heightMap, x+1, y)
	}
	if y-1 >= 0 && heightMap[x][y-1] != 9 {
		size += floodFill(heightMap, x, y-1)
	}
	if y+1 < len(heightMap[x]) && heightMap[x][y+1] != 9 {
		size += floodFill(heightMap, x, y+1)
	}

	return size
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

	minDepthsCoords := findMinimumDepths(heightMap)

	var sizes []int
	for _, coords := range minDepthsCoords {
		sizes = append(sizes, floodFill(heightMap, coords[0], coords[1]))
	}

	sort.Ints(sizes)

	result := sizes[len(sizes)-3] * sizes[len(sizes)-2] * sizes[len(sizes)-1]

	fmt.Printf("Result is %v\n", result)
}
