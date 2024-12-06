package day6

import (
	"fmt"
	"slices"
	"time"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

type point struct {
	x int
	y int
}

var directions = []point{ //-1,0 is top left
		{-1, 0},  // up
		{0, 1},  // right
		{1, 0},   // down
		{0, -1}, // left
	}

var startDirectionIndex = 0

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

	
	body, err := fileops.ReadFile(input)
	if err != nil {
		panic(err)
	}

	slice := fileops.MakeSlice(body) //create 2d rune slice of map
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
		// time.Sleep(10 * time.Millisecond)

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

// func testForLoop() {
//
// 	type point struct {
// 		x int
// 		y int
// 	}
//
// 	var directions = []point{ //0,0 is top left
// 		{-1, 0}, // up
// 		{0, 1},  // right
// 		{1, 0},  // down
// 		{0, -1}, // left
// 	}
//
// }

func Day6(input string) int {

	// puzzle 1
	score1 := findWays(input)

	// puzzle 2

	return score1
}
