package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Cache that stores the number of numbers generated for each number and each iteration
var cache map[string]int

// Transform applies the rules to a number and returns the resulting numbers
func Transform(number int) []int {
	if number == 0 {
		return []int{1}
	} else if len(strconv.Itoa(number))%2 == 0 {
		s := strconv.Itoa(number)
		mid := len(s) / 2
		left, _ := strconv.Atoi(s[:mid])
		right, _ := strconv.Atoi(s[mid:])
		return []int{left, right}
	} else {
		return []int{number * 2024}
	}
}

// CountNumbersAfterIterations uses the cache to calculate the number of resulting numbers after n iterations
func CountNumbersAfterIterations(number, iterations int) int {
	// Cache key is the comination of number and remaining iterations
	key := fmt.Sprintf("%d:%d", number, iterations)
	// fmt.Printf("Iteration: %d, Number: %d\n", iterations, number)
	if val, found := cache[key]; found {
		// returns the value from the cache.
		return val
	}

  // default case: no iterations left, 1 number remains
	if iterations == 0 {
		return 1
	}

	// apply the rules and calculate recursively
	transformed := Transform(number)
	count := 0
	for _, newNumber := range transformed {
		count += CountNumbersAfterIterations(newNumber, iterations-1)
	}

	// save result to cache
	cache[key] = count
	return count
}

// iterate and find the total amount of numbers. 
func Iterate(numbers []int, iterations int) int {
	totalCount := 0

	for _, number := range numbers {
		totalCount += CountNumbersAfterIterations(number, iterations)
	}

	return totalCount
}
func Day11(input string, iterations int) int {
	// init cache
	cache = make(map[string]int)
	body, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	numbers := []int{}
	strNumbers := strings.Split(strings.TrimSpace(string(body)), " ")
	for _, strNumber := range strNumbers {
		number, err := strconv.Atoi(strNumber)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return 0
		}
		numbers = append(numbers, number)
	}

	result := Iterate(numbers, iterations)

	return result
}
