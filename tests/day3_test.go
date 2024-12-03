package tests

import (
	"github.com/JappeHallunken/advent-of-code/puzzles"
	"testing"
)

// Mock version of readFile for testing
func MockReadFile() []byte {
	// Return the specific body string for testing, converted to []byte
	return []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
}

func TestDay3(t *testing.T) {

  expected := 161

  body := MockReadFile()
  actual := puzzles.CalculateSum(body)

  if actual != expected {
    t.Errorf("Expected %d, got %d", expected, actual)
  }

}
