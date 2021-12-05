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

type vector struct {
	x int
	y int
}

type line struct {
	a vector
	b vector
}

// https://adventofcode.com/2021/day/5
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []line
	for scanner.Scan() {
		input := scanner.Text()
		inputLine := strings.Split(input, " -> ")
		var line line
		line.a = makeVector(inputLine[0])
		line.b = makeVector(inputLine[1])
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	intersections := make(map[string]int)
	var highestIntersectionCount int
	for _, line := range lines {
		if line.a.x == line.b.x {
			// vertical lines
			min := int(math.Min(float64(line.a.y), float64(line.b.y)))
			max := int(math.Max(float64(line.a.y), float64(line.b.y)))

			for i := min; i < max+1; i++ {
				v := vector{x: line.a.x, y: i}
				key := strconv.Itoa(v.x) + "," + strconv.Itoa(v.y)
				intersections[key] = intersections[key] + 1
				if highestIntersectionCount < intersections[key] {
					highestIntersectionCount = intersections[key]
				}
			}

		} else if line.a.y == line.b.y {
			// horizontal lines
			min := int(math.Min(float64(line.a.x), float64(line.b.x)))
			max := int(math.Max(float64(line.a.x), float64(line.b.x)))

			for i := min; i < max+1; i++ {
				v := vector{x: i, y: line.a.y}
				key := strconv.Itoa(v.x) + "," + strconv.Itoa(v.y)
				intersections[key] = intersections[key] + 1
				if highestIntersectionCount < intersections[key] {
					highestIntersectionCount = intersections[key]
				}
			}
		}

	}

	var count int
	for _, val := range intersections {
		if val > 1 {
			count++
		}
	}

	fmt.Println(highestIntersectionCount)
	fmt.Println(count)
}

func makeVector(inputLineVector string) vector {
	coords := strings.Split(inputLineVector, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}
	return vector{x: x, y: y}
}
