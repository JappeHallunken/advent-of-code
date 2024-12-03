package tests

import (
	"testing"

	"github.com/JappeHallunken/advent-of-code/puzzles"
)

// Mock version of readFile for testing
func MockReadFile() []byte {
	// Return the specific body string for testing, converted to []byte
	return []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
}

func MockReadFile2() []byte {
	// Return the specific body string for testing, converted to []byte
	return []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
}

func TestCalculateSum(t *testing.T) {
	// Run the first test: Testing CalculateSum with mock input
	t.Run("TestCalculateSum", func(t *testing.T) {
		expected := 161

		// Using MockReadFile to simulate reading the file
		body := MockReadFile()

		// Call CalculateSum with the mock data
		actual, _ := puzzles.CalculateSum(body)

		// Check if the result matches the expected value
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}

func TestMakeString(t *testing.T) {
	// Run the second test: Testing MakeString and CalculateSum
	t.Run("TestMakeString", func(t *testing.T) {
		// The input string for testing MakeString
		body := MockReadFile2()
    // body := []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
    // fmt.Println(string(body))
		// Step 1: Transform the string using MakeString
		newString := puzzles.MakeString(body)

		// Step 2: Calculate sum based on the new transformed string
		actual, _ := puzzles.CalculateSum([]byte(newString))

		// Define the expected result after transformation and sum calculation
		expected := 48

		// Check if the result matches the expected value
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}
