package day6_test

import (
	"advent-of-code-2025/puzzles/day6"
	"testing"
)

func TestDay6P1(t *testing.T) {

	result := day6.P1(day6.TestInput)

	expected := 4277556
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestDay6P2(t *testing.T) {

	result := day6.P2(day6.TestInput)

	expected := 3263827
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
