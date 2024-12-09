package day8

import (
	"fmt"
	"testing"
)

func TestDay8(t *testing.T) {
	d8p1, d8p2 := Day8("../../input/day8_test.txt")

	t.Run("Day 8_1", func(t *testing.T) {
		expected := 14

		// fmt.Printf("\nDay 8 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d8p1, d8p2)

		if d8p1 != expected {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expected, d8p1)
		} else {
			fmt.Println("Test 1 Success!")
		}

	})
	t.Run("Day 8_2", func(t *testing.T) {
		expected := 34

		// fmt.Printf("\nDay 8 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d8p1, d8p2)

		if d8p2 != expected {
			t.Errorf("Erwartetes Ergebnis: %d, erhalten: %d", expected, d8p2)
		} else {
			fmt.Println("Test 2 Success!")
		}

	})

}
