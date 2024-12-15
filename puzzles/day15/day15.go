package day15

import (
	"fmt"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

type Point struct {
	X, Y int
}

func Day15(input string) (int, int) {
	commandMap := map[string]Point{
		"<": {X: -1, Y: 0}, // Left
		">": {X: 1, Y: 0},  // Right
		"^": {X: 0, Y: -1}, // Up
		"v": {X: 0, Y: 1},  // Down
	}
	body, err := parseFile(input)
	if err != nil {
		fmt.Printf("error reading file: %vn", err)
		return 0, 0
	}

  matrix, movements := getMapAndMovement(body)
	fileops.PrintRuneMatrix(matrix)
	fmt.Printf("Movements:\n%v \n", string(movements))

	start := findStartingPoint(matrix)
	fmt.Printf("Start:\n%v \n", start)
	return 0, 0
}

func findStartingPoint(matrix [][]rune) Point {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '@' {
				return Point{i, j}
			}
		}
	}
	return Point{0, 0}
}

func parseFile(input string) ([]byte, error) {
	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Printf("error reading file: %vn", err)
		return nil, err
	}
	return body, nil
}

func getMapAndMovement(body []byte) ([][]rune, []rune) {

	content := string(body)
	contents := strings.Split(content, "\n\n")

	stringMatrix := contents[0]
	movementList := contents[1]

	matrix := fileops.MakeSlice([]byte(stringMatrix))
	movements := []rune(movementList)

  return matrix, movements
}
