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
	cycles := 100

	t.Run("Day14", func(t *testing.T) {
    expected := 12
		body, err := fileops.ReadFile(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))

		robots := getRobots(string(body))
		fmt.Println(robots)
		// robots := []Robot{
		// 	{
		// 		Position: Point{X: 2, Y: 4},
		// 		Velocity: Point{X: 2, Y: -3},
		// 	},
		// }
		space := createSpace(xLength, yLength, robots)
		fileops.PrintRuneMatrix(space)

		space = moveRobots(space, robots, cycles)

		fmt.Printf("\nAfter %v cycles:\n", cycles)
		fileops.PrintRuneMatrix(space)
    score := calculateSafetyFactor(space)

    if score != expected {
      t.Errorf("expected %v, got %v", expected, score)
    }
	})
}
