package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var openingChars = "([{<"

var illegalPointsTable = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var autocompletePointsTable = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var closingChar = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

// https://adventofcode.com/2021/day/10
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	illegalCharsPoints := 0
	lineAutocompleteScores := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineChars := strings.Split(line, "")
		stack := []string{}
		illegalLineCharPoints := 0
		for _, c := range lineChars {
			if strings.Contains(openingChars, c) {
				stack = append(stack, c)
			} else {
				n := len(stack) - 1
				topOpening := stack[n]
				stack = stack[:n]
				if c != closingChar[topOpening] {
					illegalLineCharPoints += illegalPointsTable[c]
				}
			}
		}

		illegalCharsPoints += illegalLineCharPoints

		if illegalLineCharPoints == 0 {
			// autocomplete
			lineAutocompleteScore := 0
			for len(stack) > 0 {
				n := len(stack) - 1
				topOpening := stack[n]
				stack = stack[:n]
				lineAutocompleteScore = lineAutocompleteScore*5 + autocompletePointsTable[closingChar[topOpening]]
			}
			lineAutocompleteScores = append(lineAutocompleteScores, lineAutocompleteScore)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(lineAutocompleteScores)
	autocompletePoints := lineAutocompleteScores[int(math.Floor(float64(len(lineAutocompleteScores))/2.0))]

	fmt.Println("illegal chars points")
	fmt.Println(illegalCharsPoints)
	fmt.Println()
	fmt.Println("auto complete points")
	fmt.Println(autocompletePoints)
}
