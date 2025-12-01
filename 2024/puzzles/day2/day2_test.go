package day2

import (
	"testing"
)

func TestCustomSequences(t *testing.T) {
	sequences := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	t.Run("Day 2 puzzle 1", func(t *testing.T) {
		expectedScore1 := 2
		totalScore1 := 0

		for _, sequence := range sequences {
			score :=sequenceType(sequence)
			totalScore1 += score
		}

		if totalScore1 != expectedScore1 {
			t.Errorf("Failed test case for SequenceType: expected total score %d, got %d", expectedScore1, totalScore1)
		}
	})

	t.Run("Day 2 puzzle 2", func(t *testing.T) {
		expectedScore2 := 4
		totalScore2 := 0

		for _, sequence := range sequences {
			score := countValidSequencesWithOneRemoved(sequence)
			totalScore2 += score
		}

		if totalScore2 != expectedScore2 {
			t.Errorf("Failed test case for CountValidSequencesWithOneRemoved: expected total score %d, got %d", expectedScore2, totalScore2)
		}
	})
}
