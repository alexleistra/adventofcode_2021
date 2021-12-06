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
	var fishInput [9]int

	for scanner.Scan() {
		input := scanner.Text()
		inputList := strings.Split(input, ",")
		for _, inputAge := range inputList {
			age, err := strconv.Atoi(inputAge)
			if err != nil {
				log.Fatal(err)
			}
			fishInput[age]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fish := fishInput[:]
	for i := 0; i < 256; i++ {
		fishesSpawning := fish[0]
		fish = fish[1:]
		if fishesSpawning > 0 {
			fish[6] += fishesSpawning
		}
		fish = append(fish, fishesSpawning)
	}

	count := 0
	for _, f := range fish {
		count += f
	}
	fmt.Println(count)
}
