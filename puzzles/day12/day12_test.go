package day12

import (
	"testing"
)

func TestDay12(t *testing.T) {

	// t.Run("splitInHalf", func(t *testing.T) {
	//
	// 	left, right := splitInHalf(120034)
	// 	fmt.Println(left, right)
	// })
	expected1 := 1930
	expected2 := 1206

	input := "../../input/day12_test.txt"

	result1, result2 := Day12(input)
	t.Run("Day12 p1", func(t *testing.T) {

		if result1 != expected1 {
			t.Errorf("Day12() = %v, want %v", result1, expected1)
		}
	})
	t.Run("Day12 p2", func(t *testing.T) {

		if result2 != expected2 {
			t.Errorf("Day12() = %v, want %v", result2, expected2)
		}
	})

}
