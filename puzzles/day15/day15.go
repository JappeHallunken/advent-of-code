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

	for i := 0; i < len(movements); i++ {
		movement := movements[i]
		start = move(matrix, movement, start)
		fmt.Printf("After movement %v:\n", i)
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

	// Bestimme die nächste Position des Spielers
	currentPos := point
	nextPos := Point{X: currentPos.X + direction.X, Y: currentPos.Y + direction.Y}

	switch matrix[nextPos.Y][nextPos.X] {
	case '#': // Wand - der Spieler bleibt an der aktuellen Position
		return currentPos
	case 'O': // Box
		// Zähle, wie viele Kartons hinter der Box stehen
		boxCount := countBoxes(matrix, currentPos, direction)

		// Berechne die Position, wohin der erste Karton verschoben werden muss
		newBoxPos := Point{X: nextPos.X + boxCount*direction.X, Y: nextPos.Y + boxCount*direction.Y}

		// Wenn das Ziel frei ist, verschiebe nur den ersten Karton
		if matrix[newBoxPos.Y][newBoxPos.X] == '.' {
			// Verschiebe den ersten Karton
			matrix[currentPos.Y][currentPos.X] = '.'
			matrix[nextPos.Y][nextPos.X] = '@' // Der Spieler geht auf den ersten Karton
      matrix[newBoxPos.Y][newBoxPos.X] = 'O'

			// // Füge einen neuen Karton hinter der Reihe hinzu
			// behindLastBoxPos := Point{X: nextPos.X + direction.X*(boxCount+1), Y: nextPos.Y + direction.Y*(boxCount+1)}
			// matrix[behindLastBoxPos.Y][behindLastBoxPos.X] = 'O' // Setze den neuen Karton

			// Gebe die neue Position des Spielers zurück
			return nextPos
		}
		// Wenn das Ziel nicht frei ist, bleibt der Spieler an der aktuellen Position
		return currentPos
	default:
		// Wenn das nächste Feld leer ist, bewege den Spieler einfach
		matrix[currentPos.Y][currentPos.X] = '.'
		matrix[nextPos.Y][nextPos.X] = '@'
		return nextPos
	}
}
func countBoxes(matrix [][]rune, start Point, direction Point) int {
	next := Point{X: start.X + direction.X, Y: start.Y + direction.Y}



	// Wenn das nächste Feld eine Wand ist, stoppe die Zählung
	if matrix[next.Y][next.X] == '#' {
		return 0
	}

	// Wenn das nächste Feld leer ist, stoppe die Zählung
	if matrix[next.Y][next.X] == '.' {
		return 0
	}

	// Wenn wir auf einen Karton stoßen, zähle ihn und gehe weiter
	if matrix[next.Y][next.X] == 'O' {
		return 1 + countBoxes(matrix, next, direction)
	}

	// Wenn wir auf etwas anderes stoßen, stoppe die Zählung
	return 0
}
