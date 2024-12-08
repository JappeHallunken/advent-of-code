package day7

import (
	"fmt"
	"testing"
)

func TestDay7(t *testing.T) {
	inputFile := "../../input/day7_test.txt"
	result1, result2, result3 := Day7(inputFile)

	t.Run("Test Day 7 Puzzle 1", func(t *testing.T) {
    var expectedResult uint64 = 3749

		if result1 != expectedResult {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result1)
		} else {
			fmt.Println("Success!")
		}
	})

	t.Run("Test Day 7 Puzzle 2", func(t *testing.T) {
		var expectedResult uint64 = 11387
		if result3 != expectedResult {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result3)
		} else {
			fmt.Println("Success!")
		}
	})
	fmt.Println("result 2: ", result2)

}
