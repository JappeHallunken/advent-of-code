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

func parseCoordinates(input string) CMap {
	coords := make(CMap, 0)
	i := 1
	for line := range strings.SplitSeq(input, "\n") {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			fmt.Println("error parsing line: ", err)
			continue
		}
		p := Point{X: x, Y: y}
		coords[i] = p
		// coords = append(coords, Coord{ID: i, Point: p})
		i++
	}
	return coords
}

func calcRectangleArea(a, b Point) int {
	x := int(math.Abs(float64(a.X)-float64(b.X))) + 1
	y := int(math.Abs(float64(a.Y)-float64(b.Y))) + 1

	return x * y
}

// identifyExtremePoints returns an array with the 4 outer-most Coordinate IDs
func identifyExtremePoints(coords CMap) [4]int {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt

	for _, p := range coords {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	var a, b, c, d int
	used := make(map[int]bool)

	// distance to extreme point
	distance := func(p Point, targetX, targetY int) int {
		dx := p.X - targetX
		dy := p.Y - targetY
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	}

	// choose the 4 extreme points
	selectClosest := func(targetX, targetY int) int {
		bestID := -1
		bestDist := math.MaxInt
		for id, p := range coords {
			if used[id] {
				continue
			}
			d := distance(p, targetX, targetY)
			if d < bestDist {
				bestDist = d
				bestID = id
			}
		}
		used[bestID] = true
		return bestID
	}

	a = selectClosest(minX, minY) // links-oben
	b = selectClosest(maxX, minY) // rechts-oben
	c = selectClosest(minX, maxY) // links-unten
	d = selectClosest(maxX, maxY) // rechts-unten

	return [4]int{a, b, c, d}
}
func P1(input string) int {
	coords := parseCoordinates(input)
	// fmt.Println(coords)

	ep := identifyExtremePoints(coords)
	// fmt.Println(ep)

	aMax := 0
	for i := range ep {
		for j := i + 1; j < len(ep); j++ {
			a := ep[i]
			b := ep[j]

			area := calcRectangleArea(coords[a], coords[b])
			if area > aMax {
				aMax = area
			}
		}

	}
	return aMax
}
