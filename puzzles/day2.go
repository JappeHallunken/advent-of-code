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
     totalScore := 0
	for _, l := range lines {
		nmbs := extractNumbers(l)
		// fmt.Println(numbers)

		score := SequenceType(nmbs)
    totalScore += score
	}
  fmt.Println("Day 2 puzzle 2: ", totalScore)
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
