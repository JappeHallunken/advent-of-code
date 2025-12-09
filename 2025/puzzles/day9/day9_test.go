package day9

import (
	"testing"
)

func TestDay8P1(t *testing.T) {
	result := P1(TestInput)

	expected := 50
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// func TestDay8P2(t *testing.T) {
//
// 	result := P2(TestInput)
//
// 	expected := 25272
// 	if result != expected {
// 		t.Errorf("expected %d, got %d", expected, result)
// 	}
// }
