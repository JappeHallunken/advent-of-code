package day4_test

import (
	"advent-of-code-2025/puzzles/day4"
	"testing"
)

func TestDay3(t *testing.T) {

	result1, result2 := day4.P1(day4.TestInput)

	expected1 := 13
	if result1 != expected1 {
		t.Errorf("expected %d, got %d", expected1, result1)
	}

	expected2 := 43
	if result2 != expected2 {
				t.Errorf("expected %d, got %d", expected2, result2)
	}

}
