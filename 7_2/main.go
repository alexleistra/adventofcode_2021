package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/7
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	positions := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputList := strings.Split(input, ",")
		for _, inputAge := range inputList {
			position, err := strconv.Atoi(inputAge)
			if err != nil {
				log.Fatal(err)
			}
			positions = append(positions, position)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// get the range of crab positions
	var max int = positions[0]
	var min int = positions[0]
	for _, value := range positions {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	possibleRallyPoints := max - min

	// calculate the fuel spend for each possible position
	minFuel := math.MaxInt32
	rallyPoint := 0
	for i := 0; i < possibleRallyPoints; i++ {
		fuelTotal := 0
		for _, p := range positions {
			distance := int(math.Abs(float64(p - i)))
			fuel := distance * (distance + 1) / 2
			fuelTotal += fuel
		}
		if fuelTotal < minFuel {
			minFuel = fuelTotal
			rallyPoint = i
		}
	}

	fmt.Println(rallyPoint)
	fmt.Println(minFuel)
}
