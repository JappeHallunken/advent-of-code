package day13


import (
	"testing"
)

func TestDay12(t *testing.T) {

	// t.Run("splitInHalf", func(t *testing.T) {
	//
	// 	left, right := splitInHalf(120034)
	// 	fmt.Println(left, right)
	// })
	expected1 := 480
	// expected2 := 1206

	input := "../../input/day13_test.txt"
  input2 := "../../input/day13_2.txt"

	result1, _ := Day13(input, input2)
	t.Run("Day13 p1", func(t *testing.T) {

		if result1 != expected1 {
			t.Errorf("Day13() = %v, want %v", result1, expected1)
		}
	})
	// t.Run("Day12 p2", func(t *testing.T) {
	//
	// 	if result2 != expected2 {
	// 		t.Errorf("Day12() = %v, want %v", result2, expected2)
	// 	}
	// })

}
