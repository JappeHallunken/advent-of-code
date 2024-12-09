package day8

import (
	"fmt"
	"testing"
)

func TestDay8(t *testing.T) {

	t.Run("Day 8", func(t *testing.T) {
		expected := 14

		d8p1, _ := Day8("../../input/day8_test.txt")
    // fmt.Printf("\nDay 8 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d8p1, d8p2)

		if d8p1 != expected {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expected, d8p1)
		} else {
			fmt.Println("Test 1 Success!")
		}

	})
}
