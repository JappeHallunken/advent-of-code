package puzzles

import (
	"fmt"
	"github.com/JappeHallunken/advent-of-code/fileops"
	"slices"
)

///////////////////////////////////////// puzzle 1

func sortSlices(firstElements, secondElements []int) (sortedFirstElements, sortedSecondElements []int) {
	slices.Sort(firstElements)
	slices.Sort(secondElements)
	return firstElements, secondElements
}

func SumDiff(firstElements, secondElements []int) (sum int) {
	sum = 0

	//sort the slices
	sortedFirstElements, sortedSecondElements := sortSlices(firstElements, secondElements)

	// calculate the sum
	for i := range sortedFirstElements {
		diff := sortedFirstElements[i] - sortedSecondElements[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	return sum
}

///////////////////////////////////////// puzzle 2

func SimilarityScore(firstElements, secondElements []int) (score int) {
	score = 0
	for i := range firstElements {
		for j := range secondElements {
			multiplier := 0

			if firstElements[i] == secondElements[j] {
				multiplier++
			}
			score = score + (firstElements[i] * multiplier)
		}
	}
	return score
}

func Day1() {
	// puzzle 1
	body, _ := fileops.ReadFile("input/day1.txt")
	firstElements, secondElements := fileops.MakeSlices(body)

	diff := SumDiff(firstElements, secondElements)
	fmt.Println("Day 1 Puzzle 1: ", diff)

	// puzzle 2
	score := SimilarityScore(firstElements, secondElements)
	fmt.Println("Day 1 puzzle 2: ", score)
}
