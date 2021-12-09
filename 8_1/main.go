package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2021/day/8
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputs := []string{}
	outputs := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " | ")
		inputs = append(inputs, parts[0])
		outputs = append(outputs, parts[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0

	for i, _ := range inputs {
		output := outputs[i]
		outputValues := strings.Split(output, " ")
		for _, value := range outputValues {
			// 1, 4, 7, 8
			if len(value) == 2 || len(value) == 4 || len(value) == 3 || len(value) == 7 {
				count++
			}

		}
	}
	fmt.Println(count)
}
