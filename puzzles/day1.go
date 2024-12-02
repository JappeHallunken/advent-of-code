package puzzles

import (
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
