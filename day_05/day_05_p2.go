package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int

func NewGrid(width, height int) (grid Grid) {
	for i := 0; i < (height + 1); i++ {
		var col []int
		for i := 0; i < (width + 1); i++ {
			col = append(col, 0)
		}
		grid = append(grid, col)
	}

	return
}

func convertPointStringsToInts(pointsString string) (points []int) {
	for _, point := range strings.Split(pointsString, ",") {
		pointNum, _ := strconv.Atoi(point)
		points = append(points, pointNum)
	}

	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

// Why did I not just use a map and count the points :(((((
func seedGrid(points [][][]int, grid Grid) Grid {
	for _, point_set := range points {
		start := point_set[0]
		end := point_set[1]

		if start[1] == end[1] {
			// seed rows
			for i := min(start[0], end[0]); i < (max(start[0], end[0]) + 1); i++ {
				grid[start[1]][i]++
			}
		} else if start[0] == end[0] {
			// seed columns
			for i := min(start[1], end[1]); i < (max(start[1], end[1]) + 1); i++ {
				grid[i][start[0]]++
			}
		} else {
			// seed diagonals
			seedX := min(start[0], end[0])

			forward := (start[1]-end[1])/(start[0]-end[0]) == -1
			if !forward {
				seedX = max(start[0], end[0])
			}

			for i := max(start[1], end[1]); i > (min(start[1], end[1]) - 1); i-- {
				grid[i][seedX]++

				if forward {
					seedX++
				} else {
					seedX--
				}
			}
		}
	}

	return grid
}

func processGrid(grid Grid) int {
	points := 0
	for _, row := range grid {
		for _, point := range row {
			if point > 1 {
				points++
			}
		}
	}

	return points
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points [][][]int

	width := 0
	height := 0

	for scanner.Scan() {
		pointsStrings := strings.Split(scanner.Text(), " -> ")
		start := convertPointStringsToInts(pointsStrings[0])
		end := convertPointStringsToInts(pointsStrings[1])

		points = append(points, [][]int{start, end})
		width = max(width, max(start[0], end[0]))
		height = max(height, max(start[1], end[1]))
	}

	grid := NewGrid(width, height)
	grid = seedGrid(points, grid)

	intersections := processGrid(grid)

	fmt.Printf("Points of intersection: %v\n", intersections)

	// At least I can add a cool accompanying md?
	for _, row := range grid {
		fmt.Printf("%v\n", row)
	}
}
