package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoTile struct {
	number int
	marked bool
}

type BingoBoard [][]BingoTile

func newBingoTile(number int) BingoTile {
	return BingoTile{
		number: number,
		marked: false,
	}
}

func stringsToTiles(strings []string) (bingoTiles []BingoTile) {
	for _, character := range strings {
		digit, _ := strconv.Atoi(character)
		bingoTiles = append(bingoTiles, newBingoTile(digit))
	}

	return
}

func stringsToInts(strings []string) (ints []int) {
	for _, character := range strings {
		digit, _ := strconv.Atoi(character)
		ints = append(ints, digit)
	}

	return
}

func initBingoGame(bingoData string) (bingoBoards []BingoBoard, randomNumbers []int) {
	file, _ := os.Open(bingoData)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bingoBoard BingoBoard

	for scanner.Scan() {
		readLine := scanner.Text()
		if randomNumbers == nil {
			randomNumbers = stringsToInts(strings.Split(readLine, ","))
			continue
		}

		if readLine != "" {
			bingoBoard = append(bingoBoard, stringsToTiles(strings.Fields(readLine)))
		}

		if len(bingoBoard) == 5 {
			bingoBoards = append(bingoBoards, bingoBoard)
			bingoBoard = nil
		}
	}

	return
}

func isBingo(bingoBoard BingoBoard) bool {
	vertWinner := true
	hozWinner := true

	for i := 0; i < 5; i++ { // row
		for j := 0; j < 5; j++ { // column
			vertWinner = vertWinner && bingoBoard[i][j].marked
		}

		if vertWinner {
			return vertWinner
		}

		vertWinner = true
	}

	for i := 0; i < 5; i++ { // column
		for j := 0; j < 5; j++ { // row
			hozWinner = hozWinner && bingoBoard[j][i].marked
		}

		if hozWinner {
			return hozWinner
		}

		hozWinner = true
	}

	return false
}

func findWinningBoard(bingoBoards []BingoBoard, numbers []int) (winningBoard BingoBoard, winningNumber int) {
	for _, num := range numbers {
		for _, board := range bingoBoards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					// i row, j column
					if board[i][j].number == num {
						board[i][j].marked = true
					}
				}
			}

			if isBingo(board) {
				winningBoard = board
				winningNumber = num
				return
			}
		}
	}

	return
}

func sumUnmarkedNumbers(bingoBoard BingoBoard) int {
	unmarkedNumbers := 0

	for _, row := range bingoBoard {
		for _, tile := range row {
			if !tile.marked {
				unmarkedNumbers += tile.number
			}
		}
	}

	return unmarkedNumbers
}

func main() {
	bingoBoards, numbers := initBingoGame("./puzzle_input.txt")

	winningBoard, winningNumber := findWinningBoard(bingoBoards, numbers)

	if winningBoard != nil && winningNumber != 0 {
		fmt.Printf("Final Score: %v\n", winningNumber*sumUnmarkedNumbers(winningBoard))
	}
}
