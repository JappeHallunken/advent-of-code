package day1_test

import (
	"advent-of-code-2025/puzzles/day1"
	"path/filepath"
	"runtime"
	"testing"
)

func TestD1P2(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0) 
	testFile := filepath.Join(filepath.Dir(filename), "test.txt")

	result, err := day1.P2(testFile)
	if err != nil {
		t.Fatalf("D1P1() returned an error: %v", err)
	}

	expected := 6
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
