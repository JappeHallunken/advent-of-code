package tests

import ( 
  "testing"
  "github.com/JappeHallunken/advent-of-code/puzzles"
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

    expectedScore := 2
    totalScore := 0

    for _, sequence := range sequences {
        score := puzzles.SequenceType(sequence)
        totalScore += score
    }

    if totalScore != expectedScore {
        t.Errorf("Failed test case for custom sequences: expected total score %d, got %d", expectedScore, totalScore)
    }
}
