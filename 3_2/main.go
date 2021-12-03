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
	var inputValues []string
	for scanner.Scan() {
		bitsInput := scanner.Text()
		inputValues = append(inputValues, bitsInput)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	number := findLifeSupportRating(inputValues)

	fmt.Println(number)
}

func findLifeSupportRating(inputValues []string) int {
	pos := 0
	mostCommonCriteria, _ := getBitCriterias(inputValues, pos)

	// split slice between the most common and least common
	var mostCommonValues, leastCommonValues []string
	for i := 0; i < len(inputValues); i++ {
		if string(inputValues[i][pos]) == mostCommonCriteria {
			mostCommonValues = append(mostCommonValues, inputValues[i])
		} else {
			leastCommonValues = append(leastCommonValues, inputValues[i])
		}
	}

	// next pos
	pos++

	// get most common and least common numbers
	mostCommonCriteriaResult := findValueWithMostCommonCriteria(mostCommonValues, pos)
	leastCommonCriteriaResult := findValueWithLeastCommonCriteria(leastCommonValues, pos)

	var oxygen int
	for i := len(mostCommonCriteriaResult) - 1; i >= 0; i-- {
		parsePos := len(mostCommonCriteriaResult) - 1 - i
		oxygen += bitWithPositionToDecimal(string(mostCommonCriteriaResult[i]), parsePos)
	}

	var c02 int
	for i := len(mostCommonCriteriaResult) - 1; i >= 0; i-- {
		parsePos := len(mostCommonCriteriaResult) - 1 - i
		c02 += bitWithPositionToDecimal(string(leastCommonCriteriaResult[i]), parsePos)
	}

	// result
	return oxygen * c02
}

func findValueWithMostCommonCriteria(values []string, pos int) string {
	criteria, _ := getBitCriterias(values, pos)
	values = removeValues(values, pos, criteria)

	if len(values) == 1 {
		return values[0]
	}

	pos++
	return findValueWithMostCommonCriteria(values, pos)
}

func findValueWithLeastCommonCriteria(values []string, pos int) string {
	_, criteria := getBitCriterias(values, pos)
	values = removeValues(values, pos, criteria)

	if len(values) == 1 {
		return values[0]
	}

	pos++
	return findValueWithLeastCommonCriteria(values, pos)
}

func removeValues(values []string, pos int, criteria string) []string {
	for i := 0; i < len(values); i++ {
		if string(values[i][pos]) != criteria {
			values = append(values[:i], values[i+1:]...)
			i--
		}
	}
	return values
}

func getBitCriterias(values []string, pos int) (string, string) {
	// find the bit criteria
	var oneCount int
	var zeroCount int
	halfLen := len(values) / 2
	for i := 0; i < len(values); i++ {
		if string(values[i][pos]) == "1" {
			oneCount++
		} else {
			zeroCount++
		}

		// stop if we already found the most common
		// the most common or least common will be the bit criteria
		if oneCount > halfLen || zeroCount > halfLen {
			break
		}
	}

	// the bit criteria
	// default values if there is a tie
	mostCommonCriteria := "1"
	leastCommonCriteria := "0"
	if oneCount < zeroCount {
		mostCommonCriteria = "0"
		leastCommonCriteria = "1"
	}

	return mostCommonCriteria, leastCommonCriteria
}

func bitWithPositionToDecimal(bitInput string, pos int) int {
	if bitInput == "1" {
		return int(math.Pow(2, float64(pos)))
	}
	return 0
}
