package day9

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	X, Y int
}

type Coord struct {
	ID    int
	Point Point
}

type CMap map[int]Point

func parseCoordinates(input string) []Coord {
	var coords []Coord
	i := 1
	for line := range strings.SplitSeq(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		p := Point{X: x, Y: y}
		coords = append(coords, Coord{ID: i, Point: p})
		i++
	}
	return coords
}

func calcRectangleArea(a, b Coord) int {
	x := int(math.Abs(float64(a.Point.X)-float64(b.Point.Y))) + 1
	y := int(math.Abs(float64(a.Point.X)-float64(b.Point.Y))) + 1

	return x * y
}

// identifyExtremePoints returns an array with the 4 outer-most Coordinate IDs
func identifyExtremePoints(coords []Coord) [4]int {
	minX, maxX := coords[0].Point.X, coords[0].Point.X
	minY, maxY := coords[0].Point.Y, coords[0].Point.Y

	for _, ch := range coords {
		if ch.Point.X < minX {
			minX = ch.Point.X
		}
		if ch.Point.X > maxX {
			maxX = ch.Point.X
		}
		if ch.Point.Y < minY {
			minY = ch.Point.Y
		}
		if ch.Point.Y > maxY {
			maxY = ch.Point.Y
		}
	}

	var a, b, c, d int

	for _, ch := range coords {
		if ch.Point.X == minX && ch.Point.Y == minY {
			a = ch.ID
		}
		if ch.Point.X == maxX && ch.Point.Y == minY {
			b = ch.ID
		}
		if ch.Point.X == minX && ch.Point.Y == maxY {
			c = ch.ID
		}
		if ch.Point.X == maxX && ch.Point.Y == maxY {
			d = ch.ID
		}
	}

	return [4]int{a, b, c, d}
}

func P1(input string) int {
	coords := parseCoordinates(input)
	fmt.Println(coords)

	ep := identifyExtremePoints(coords)
	fmt.Println(ep)

	aMax := 0
	for i := 0; i < len(ep); i++ {
		for j := i + 1; j < len(ep); j++ {

			area := calcRectangleArea(ep[i], ep[j])
			if v > aMax {
				aMax = v
			}
		}

	}
	return 0
}
