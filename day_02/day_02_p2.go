package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y, aim := 0, 0, 0

	for scanner.Scan() {
		movement := strings.Split(scanner.Text(), " ")
		direction := movement[0]
		units, _ := strconv.Atoi(movement[1])

		switch direction {
		case "forward":
			x += units
			y += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}

	fmt.Printf("Final position %v\n", x*y)
}
