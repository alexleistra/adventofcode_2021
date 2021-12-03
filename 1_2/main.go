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
	var prevSum int // previous sum
	var count int   // times the sum of measurements increases
	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		// add new
		nums = append(nums, num)

		if len(nums) == 3 {
			// we can get the sum
			currSum := nums[0] + nums[1] + nums[2]

			// fmt.Printf("%v + %v + %v = ", nums[0], nums[1], nums[2])
			// fmt.Println(currSum)

			// compare the sum with previous sum
			if prevSum > 0 {
				if currSum > prevSum {
					// increased
					count++
				}
			}

			prevSum = currSum

			// remove first
			_, nums = nums[0], nums[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}
