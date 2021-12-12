package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/alexleistra/adventofcode_2021/day12/entities"
)

// https://adventofcode.com/2021/day/12
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allCaves := make(map[string]*entities.Cave)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputCaveNames := strings.Split(line, "-")
		caveA, okA := allCaves[inputCaveNames[0]]
		caveB, okB := allCaves[inputCaveNames[1]]

		if !okA {
			caveName := inputCaveNames[0]
			caveType := getCaveType(caveName)
			cave := entities.NewCave(caveName, caveType)
			allCaves[caveName] = &cave
			caveA = &cave
		}
		if !okB {
			caveName := inputCaveNames[1]
			caveType := getCaveType(caveName)
			cave := entities.NewCave(caveName, caveType)
			allCaves[caveName] = &cave
			caveB = &cave
		}
		caveA.AddLinkedCave(caveB)
		caveB.AddLinkedCave(caveA)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	start, ok := allCaves["start"]
	if ok {
		delete(allCaves, "start")
	}

	partOne(start, allCaves)
	partTwo(start, allCaves)
}

func getCaveType(caveName string) entities.CaveType {
	if caveName == "start" {
		return entities.Start
	} else if caveName == "end" {
		return entities.End
	} else if unicode.IsUpper(rune(caveName[0])) && unicode.IsLetter(rune(caveName[0])) {
		return entities.Large
	}
	return entities.Small
}

func findPaths(cave *entities.Cave, currentPath []string, visitedSmallCaves []string, caveAllowedMultipleVisits string, maxSmallVisits int) int {
	paths := 0

	currentPath = append(currentPath, cave.GetName())
	if cave.GetType() == entities.Small {
		if maxSmallVisits > 1 {
			// check if we have already visited this small cave
			for _, visitedCave := range visitedSmallCaves {
				if visitedCave == cave.GetName() {
					// we visited this cave already
					if caveAllowedMultipleVisits == "" {
						// this cave can have multipe visits
						caveAllowedMultipleVisits = cave.GetName()
					}
					break
				}
			}
		}

		visitedSmallCaves = append(visitedSmallCaves, cave.GetName())
	}

	if cave.GetType() == entities.End {
		// fmt.Println(currentPath)
		paths++
	} else {
		for _, nextCave := range cave.GetLinkedCaves() {
			caveBlocked := false
			if nextCave.GetType() == entities.Small {
				thisCaveVisited := 0
				for _, visitedSmallCave := range visitedSmallCaves {
					if visitedSmallCave == nextCave.GetName() {
						thisCaveVisited++
					}
				}
				if maxSmallVisits > 1 {
					if caveAllowedMultipleVisits != "" {
						if nextCave.GetName() == caveAllowedMultipleVisits {
							if thisCaveVisited == maxSmallVisits {
								caveBlocked = true
							}
						} else if thisCaveVisited > 0 {
							caveBlocked = true
						}
					}
				} else if thisCaveVisited > 0 {
					caveBlocked = true
				}
			}

			if !caveBlocked {
				paths += findPaths(nextCave, currentPath, visitedSmallCaves, caveAllowedMultipleVisits, maxSmallVisits)
			}
		}
	}
	return paths
}

func partOne(start *entities.Cave, allCaves map[string]*entities.Cave) {
	paths := findPaths(start, []string{}, []string{}, "", 1)

	fmt.Println("paths")
	fmt.Println(paths)
}

func partTwo(start *entities.Cave, allCaves map[string]*entities.Cave) {
	paths := findPaths(start, []string{}, []string{}, "", 2)

	fmt.Println("paths")
	fmt.Println(paths)
}
