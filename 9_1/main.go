package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vector struct {
	x int
	y int
}

type bounds struct {
	min vector
	max vector
}

// https://adventofcode.com/2021/day/9
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	heightmap := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineValues := strings.Split(line, "")
		row := []int{}
		for _, value := range lineValues {
			x, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			row = append(row, x)
		}
		heightmap = append(heightmap, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(heightmap) == 0 {
		fmt.Println("no rows")
		return
	}

	// min x,y and max x,y
	heightmapBounds := bounds{min: vector{x: 0, y: 0}, max: vector{x: len(heightmap[0]) - 1, y: len(heightmap) - 1}}

	lowestPositions := make(map[vector]int)

	for y, row := range heightmap {
		for x, height := range row {
			isLow, lowPosition, lowHeight := findLowestPositionFromPoint(heightmap, height, vector{x: x, y: y}, heightmapBounds)
			if isLow {
				lowestPositions[lowPosition] = lowHeight
			}
		}
	}

	riskLevelSum := 0
	for _, height := range lowestPositions {
		riskLevelSum += height + 1
	}

	fmt.Println(riskLevelSum)
}

func findLowestPositionFromPoint(heightmap [][]int, height int, position vector, heightmapBounds bounds) (bool, vector, int) {
	isLowest := true
	isHighest := true
	lowerPositions := make(map[int]vector)

	if position.x > heightmapBounds.min.x {
		// check left
		left := heightmap[position.y][position.x-1]
		if height > left {
			isLowest = false
			lowerPositions[left] = vector{x: position.x - 1, y: position.y}
		} else if height < left {
			isHighest = false
		}
	}
	if position.x < heightmapBounds.max.x {
		// check right
		right := heightmap[position.y][position.x+1]
		if height > right {
			isLowest = false
			lowerPositions[right] = vector{x: position.x + 1, y: position.y}
		} else if height < right {
			isHighest = false
		}
	}
	if position.y > heightmapBounds.min.y {
		// check up
		up := heightmap[position.y-1][position.x]
		if height > up {
			isLowest = false
			lowerPositions[up] = vector{x: position.x, y: position.y - 1}
		} else if height < up {
			isHighest = false
		}
	}
	if position.y < heightmapBounds.max.y {
		// check down
		down := heightmap[position.y+1][position.x]
		if height > down {
			isLowest = false
			lowerPositions[down] = vector{x: position.x, y: position.y + 1}
		} else if height < down {
			isHighest = false
		}
	}

	if isLowest {
		if isHighest {
			// all surrounding heights are equal to this height
			return false, vector{}, 0
		}
		return true, vector{x: position.x, y: position.y}, height
	}
	keys := make([]int, 0, len(lowerPositions))
	for k := range lowerPositions {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	height = keys[0]
	position = lowerPositions[height]
	return findLowestPositionFromPoint(heightmap, height, position, heightmapBounds)
}
