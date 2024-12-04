package day1

import (
	"testing"
)

// Helper function to avoid code repetition
func runTest(t *testing.T, testCases []struct {
	name           string
	expectedScore  int
}, similarityFunc func(firstElements, secondElements []int) int, firstElements, secondElements []int) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actualScore := similarityFunc(firstElements, secondElements)
			if actualScore != test.expectedScore {
				t.Errorf("Expected score %d, got %d", test.expectedScore, actualScore)
			}
		})
	}
}

func TestPuzzles(t *testing.T) {
	// Define shared arrays
	firstElements := []int{3, 4, 2, 1, 3, 3}
	secondElements := []int{4, 3, 5, 3, 9, 3}

	// Test cases for SimilarityScore
	similarityTests := []struct {
		name          string
		expectedScore int
	}{
		{
			name:          "Day 1 part 2",
			expectedScore: 31,
		},
	}

	// Run SimilarityScore tests
	t.Run("SimilarityScore", func(t *testing.T) {
		runTest(t, similarityTests, similarityScore, firstElements, secondElements)
	})

	// Test cases for SumDiff
	sumDiffTests := []struct {
		name        string
		expectedScore int
	}{
		{
			name:        "Day1 part 1",
			expectedScore: 11,
		},
	}

	// Run SumDiff tests
	t.Run("SumDiff", func(t *testing.T) {
		runTest(t, sumDiffTests, sumDiff, firstElements, secondElements)
	})
}
