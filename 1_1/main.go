package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prev int
	var count int
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if prev > 0 {
			if curr > prev {
				// increased
				count++
			}
		}

		// compare the next number with curr
		prev = curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
