package main

import "testing"


func TestSimilarityScore(t *testing.T) {
	tests := []struct {
		name           string
		firstElements  []int
		secondElements []int
		expectedScore    int
	}{
		{
			name:           "Test 1",
			firstElements:  []int{3, 4, 2, 1, 3, 3},
			secondElements: []int{4, 3, 5, 3, 9, 3},
			expectedScore:      31,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualScore := similarityScore(test.firstElements, test.secondElements)
			if actualScore != test.expectedScore {
				t.Errorf("Expected score %d, got %d", test.expectedScore, actualScore)
			}
		})
	}

}

