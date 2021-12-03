package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BitCount struct {
	on  int
	off int
}

func newBitCount() BitCount {
	return BitCount{
		on:  0,
		off: 0,
	}
}

func initBitCounts(bitSlice []string) (bitCounts []BitCount) {
	for _, num := range bitSlice {
		bitCount := newBitCount()

		if num == "1" {
			bitCount.on++
		} else {
			bitCount.off--
		}

		bitCounts = append(bitCounts, bitCount)
	}

	return
}

func getGammaRate(bitCounts []BitCount) (gammaRate string) {
	gammaRate = ""

	for _, bitCount := range bitCounts {
		if bitCount.on > bitCount.off {
			gammaRate += "1"
			continue
		}

		gammaRate += "0"
	}

	return
}

func getEpisolonRate(gammaRate string) (epsilonRate string) {
	epsilonRate = ""

	for _, char := range gammaRate {
		if char == '1' {
			epsilonRate += "0"
			continue
		}

		epsilonRate += "1"
	}

	return
}

func main() {
	file, _ := os.Open("./puzzle_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bitCounts []BitCount

	for scanner.Scan() {
		bitSlice := strings.Split(scanner.Text(), "")

		if bitCounts == nil {
			bitCounts = initBitCounts(bitSlice)
			continue
		}

		for index, bit := range bitSlice {
			if bit == "1" {
				bitCounts[index].on++
			} else {
				bitCounts[index].off++
			}
		}
	}

	gammaRateString := getGammaRate(bitCounts)
	gammaRate, _ := strconv.ParseInt(gammaRateString, 2, 32)
	epsilonRate, _ := strconv.ParseInt(getEpisolonRate(gammaRateString), 2, 32)

	fmt.Printf("Power consumption %v\n", gammaRate*epsilonRate)
}
