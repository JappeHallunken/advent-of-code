package puzzles

import (
	"fmt"
	"github.com/JappeHallunken/advent-of-code/fileops"
	"strconv"
	"strings"
)

func SafetyCheck() {
	// fmt.Println("Day 2")
	body := fileops.ReadFile("input/day2.txt")

	lines := strings.Split(string(body), "\n")
	totalScore1 := 0
	totalScore2 := 0
	for _, l := range lines {
		nmbs := extractNumbers(l)
		// fmt.Println(numbers)
		score := SequenceType(nmbs)
    // fmt.Println("Score: ", score)
		totalScore1 += score
	}
	fmt.Println("Day 2 puzzle 1: ", totalScore1)

	for _, l := range lines {
		nmbs := extractNumbers(l)
		// fmt.Println(numbers)
		score2 := CountValidSequencesWithOneRemoved(nmbs)
		totalScore2 += score2
	}

	fmt.Println("Day 2 puzzle 2: ", totalScore2)
}

func extractNumbers(line string) (numbers []int) {
	elements := strings.Fields(line)
	for _, element := range elements {
		number, _ := strconv.Atoi(element)
		numbers = append(numbers, number)
	}
	return numbers
}

func SequenceType(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}

	increasing, decreasing := true, true
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

func CountValidSequencesWithOneRemoved(numbers []int) int {
    validCount := 0
    for i := range numbers {
        removedSlice := append(numbers[:i:i], numbers[i+1:]...)
        if SequenceType(removedSlice) > 0 {
            validCount++
            break
        }
    }
    return validCount
}
