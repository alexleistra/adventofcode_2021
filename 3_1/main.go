package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// https://adventofcode.com/2021/day/3
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var gammaBits, epsilonBits []int
	var inputBitsCounts []int
	var inputCount int
	for scanner.Scan() {
		bitsInput := scanner.Text()
		inputCount++
		for i, c := range bitsInput {
			if c == '1' {
				if len(inputBitsCounts) < i+1 {
					inputBitsCounts = append(inputBitsCounts, 1)
				} else {
					inputBitsCounts[i] = inputBitsCounts[i] + 1
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, c := range inputBitsCounts {
		if c > inputCount/2 {
			gammaBits = append(gammaBits, 1)
			epsilonBits = append(epsilonBits, 0)
		} else {
			gammaBits = append(gammaBits, 0)
			epsilonBits = append(epsilonBits, 1)
		}
	}

	var gamma, epsilon int
	for i := len(gammaBits) - 1; i >= 0; i-- {
		pos := len(gammaBits) - 1 - i
		gamma += bitWithPositionToDecimal(gammaBits[i], pos)
		epsilon += bitWithPositionToDecimal(epsilonBits[i], pos)
	}

	fmt.Println(gamma * epsilon)
}

func bitWithPositionToDecimal(bitInput, pos int) int {
	if bitInput == 1 {
		return int(math.Pow(2, float64(pos)))
	}
	return 0
}
