package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	sort.Ints(positions)
	m := len(positions) / 2
	rallyPoint := positions[m]
	fmt.Println(rallyPoint)

	count := 0
	for _, p := range positions {
		distance := int(math.Abs(float64(p - rallyPoint)))
		// fmt.Printf("move from %v to %v: %v\r\n", p, rallyPoint, distance)
		count += distance
	}

	fmt.Println(count)
}
