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

func Day14(input string) {
	 xLength:= 101
	 yLength:= 103

	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	robots := getRobots(string(body))
	fmt.Println(robots)
  space := createSpace(xLength, yLength, robots)
  fileops.PrintRuneMatrix(space)
}

func getRobots(body string) []Robot {
	var robot Robot
	var robots []Robot
	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	lines := strings.Split(string(body), "\n")
	for i, line := range lines {
		match := pattern.FindStringSubmatch(line)
		if len(match) == 5 {
			fmt.Sscanf(match[1], "%d", &robot.Position.X)
			fmt.Sscanf(match[2], "%d", &robot.Position.Y)
			fmt.Sscanf(match[3], "%d", &robot.Velocity.X)
			fmt.Sscanf(match[4], "%d", &robot.Velocity.Y)
			robots = append(robots, robot)
			fmt.Printf("robot %v; pos: %v, %v; velocity: %v, %v\n", i+1, robot.Position.X, robot.Position.Y, robot.Velocity.X, robot.Velocity.Y)
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
