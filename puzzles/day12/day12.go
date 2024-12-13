package day12

import (
	// "fmt"

	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

type Coordinates struct {
	X, Y int
}

type PerimeterDetail struct {
	Coordinate Coordinates // Koordinate der aktuellen Zelle
	Direction  Coordinates // Richtung des Perimeterabschnitts
}
type RegionProperties struct {
	Rune             rune
	Area             int
	Perimeter        int
	PerimeterDetails []PerimeterDetail
}

func calculateRegionProperties(matrix [][]rune, startX, startY int, visited map[Coordinates]bool) (int, int, []PerimeterDetail) {
	directions := []Coordinates{ // mögliche Bewegungsrichtungen
		{0, 1},  // rechts
		{0, -1}, // links
		{1, 0},  // unten
		{-1, 0}, // oben
	}

	stack := []Coordinates{{startX, startY}}
	regionRune := matrix[startX][startY]
	visited[Coordinates{startX, startY}] = true

	area := 0
	perimeter := 0
	perimeterDetails := []PerimeterDetail{} // Liste der Perimeterdetails

	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area++ // Jede besuchte Zelle erhöht die Fläche

		for _, dir := range directions {
			newX, newY := cell.X+dir.X, cell.Y+dir.Y
			neighbor := Coordinates{newX, newY}

			if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) {
				// Rand der Matrix zählt zum Perimeter
				perimeter++
				perimeterDetails = append(perimeterDetails, PerimeterDetail{Coordinate: cell, Direction: dir})
			} else if matrix[newX][newY] != regionRune {
				// Andere Region zählt zum Perimeter
				perimeter++
				perimeterDetails = append(perimeterDetails, PerimeterDetail{Coordinate: cell, Direction: dir})
			} else if !visited[neighbor] {
				// Gleiche Region, aber noch nicht besucht
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return area, perimeter, perimeterDetails
}

func findRegions(matrix [][]rune) []RegionProperties {
	visited := make(map[Coordinates]bool)
	var regions []RegionProperties

	for i := range matrix {
		for j := range matrix[i] {
			coord := Coordinates{i, j}
			if !visited[coord] {
				area, perimeter, details := calculateRegionProperties(matrix, i, j, visited)
				region := matrix[i][j]

				// Region zur Liste hinzufügen
				regions = append(regions, RegionProperties{
					Rune:             region,
					Area:             area,
					Perimeter:        perimeter,
					PerimeterDetails: details,
				})
			}
		}
	}

	// Ausgabe
	for i, props := range regions {
		fmt.Printf("Region %d (%c): Fläche = %d, Perimeter = %d\n", i+1, props.Rune, props.Area, props.Perimeter)
		for _, detail := range props.PerimeterDetails {
			fmt.Printf("  Perimeter bei %v in Richtung %v\n", detail.Coordinate, detail.Direction)
		}
	}
  fmt.Println(regions)
	return regions
}

func Day12(input string) int {

	matrix := fileops.FileToIntRune(input)
	// fileops.PrintRuneMatrix(matrix)

  regions :=	findRegions(matrix)
var  puzzle1Sum int
  for _, v := range regions {
    puzzle1Sum += v.Area * v.Perimeter
  }
	return puzzle1Sum
}
