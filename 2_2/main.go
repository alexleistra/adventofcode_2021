package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2021/day/2
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontal, vertical, aim int
	for scanner.Scan() {
		velocity := scanner.Text()
		params := getParams(`(?P<direction>[a-z]+)\s(?P<count>\d)`, velocity)

		count, err := strconv.Atoi(params["count"])
		if err != nil {
			log.Fatal(err)
		}

		switch params["direction"] {
		case "up":
			aim -= count
		case "down":
			aim += count
		case "forward":
			horizontal += count
			vertical += aim * count
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontal * vertical)
}

func getParams(regEx, value string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(value)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
