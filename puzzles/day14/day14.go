package day14

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

type Point struct {
	X, Y int
}

type Robot struct {
	Position Point
	Velocity Point
}

func Day14(input string) int {
	xLength := 101
	yLength := 103
	cycles := 100 // one cycle is one second

	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	robots := getRobots(string(body))
	space := createSpace(xLength, yLength, robots)
	// fileops.PrintRuneMatrix(space)

	space = moveRobots(space, robots, cycles)
	// fmt.Println("After", cycles, "cycles:")
	// fileops.PrintRuneMatrix(space)
	score := calculateSafetyFactor(space)
	return score
}

func getRobots(body string) []Robot {
	var robot Robot
	var robots []Robot
	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		match := pattern.FindStringSubmatch(line)
		if len(match) == 5 {
			fmt.Sscanf(match[1], "%d", &robot.Position.X)
			fmt.Sscanf(match[2], "%d", &robot.Position.Y)
			fmt.Sscanf(match[3], "%d", &robot.Velocity.X)
			fmt.Sscanf(match[4], "%d", &robot.Velocity.Y)
			robots = append(robots, robot)
			// fmt.Printf("robot %v; pos: %v, %v; velocity: %v, %v\n", i+1, robot.Position.X, robot.Position.Y, robot.Velocity.X, robot.Velocity.Y)
		}
	}
	return robots
}

// create the space and place the robots
func createSpace(xLength, yLength int, robots []Robot) [][]rune {
	space := make([][]rune, yLength)
	for i := range space {
		space[i] = make([]rune, xLength)
		for j := range space[i] {
			space[i][j] = '.'
		}
	}
	// place init positions of robots
	for _, robot := range robots {
		position := robot.Position
		if space[position.Y][position.X] == '.' {
			space[position.Y][position.X] = '1'
		} else {
			val := int(space[position.Y][position.X] - '0')
			val++
			space[position.Y][position.X] = rune(val + '0')
		}
	}
	return space
}

func moveRobots(space [][]rune, robots []Robot, cycles int) [][]rune {
	for i := 0; i < cycles; i++ {
		// fmt.Printf("\n\n###### CYCLE %v ######\n", i+1)

		for j, robot := range robots {
			position := robot.Position
			velocity := robot.Velocity
			var newPos Point
			// fmt.Println(len(space[0]), len(space))

			// fmt.Printf("robot %d; pos: %v, %v; velocity: %v, %v\n", j+1, robot.Position.X, robot.Position.Y, robot.Velocity.X, robot.Velocity.Y)
			// fmt.Printf("new position is: \npos x %v + vel x %v = %v\npos y %v + vel y %v = %v\n", position.X, velocity.X, position.X+velocity.X, position.Y, velocity.Y, position.Y+velocity.Y)
			newPos.X = (position.X + velocity.X + len(space[0])) % len(space[0])
			newPos.Y = (position.Y + velocity.Y + len(space)) % len(space)

			// fmt.Println("new Pos is: ", newPos)
			if space[newPos.Y][newPos.X] == '.' {
				space[newPos.Y][newPos.X] = '1'
			} else {
				val := int(space[newPos.Y][newPos.X] - '0')
				val++
				space[newPos.Y][newPos.X] = rune(val + '0')
			}
			if space[position.Y][position.X] != '.' {
				if space[position.Y][position.X] == '1' {
					space[position.Y][position.X] = '.'
				} else {
					val := int(space[position.Y][position.X] - '0')
					val--
					space[position.Y][position.X] = rune(val + '0')
				}
			} else {
				space[position.Y][position.X] = '.'
			}
			robots[j].Position = newPos

			// fileops.PrintRuneMatrix(space)
		}
	}
	return space
}

func calculateSafetyFactor(space [][]rune, startX, endX, startY, endY int) int {
	safetyFactor := 0
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			if space[j][i] != '.' {
				safetyFactor += int(space[j][i] - '0')
			}
		}
	}
	return safetyFactor
}

func getSafetyFactor(space [][]rune, startX, endX, startY, endY int) int {
	lenX := len(space[0])
	lenY := len(space)
	midX, midY := lenX/2, lenY/2

	// Quadranten berechnen
	safetyFactor1 := calculateSafetyFactor(space, 0, midX, 0, midY)      // Quadrant 1
	safetyFactor2 := calculateSafetyFactor(space, midX+1, lenX, 0, midY) // Quadrant 2
	safetyFactor3 := calculateSafetyFactor(space, 0, midX, midY+1, lenY) // Quadrant 3
	safetyFactor4 := calculateSafetyFactor(space, midX+1, lenX, midY+1, lenY)

  safetyFactor := safetyFactor1 * safetyFactor2 * safetyFactor3 * safetyFactor4
  return safetyFactor
}