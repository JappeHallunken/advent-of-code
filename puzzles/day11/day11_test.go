package day11

import (
	"testing"
)

func TestDay11(t *testing.T) {

	// t.Run("splitInHalf", func(t *testing.T) {
	//
	// 	left, right := splitInHalf(120034)
	// 	fmt.Println(left, right)
	// })

		input := "../../input/day11_test.txt"
	t.Run("Day11 p1, 6 cycles", func(t *testing.T) {
		cycles := 6
		expected := 22

    result := Day11(input, cycles)

		if result != expected {
			t.Errorf("Day11() = %v, want %v", result, expected)
		}
	})
t.Run("Day11 p1, 25 cycles", func(t *testing.T) {
		cycles := 25
		expected := 55312
		result := Day11(input, cycles)
			// fmt.Printf("After %v Cycle: %v\n", j, input)

		if result != expected {
			t.Errorf("Day11() = %v, want %v", result, expected)
		}
	})

}
