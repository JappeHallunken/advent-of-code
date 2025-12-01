package day2

import (
	"fmt"
	"github.com/JappeHallunken/advent-of-code/fileops"
	"strconv"
	"strings"
)

func extractNumbers(line string) (numbers []int) {
	elements := strings.Fields(line)
	for _, element := range elements {
		number, _ := strconv.Atoi(element)
		numbers = append(numbers, number)
	}
	return numbers
}

func sequenceType(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}

	increasing := true
	decreasing := true
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff < 1 || diff > 3 {
			increasing = false
		}
		if diff > -1 || diff < -3 {
			decreasing = false
		}
	}
	if increasing || decreasing {
		return 1
	}

	return 0
}

func countValidSequencesWithOneRemoved(numbers []int) int {
	validCount := 0
	for i := range numbers {
		removedSlice := append([]int(nil), numbers[:i]...)
		removedSlice = append(removedSlice, numbers[i+1:]...)
		if sequenceType(removedSlice) > 0 {
			validCount++
			break
		}
	}
	return validCount
}

func Day2(input string) (totalScore1, totalscore2 int) {
	// fmt.Println("Day 2")
	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Printf("Day 2: Error reading file: %v\n", err)
		return
	} else {

		lines := strings.Split(string(body), "\n")
		totalScore1 := 0
		totalScore2 := 0
		for _, l := range lines {
			nmbs := extractNumbers(l)
			// fmt.Println(numbers)
			score := sequenceType(nmbs)
			// fmt.Println("Score: ", score)
			totalScore1 += score
		}

		for _, l := range lines {
			nmbs := extractNumbers(l)
			// fmt.Println(numbers)
			score2 := countValidSequencesWithOneRemoved(nmbs)
			totalScore2 += score2
		}

		return totalScore1, totalScore2
	}
}
