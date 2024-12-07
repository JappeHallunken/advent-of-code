package day6

import (
	// "fmt"
	"slices"

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

var visitedPositions []point

func findWays(input string) int {

	slice := createRuneSlice(input)
	start := getStartPoint(slice)

	position := start
	directionsIndex := startDirectionIndex
	startDirection := directions[directionsIndex]
	direction := startDirection

	xLength := len(slice[0])
	yLength := len(slice)
	visitedPositions = append(visitedPositions, start)

	// fmt.Println("dimensions: ", xLength, yLength)
	// fmt.Println("start: ", start)
	// fmt.Println("direction: ", startDirection)

	outsideArea := false

	for !outsideArea {
		// fmt.Println("steps: ", steps)
		// fmt.Println("current position: ", position)

		nextPosition := point{position.x + direction.x, position.y + direction.y}

		if nextPosition.x < 0 || nextPosition.x >= xLength || nextPosition.y < 0 || nextPosition.y >= yLength {
			// fmt.Println("leaving area")
			outsideArea = true

			break
		}

		if slice[nextPosition.x][nextPosition.y] == '#' {

			directionsIndex = (directionsIndex + 1) % len(directions) // get next direction of directionsSlice, wrap around if necessary
			direction = directions[directionsIndex]

			//    fmt.Println("found obstacle")
			// fmt.Println("new direction: ", direction)
			//
			continue
		}
		if slice[nextPosition.x][nextPosition.y] == '.' {

			position = nextPosition
			if !slices.Contains(visitedPositions, position) {
				visitedPositions = append(visitedPositions, position)
			}
		}
		// fmt.Println("next position: ", position)
	}
	// fmt.Println(visitedPositions)
	return len(visitedPositions)
}

func testForLoop(input string) (loopCount int) {
	slice := createRuneSlice(input)

	xLength := len(slice[0])
	yLength := len(slice)

	start := getStartPoint(slice)
	// fmt.Println("start: ", start)
	loopcount := 0

	for _, p := range visitedPositions {
		
		// set every point in the [][]rune to '#' and see if we get a loop
		// record the turning positions, if they repeat we are in a loop
		// turningPosIdx := 0
		position := start
		directionsIndex := startDirectionIndex
		direction := directions[directionsIndex]

		var turningPos []point
		if slice[p.x][p.y] == '#' {
			continue
		}

		slice[p.x][p.y] = '#'

		outsideArea, foundLoop := false, false
		for !outsideArea || !foundLoop {

			nextPosition := point{position.x + direction.x, position.y + direction.y}

			if nextPosition.x < 0 || nextPosition.x >= xLength || nextPosition.y < 0 || nextPosition.y >= yLength {
				// fmt.Println("###### leaving area ######", nextPosition)
				outsideArea = true
				break
			}

			if slice[nextPosition.x][nextPosition.y] == '#' {

				directionsIndex = (directionsIndex + 1) % len(directions) // get next direction of directionsSlice, wrap around if necessary
				direction = directions[directionsIndex]
				turningPos = append(turningPos, position)

				// fmt.Println("found obstacle", position)

				continue
			}
			if compareLastWithSubset(turningPos) {
				loopcount++
				foundLoop = true
				// fmt.Println("loopcount: ", loopcount, "  ", i, j)
				// fmt.Println("turning points: ", turningPos)
				// fmt.Println("found loop at: ", turningPos[len(turningPos)-1])
				break
			}

			if slice[nextPosition.x][nextPosition.y] != '#' {
				position = nextPosition
			}
			// fmt.Println(turningPos)
		}
		slice[p.x][p.y] = '.'
	}

	return loopcount
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
	score1 = findWays(input)

	// puzzle 2
	score2 = testForLoop(input)

	return score1, score2
}
