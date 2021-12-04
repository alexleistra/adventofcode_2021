package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/4
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numberPicker *NumberPicker

	first := true
	boardId := 0
	lineIndex := 0
	lineMax := 5
	currentBoardInput := [][]string{}
	boardLineReg := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		if first {
			// create number picker
			inputNumbers := strings.Split(inputLine, ",")
			numbers := make([]int, len(inputNumbers))
			for i, s := range inputNumbers {
				numbers[i], _ = strconv.Atoi(s)
			}
			numberPicker = NewNumberPicker(numbers)
			first = false
			continue
		}

		if boardLineReg.MatchString(inputLine) {
			findResult := boardLineReg.FindStringSubmatch(inputLine)
			currentBoardInput = append(currentBoardInput, findResult[1:])
			lineIndex++
			if lineIndex == lineMax {
				// register a new board
				currentBoard := make([][]int, len(currentBoardInput))
				for i, r := range currentBoardInput {
					for _, c := range r {
						num, _ := strconv.Atoi(c)
						currentBoard[i] = append(currentBoard[i], num)
					}
				}

				board := NewBingoBoard(boardId, currentBoard)
				numberPicker.register(board)
				boardId++

				// reset board building
				currentBoardInput = [][]string{}
				lineIndex = 0
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		if !numberPicker.hasNextNumber() {
			break
		}
		numberPicker.pickNextNumber()
		hasWinner, id, winningNumbers := numberPicker.findWinningBoard()
		if hasWinner {
			fmt.Printf("winning id: %v, final score: %v\n", id, winningNumbers)
			for _, board := range numberPicker.observerList {
				numberPicker.deregister(board)
			}
			break
		}
	}

}
