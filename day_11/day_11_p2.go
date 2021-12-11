package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OctopusGrid [][]int

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

func getOctopusStartGrid(input string) (grid OctopusGrid) {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		octopi, _ := sliceAtoi(strings.Split(scanner.Text(), ""))
		grid = append(grid, octopi)
	}

	return
}

func processFlash(grid OctopusGrid, y, x int) (flashes int) {
	if grid[y][x] > 9 {
		flashes = 1

		// horizontal
		if x-1 >= 0 && grid[y][x-1] < 10 {
			grid[y][x-1]++
			if grid[y][x-1] > 9 {
				flashes += processFlash(grid, y, x-1)
			}
		}
		if x+1 < len(grid) && grid[y][x+1] < 10 {
			grid[y][x+1]++
			if grid[y][x+1] > 9 {
				flashes += processFlash(grid, y, x+1)
			}
		}

		// vertical
		if y-1 >= 0 && grid[y-1][x] < 10 {
			grid[y-1][x]++
			if grid[y-1][x] > 9 {
				flashes += processFlash(grid, y-1, x)
			}
		}
		if y+1 < len(grid[x]) && grid[y+1][x] < 10 {
			grid[y+1][x]++
			if grid[y+1][x] > 9 {
				flashes += processFlash(grid, y+1, x)
			}
		}

		// diagonal
		if x-1 >= 0 && y-1 >= 0 && grid[y-1][x-1] < 10 { // top left
			grid[y-1][x-1]++
			if grid[y-1][x-1] > 9 {
				flashes += processFlash(grid, y-1, x-1)
			}
		}
		if x+1 < len(grid) && y-1 >= 0 && grid[y-1][x+1] < 10 { // top right
			grid[y-1][x+1]++
			if grid[y-1][x+1] > 9 {
				flashes += processFlash(grid, y-1, x+1)
			}
		}
		if x+1 < len(grid) && y+1 < len(grid[x]) && grid[y+1][x+1] < 10 { // bottom right
			grid[y+1][x+1]++
			if grid[y+1][x+1] > 9 {
				flashes += processFlash(grid, y+1, x+1)
			}
		}
		if x-1 >= 0 && y+1 < len(grid[x]) && grid[y+1][x-1] < 10 { // bottom left
			grid[y+1][x-1]++
			if grid[y+1][x-1] > 9 {
				flashes += processFlash(grid, y+1, x-1)
			}
		}
	}

	return flashes
}

func processOctopusGrid(grid OctopusGrid) bool {
	for y := range grid {
		for x := range grid {
			grid[y][x]++
			if grid[y][x] == 10 {
				_ = processFlash(grid, y, x)
			}
		}
	}

	synced := true
	for y := range grid {
		for x := range grid {
			if grid[y][x] > 9 {
				grid[y][x] = 0
			}

			synced = synced && grid[y][x] == 0
		}
	}

	return synced
}

func main() {
	octopusGrid := getOctopusStartGrid("./puzzle_input.txt")

	synced := false
	day := 0
	for synced != true {
		synced = processOctopusGrid(octopusGrid)
		day++
	}

	fmt.Printf("%v\n", day)
}
