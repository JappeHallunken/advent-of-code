package day1

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

func sumDiff(firstElements, secondElements []int) (sum int) {
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

func similarityScore(firstElements, secondElements []int) (score int) {
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

func Day1(input string) {
	//prepare
	 body, err := fileops.ReadFile(input)
	firstElements, secondElements := fileops.MakeSlices(body)
	if err != nil {
		fmt.Printf("Day 1: Error reading file: %v\n", err)
	} else {

		// puzzle 1
		diff := sumDiff(firstElements, secondElements)
		fmt.Println("Day 1 Puzzle 1: ", diff)
		// puzzle 2
		score := similarityScore(firstElements, secondElements)
		fmt.Println("Day 1 puzzle 2: ", score)
	}
}
