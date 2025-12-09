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

func TestDay8P2(t *testing.T) {

	result := P2(TestInput)

	expected := 25272
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
