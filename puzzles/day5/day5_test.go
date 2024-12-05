package day5

import (
	"testing"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func TestDay5(t *testing.T) {
	inputFile := "../../input/day5_test.txt"
	body, err := fileops.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	rules, pages := splitAndMakeSlices(body)
	validUpdatesIdx, invalidUpdatesIdx := getIdxValidInvalid(rules, pages)

	validNumbers := createNmbSlices(pages, validUpdatesIdx)
	invalidNumbers := createNmbSlices(pages, invalidUpdatesIdx)

	t.Run("Test Day 5 Puzzle 1", func(t *testing.T) {

		expectedResult := 143

		result := findMiddleAndSum(validNumbers)
		if result != expectedResult {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result)
		}
	})
	t.Run("Test Day 5 Puzzle 2", func(t *testing.T) {

		expectedResult := 123
		fixedInvalidNumbers := fixOrder(rules, invalidNumbers)
		result := findMiddleAndSum(fixedInvalidNumbers)
		if result != expectedResult {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result)
		}
	})

}
