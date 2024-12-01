package main

import (
	"testing"
	"github.com/JappeHallunken/advent-of-code/puzzles"
)

func TestPuzzles(t *testing.T) {
	// Group for SimilarityScore tests
	t.Run("SimilarityScore", func(t *testing.T) {
		tests := []struct {
			name           string
			firstElements  []int
			secondElements []int
			expectedScore  int
		}{
			{
				name:           "Test 0",
				firstElements:  []int{3, 4, 2, 1, 3, 3},
				secondElements: []int{4, 3, 5, 3, 9, 3},
				expectedScore:  31,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualScore := puzzles.SimilarityScore(test.firstElements, test.secondElements)
				if actualScore != test.expectedScore {
					t.Errorf("Expected score %d, got %d", test.expectedScore, actualScore)
				}
			})
		}
	})

	// Group for SumDiff tests
	t.Run("SumDiff", func(t *testing.T) {
		tests := []struct {
			name           string
			firstElements  []int
			secondElements []int
			expectedSum    int
		}{
			{
				name:           "Test 0",
				firstElements:  []int{3, 4, 2, 1, 3, 3},
				secondElements: []int{4, 3, 5, 3, 9, 3},
				expectedSum:    11,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actualSum := puzzles.SumDiff(test.firstElements, test.secondElements)
				if actualSum != test.expectedSum {
					t.Errorf("Expected sum %d, got %d", test.expectedSum, actualSum)
				}
			})
		}
	})
}
