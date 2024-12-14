package day14

import (
	"fmt"
	"testing"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func TestDay14(t *testing.T) {
	input := "../../input/day14_test.txt"
	xLength := 11
	yLength := 7

	t.Run("Day14", func(t *testing.T) {
		body, err := fileops.ReadFile(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))

		robots := getRobots(string(body))
		fmt.Println(robots)
		space := createSpace(xLength, yLength, robots)
		fileops.PrintRuneMatrix(space)
	})
}
