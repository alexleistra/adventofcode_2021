package main

import (
	"bufio"
	"fmt"
	"log"
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
		xInc := 1
		if line.a.x > line.b.x {
			xInc = -1
		} else if line.a.x == line.b.x {
			xInc = 0
		}
		yInc := 1
		if line.a.y > line.b.y {
			yInc = -1
		} else if line.a.y == line.b.y {
			yInc = 0
		}
		x := line.a.x
		y := line.a.y
		for true {
			key := strconv.Itoa(x) + "," + strconv.Itoa(y)
			intersections[key] = intersections[key] + 1
			if highestIntersectionCount < intersections[key] {
				highestIntersectionCount = intersections[key]
			}

			if x == line.b.x && y == line.b.y {
				break
			}

			x = x + xInc
			y = y + yInc
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
