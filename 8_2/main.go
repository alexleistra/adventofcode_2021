package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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

	for i, lineInput := range inputs {
		output := outputs[i]
		outputValues := strings.Split(output, " ")

		// the start of the dumbest shit ever
		reFour := regexp.MustCompile(`\b\w{4}\b`)
		four := sortString(reFour.FindString(lineInput))
		reOne := regexp.MustCompile(`\b\w{2}\b`)
		one := sortString(reOne.FindString(lineInput))

		outputDisplay := ""
		for _, value := range outputValues {

			// sort value
			value = sortString(value)

			// the realization of the dumbest shit ever
			converted := ""
			if len(value) == 2 {
				converted = "1"
			} else if len(value) == 4 {
				converted = "4"
			} else if len(value) == 3 {
				converted = "7"
			} else if len(value) == 7 {
				converted = "8"
			} else if containsAll(value, one) {
				if containsAll(value, four) {
					converted = "9"
				} else if len(value) == 6 {
					converted = "0"
				} else {
					converted = "3"
				}
			} else {
				if countContains(value, four) == 2 {
					converted = "2"
				} else if len(value) == 5 {
					converted = "5"
				} else {
					converted = "6"
				}
			}

			outputDisplay += converted
		}

		outputNumber, err := strconv.Atoi(outputDisplay)
		if err != nil {
			log.Fatal(err)
		}
		count += outputNumber
	}
	fmt.Println(count)
}

func sortString(value string) string {
	s := strings.Split(value, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func containsAll(str string, chstr string) bool {
	for _, c := range chstr {
		if !strings.Contains(str, string(c)) {
			return false
		}
	}
	return true
}

func countContains(str string, chstr string) int {
	count := 0
	for _, c := range chstr {
		if strings.Contains(str, string(c)) {
			count++
		}
	}
	return count
}
