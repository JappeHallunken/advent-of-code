package day14

import (
	"fmt"
	"testing"

	"github.com/JappeHallunken/advent-of-code/2024/fileops"
)

func TestDay14(t *testing.T) {
	input := "../../input/day14_test.txt"
	xLength := 11
	yLength := 7
	cycles := 100
	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
  score := 0
	t.Run("Day14", func(t *testing.T) {
		expected := 12
		robots := getRobots(string(body))
		space := createSpace(xLength, yLength, robots)
		for i := 0; i < cycles; i++ {
			score = getSafetyFactor(space)
		}

		if score != expected {
			t.Errorf("expected %v, got %v", expected, score)
		}
	})
}
