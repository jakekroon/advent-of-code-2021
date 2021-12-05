package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func removeBoards(bingoBoards []BingoBoard, indexesToRemove []int) []BingoBoard {
	sort.Slice(indexesToRemove, func(i, j int) bool {
		return indexesToRemove[i] > indexesToRemove[j]
	})

	for _, pos := range indexesToRemove {
		bingoBoards = append(bingoBoards[:pos], bingoBoards[pos+1:]...)
	}

	return bingoBoards
}

func findWinningBoards(bingoBoards []BingoBoard, numbers []int) (winningBoards []BingoBoard, winningNumbers []int) {
	var boardsThatWon []int

	for _, num := range numbers {
		bingoBoards = removeBoards(bingoBoards, boardsThatWon)
		boardsThatWon = nil

		for boardPos, board := range bingoBoards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					// i row, j column
					if board[i][j].number == num {
						board[i][j].marked = true
					}
				}
			}

			if isBingo(board) {
				winningBoards = append(winningBoards, board)
				boardsThatWon = append(boardsThatWon, boardPos)
				winningNumbers = append(winningNumbers, num)
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

	winningBoards, winningNumbers := findWinningBoards(bingoBoards, numbers)

	if winningBoards != nil && winningNumbers != nil {
		lastWinningNumber := winningNumbers[len(winningNumbers)-1]
		lastWinningBoard := winningBoards[len(winningBoards)-1]

		fmt.Printf("Final Score: %v\n", lastWinningNumber*sumUnmarkedNumbers(lastWinningBoard))
	}
}
