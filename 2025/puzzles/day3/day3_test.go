package day3_test

import (
	"advent-of-code-2025/puzzles/day3"
	"testing"
)

func TestDay3(t *testing.T) {

	result1, result2 := day3.Day3(day3.TestInput)

	expected1 := 357
	if result1 != expected1 {
		t.Errorf("expected %d, got %d", expected1, result1)
	}

	expected2 := 3121910778619
	if result2 != expected2 {
				t.Errorf("expected %d, got %d", expected2, result2)
	}
	
}
