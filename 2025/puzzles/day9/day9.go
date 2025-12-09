package day9

import (
	"fmt"
	"math"
	"strings"
)

type Coord struct {
	ID   int
	X, Y int
}

func calcRectangleArea(a, b Coord) int {
	x := int(math.Abs(float64(a.X)-float64(b.X))) + 1
	y := int(math.Abs(float64(a.Y)-float64(b.Y))) + 1

	return x * y
}

func parseCoordinates(input string) []Coord {
	var coords []Coord
	i := 1
	for line := range strings.SplitSeq(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		coords = append(coords, Coord{ID: i, X: x, Y: y})
		i++
	}
	return coords
}

func P1(input string) int {
	coords := parseCoordinates(input)
	fmt.Println(coords)
	return 0
}
