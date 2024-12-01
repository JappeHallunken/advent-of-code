package puzzles

// puzzle 1

func SumDiff(firstElements, secondElements []int) (sum int) {
	sum = 0
	for i := range firstElements {
		diff := firstElements[i] - secondElements[i]
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
