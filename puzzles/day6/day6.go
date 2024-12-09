package day6

import (
	// "fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

type point struct {
	x int
	y int
}

var directions = []point{ //-1,0 is top left
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

var newStart = []point{
	{0, 1},  // start is 1 to the right
	{1, 0},  // start is 1 down
	{0, -1}, // start is 1 the left
	{-1, 0}, // start is 1 up
}

var startDirectionIndex = 0

func createRuneSlice(input string) (slice [][]rune) {
	body, err := fileops.ReadFile(input)
	if err != nil {
		panic(err)
	}

	slice = fileops.MakeSlice(body) //create 2d rune slice of map

	// lines := fileops.MakeStringSlice(body)
	// fileops.PrintMap(lines)

	return slice
}

func getStartPoint(slice [][]rune) (startPoint point) {

	for i := range slice {
		for j := range slice[i] {

			if slice[i][j] == '^' { //find starting point, we know it must be ^; TODO: check for other directions
				startPoint = point{x: i, y: j}
				slice[i][j] = '.'
			}
		}
	}

	return startPoint
}

// we use this for the second puzzle

func findWays(input string) (visitedPositions []point, length int) {

	slice := createRuneSlice(input)
	start := getStartPoint(slice)

	position := start
	directionsIndex := startDirectionIndex
	direction := directions[directionsIndex]

	xLength := len(slice[0])
	yLength := len(slice)

	visited := make(map[point]bool)
	visited[start] = true
	visitedPositions = append(visitedPositions, start) //
	for {
		nextPosition := point{
			x: position.x + direction.x,
			y: position.y + direction.y,
		}

		if nextPosition.x < 0 || nextPosition.x >= xLength || nextPosition.y < 0 || nextPosition.y >= yLength {
			break
		}

		if slice[nextPosition.x][nextPosition.y] == '#' {

			directionsIndex = (directionsIndex + 1) % len(directions) // get next direction of directionsSlice, wrap around if necessary
			direction = directions[directionsIndex]

			continue
		}

		position = nextPosition

		if !visited[position] {
			visited[position] = true
			visitedPositions = append(visitedPositions, position)
			// fmt.Println(visited)
		}
	}
	return visitedPositions, len(visitedPositions)
}

func testForLoop(input string) int {
	slice := createRuneSlice(input)
	visitedPositions, _ := findWays(input)
	start := getStartPoint(slice)

	loopCount := 0

	for _, p := range visitedPositions {
		if slice[p.x][p.y] == '#' {
			continue
		}

		slice[p.x][p.y] = '#'

		if detectsLoop(slice, start) {
			loopCount++
		}

		slice[p.x][p.y] = '.'
	}

	return loopCount
}

func detectsLoop(slice [][]rune, start point) bool {
	xLength := len(slice[0])
	yLength := len(slice)
	position := start
	directionsIndex := startDirectionIndex
	direction := directions[directionsIndex]

	var turningPos []point

	for {
		nextPosition := point{
			x: position.x + direction.x,
			y: position.y + direction.y,
		}

		if nextPosition.x < 0 || nextPosition.x >= yLength || nextPosition.y < 0 || nextPosition.y >= xLength {
			return false
		}

		if slice[nextPosition.x][nextPosition.y] == '#' {
			directionsIndex = (directionsIndex + 1) % len(directions)
			direction = directions[directionsIndex]
			turningPos = append(turningPos, position)
			continue
		}

		if compareLastWithSubset(turningPos) {
			return true
		}

		position = nextPosition
	}
}

func compareLastWithSubset(arr []point) bool {
	if len(arr) < 4 {
		return false
	}

	last := arr[len(arr)-1]    // Letztes Element
	subset := arr[:len(arr)-3] // Alle außer den letzten drei Elementen
	for _, v := range subset {
		if v == last {
			return true // Das letzte Element ist in diesem Subset vorhanden
		}
	}
	return false // Kein Duplikat gefunden
}

func Day6(input string) (score1, score2 int) {

	// puzzle 1
	_, score1 = findWays(input)

	// puzzle 2
	score2 = testForLoop(input)

	return score1, score2
}
