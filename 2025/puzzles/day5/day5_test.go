package day5_test

import (
	"advent-of-code-2025/puzzles/day5"
	"testing"
)

func TestDay5(t *testing.T) {

	result1, result2 := day5.P1(day5.TestInput)

	expected1 := 3
	if result1 != expected1 {
		t.Errorf("expected %d, got %d", expected1, result1)
	}

	expected2 := 14
	if result2 != expected2 {
				t.Errorf("expected %d, got %d", expected2, result2)
	}

}
