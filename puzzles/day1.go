package puzzles

import "github.com/JappeHallunken/advent-of-code/fileops"

// puzzle 1

func SumDiff(firstElements, secondElements []int) (sum int) {
	sum = 0

  //sort the slices
	sortedFirstElements, sortedSecondElements := fileops.SortSlices(firstElements, secondElements)
  
  // calculate the sum
  for i := range sortedFirstElements {
		diff := sortedFirstElements[i] - sortedSecondElements[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	// fmt.Println(sum)
	return sum
}

// puzzle 2

func SimilarityScore(firstElements, secondElements []int) (score int) {
	score = 0
	for i := range firstElements {
		for j := range secondElements {
			multiplier := 0

			if firstElements[i] == secondElements[j] {
				multiplier++
				// fmt.Println(i, j, multiplier)
			}
			score = score + (firstElements[i] * multiplier)
			// fmt.Println(score)
		}
	}
	return score
}
