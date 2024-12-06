package day6

import (
	// "fmt"
	"fmt"
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

func findWays(input string) int {

	slice := createRuneSlice(input)
	start := getStartPoint(slice)

	position := start
	directionsIndex := startDirectionIndex
	startDirection := directions[directionsIndex]
	direction := startDirection

	xLength := len(slice[0])
	yLength := len(slice)

	var visitedPositions []point
	steps := 0 //starting position is a step

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
				steps++
				visitedPositions = append(visitedPositions, position)
			}
		}
		// fmt.Println("next position: ", position)
	}
	return steps
}

func testForLoop(input string) (loopCount int) {
	slice := createRuneSlice(input)

	xLength := len(slice[0])
	yLength := len(slice)

	start := getStartPoint(slice)
	// fmt.Println("start: ", start)
	loopcount := 0
	for i := range slice {
		for j := range slice[i] { // set every point in the [][]rune to '#' and see if we get a loop
			var turningPos []point // record the turning positions, if they repeat we are in a loop
			turningPosIdx := 0
			position := start
			directionsIndex := startDirectionIndex
			direction := directions[directionsIndex]

			// fmt.Println("i, j: ", i, j)

			if slice[i][j] == '#' {
				continue
			}

			slice[i][j] = '#'

			outsideArea, foundLoop := false, false
			for !outsideArea || !foundLoop {

				nextPosition := point{position.x + direction.x, position.y + direction.y}

				if nextPosition.x < 0 || nextPosition.x >= xLength || nextPosition.y < 0 || nextPosition.y >= yLength {
					// fmt.Println("leaving area")
					outsideArea = true
					break
				}

				if slice[nextPosition.x][nextPosition.y] == '#' {

					directionsIndex = (directionsIndex + 1) % len(directions) // get next direction of directionsSlice, wrap around if necessary
					direction = directions[directionsIndex]

					turningPos = append(turningPos, position)

					turningPosIdx = (turningPosIdx + 1)

					//      fmt.Println("turnpoints: ", turningPos)
					// fmt.Println("found obstacle")
					// fmt.Println("new direction: ", direction)

				

					continue
				}
				if slice[nextPosition.x][nextPosition.y] != '#' {

					position = nextPosition

				}
        if turningPosIdx > 2 && (hasDuplicates(turningPos[:len(turningPos)-2])) {
					loopcount++
					foundLoop = true	
          // fmt.Println("loopcount: ", loopcount, "  ", i, j)

          // fmt.Println("turning points: ", turningPos)
          // fmt.Println("found loop at: ", i, j)
					//      slice[position.x][position.y] = 'O'
					// strings := make([][]string, len(slice))
					// for i, row := range slice {
					// 	strings[i] = make([]string, len(row))
					// 	for j, r := range row {
					// 		strings[i][j] = string(r)
					// 	}
					// }
					// for _, row := range strings {
					// 	fmt.Println(row)
					// }
					// fmt.Println()
					// slice[position.x][position.y] = '.'
					break
				}
				// fmt.Println("next position: ", position)

			}
			slice[i][j] = '.'
		}
    fmt.Println(i)
	}
  
	return loopcount
}

func hasDuplicates(arr []point) bool {
	seen := make(map[point]bool)
	for _, v := range arr {
		if seen[v] {
			return true // duplicate found
		}
		seen[v] = true
	}
	return false // no duplicate
}



func Day6(input string) (score1, score2 int) {

	// puzzle 1
	score1 = findWays(input)

	// puzzle 2
  score2 = testForLoop(input)
  return score1, score2
}
