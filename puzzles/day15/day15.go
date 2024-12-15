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

  for _, movement := range movements {
    start = move(matrix, movement, start)
    fmt.Printf("After movement %v:\n", movement)
    fileops.PrintRuneMatrix(matrix)
  }



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

func move(matrix [][]rune, movement rune, point Point) Point {
	commandMap := map[rune]Point{
		'<': {X: -1, Y: 0}, // Left
		'>': {X: 1, Y: 0},  // Right
		'^': {X: 0, Y: -1}, // Up
		'v': {X: 0, Y: 1},  // Down
	}

	direction, ok := commandMap[movement]

	if !ok {
		fmt.Printf("Unknown movement: %v\n", movement)
		return point
	}

	currentPos := point

	nextPos := Point{X: currentPos.X + direction.X, Y: currentPos.Y + direction.Y}

	switch matrix[nextPos.Y][nextPos.X] {
	case '#': // wall
		return currentPos
	case 'O': // box
		if pushBoxes(matrix, nextPos, direction) {
			matrix[currentPos.Y][currentPos.X] = ' ' 
			matrix[nextPos.Y][nextPos.X] = '@'       
			return nextPos
		} else {
			return currentPos
		}
	default: 
		matrix[currentPos.Y][currentPos.X] = ' ' 
		matrix[nextPos.Y][nextPos.X] = '@'
		return nextPos
	}
}

func pushBoxes(matrix [][]rune, start Point, direction Point) bool {
	current := start
	for {
		next := Point{X: current.X + direction.X, Y: current.Y + direction.Y}

		switch matrix[next.Y][next.X] {
		case ' ': 
			matrix[next.Y][next.X] = 'O'
			matrix[current.Y][current.X] = ' '
			return true
		case 'O': 
			current = next
		default: 
			return false
		}
	}
}
