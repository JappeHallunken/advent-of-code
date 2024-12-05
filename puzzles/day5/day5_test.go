package day5

import (
	"testing"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func TestDay5(t *testing.T) {
	t.Run("Test Day 5 Puzzle 1", func(t *testing.T) {
		inputFile := "../../input/day5_test.txt"
		expectedResult := 143 // Erwartetes Ergebnis

    body, err := fileops.ReadFile(inputFile)
    if err != nil {
        t.Fatal(err)
    }
    data, data2 := splitAndMakeSlices(body)

    orderArray := checkForRightOrder(data, data2)

    result := findMiddleAndSum(orderArray, data2)

    if  result != expectedResult {
        t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result)
    }

	})
}
