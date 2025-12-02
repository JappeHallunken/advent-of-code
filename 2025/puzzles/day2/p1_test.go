package day2_test

import (
	"advent-of-code-2025/puzzles/day2"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDay2(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	testFile := filepath.Join(filepath.Dir(filename), "test.txt")

	result, err := day2.P1(testFile)
	if err != nil {
		t.Fatalf("D2P1 returned an error: %v", err)
	}
	expected := 1227775554
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	result2, err := day2.P2(testFile)
	if err != nil {
		t.Fatalf("D2P2 returned an error: %v", err)
	}
	expected2 := 4174379265
	if result2 != expected2 {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
