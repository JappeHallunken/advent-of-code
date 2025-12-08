package day7_test

import (
	"advent-of-code-2025/puzzles/day7"
	"testing"
)

func TestDay7P1(t *testing.T) {

	result := day7.P1(day7.TestInput)

	expected := 21
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func TestDay6P2(t *testing.T) {
//
// 	result := day6.P2(day6.TestInput)
//
// 	expected := 3263827
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }
