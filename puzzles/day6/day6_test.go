package day6

import (
	"fmt"
	"testing"

)

func TestDay5(t *testing.T) {
	inputFile := "../../input/day6_test.txt"
	
	t.Run("Test Day 6 Puzzle 1", func(t *testing.T) {
  expectedResult := 41
  result :=	findWays(inputFile)
		if result != expectedResult {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result)

		}
	})
	t.Run("Test Day 5 Puzzle 2", func(t *testing.T) {
    fmt.Println("TBD")
		//
		// expectedResult := 123
		// fixedInvalidNumbers := fixOrder(rules, invalidNumbers)
		// result := findMiddleAndSum(fixedInvalidNumbers)
		// if result != expectedResult {
		// 	t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expectedResult, result)
		// }
	})

}
