package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/6
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fish := []int{}
	for scanner.Scan() {
		input := scanner.Text()
		inputList := strings.Split(input, ",")
		for _, inputAge := range inputList {
			age, err := strconv.Atoi(inputAge)
			if err != nil {
				log.Fatal(err)
			}
			fish = append(fish, age)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 80 days
	for i := 0; i < 80; i++ {
		todaysSpawn := []int{}
		for k, f := range fish {
			if f == 0 {
				f = 6
				todaysSpawn = append(todaysSpawn, 8) // new fish takes 2 days longer to spawn on first cycle
			} else {
				f--
			}
			fish[k] = f
		}
		fish = append(fish, todaysSpawn...)
		fmt.Println(len(fish))
	}
}
