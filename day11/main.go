package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alexleistra/adventofcode_2021/day11/entities"
)

// https://adventofcode.com/2021/day/11
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allSquid := [][]*entities.Squid{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineChars := strings.Split(line, "")
		squidRow := []*entities.Squid{}
		for _, c := range lineChars {
			energy, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			squid := entities.NewSquid(energy)
			squidRow = append(squidRow, &squid)
		}
		allSquid = append(allSquid, squidRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for y, squidRow := range allSquid {
		for x, squid := range squidRow {
			addNeighbours(allSquid, squid, x, y)
		}
	}

	//partOne(allSquid)
	partTwo(allSquid)
}

func partOne(allSquid [][]*entities.Squid) {
	steps := 100
	flashes := 0
	for i := 0; i < steps; i++ {
		for _, squidRow := range allSquid {
			for _, squid := range squidRow {
				squid.IncreaseEnergy()
			}
		}

		for _, squidRow := range allSquid {
			for _, squid := range squidRow {
				flashes += squid.FlashIfCanFlash()
			}
		}
	}

	fmt.Println("flashes")
	fmt.Println(flashes)
}

func partTwo(allSquid [][]*entities.Squid) {
	step := 0
	for true {
		step++
		flashes := 0

		for _, squidRow := range allSquid {
			for _, squid := range squidRow {
				squid.IncreaseEnergy()
			}
		}

		for _, squidRow := range allSquid {
			for _, squid := range squidRow {
				flashes += squid.FlashIfCanFlash()
			}
		}

		if flashes == len(allSquid)*len(allSquid[0]) {
			break
		}
	}

	fmt.Println("step")
	fmt.Println(step)
}

func addNeighbours(allSquid [][]*entities.Squid, squid *entities.Squid, x, y int) {
	left, right := false, false
	if x > 0 {
		// add left
		left = true
		squid.AddNeighbour(allSquid[y][x-1])
	}
	if x < len(allSquid[0])-1 {
		// add right
		right = true
		squid.AddNeighbour(allSquid[y][x+1])
	}
	if y > 0 {
		// add up
		squid.AddNeighbour(allSquid[y-1][x])
		if left {
			squid.AddNeighbour(allSquid[y-1][x-1])
		}
		if right {
			squid.AddNeighbour(allSquid[y-1][x+1])
		}
	}
	if y < len(allSquid)-1 {
		// add down
		squid.AddNeighbour(allSquid[y+1][x])
		if left {
			squid.AddNeighbour(allSquid[y+1][x-1])
		}
		if right {
			squid.AddNeighbour(allSquid[y+1][x+1])
		}
	}
}
