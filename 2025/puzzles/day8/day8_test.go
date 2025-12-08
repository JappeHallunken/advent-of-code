package day8
import (
	"testing"
)

func TestDay8P1(t *testing.T) {
	old := TopN                   // save the old variable value
	TopN = 10                      // modify temporarily
	defer func() { TopN = old }() // restore
	result := P1(TestInput)

	expected := 40
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
