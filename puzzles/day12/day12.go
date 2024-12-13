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
	Coordinate Coordinates
	Direction  Coordinates
}
type RegionProperties struct {
	Rune             rune
	Area             int
	Perimeter        int
	PerimeterDetails []PerimeterDetail // coordinate and direction of the perimeter area
	Coords           []Coordinates     // all coordinates of an area
	Sides            int
}

func calculateRegionProperties(matrix [][]rune, startX, startY int, visited map[Coordinates]bool) (int, int, []PerimeterDetail, []Coordinates) {
	directions := []Coordinates{
		{0, 1},  // right
		{-1, 0}, // down
		{0, -1}, // left
		{1, 0},  // up

	}

	stack := []Coordinates{{startX, startY}}
	regionRune := matrix[startX][startY]
	visited[Coordinates{startX, startY}] = true

	area := 0
	perimeter := 0
	perimeterDetails := []PerimeterDetail{}
	var coords []Coordinates
	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area++
		visited[cell] = true
		coords = append(coords, cell)

		for _, dir := range directions {
			newX, newY := cell.X+dir.X, cell.Y+dir.Y
			neighbor := Coordinates{newX, newY}

			if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) {

				perimeter++
				perimeterDetails = append(perimeterDetails, PerimeterDetail{Coordinate: cell, Direction: dir})
			} else if matrix[newX][newY] != regionRune {

				perimeter++
				perimeterDetails = append(perimeterDetails, PerimeterDetail{Coordinate: cell, Direction: dir})
			} else if !visited[neighbor] {

				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return area, perimeter, perimeterDetails, coords
}

func findRegions(matrix [][]rune) []RegionProperties {
	visited := make(map[Coordinates]bool)
	var regions []RegionProperties
	for i := range matrix {
		for j := range matrix[i] {
			coord := Coordinates{i, j}
			if !visited[coord] {
				area, perimeter, details, coords := calculateRegionProperties(matrix, i, j, visited)
				region := matrix[i][j]
				// coords = append(coords, coord)

				// Region zur Liste hinzufügen
				regions = append(regions, RegionProperties{
					Rune:             region,
					Area:             area,
					Perimeter:        perimeter,
					PerimeterDetails: details,
					Coords:           coords,
				})
			}
		}
	}

	// Ausgabe
	for i, props := range regions {
		fmt.Printf("Region %d (%c): Fläche = %d, Perimeter = %d, Coords: %v\n", i+1, props.Rune, props.Area, props.Perimeter, props.Coords)
		for _, detail := range props.PerimeterDetails {
			perimeter := Coordinates{X: detail.Coordinate.X + detail.Direction.X, Y: detail.Coordinate.Y + detail.Direction.Y}
			fmt.Printf("  Perimeter bei %v in Richtung %v. Perimeter cell: %v\n", detail.Coordinate, detail.Direction, perimeter)
			// fmt.Printf("%v \n", detail.Coordinate )
		}
	}
	// fmt.Println(regions)
	return regions
}

func calculateSides(region RegionProperties) int {
	// Wenn weniger als 2 Punkte vorhanden sind, können keine Seiten berechnet werden
	perimeterDetails := region.PerimeterDetails
	if len(perimeterDetails) < 2 {
		return 0
	}

	sides := 1 // Es gibt immer mindestens eine Seite, wenn es 2 oder mehr Punkte gibt
	var lastChange string

	// Start bei i := 1, um die erste Änderung zu überprüfen
	for i := 1; i < len(perimeterDetails); i++ {
		prev := perimeterDetails[i-1]
		current := perimeterDetails[i]

		// Überprüfen, welche Koordinate sich geändert hat
		var change string

		// Überprüfen, ob sich nur X oder nur Y ändert
		if current.Coordinate.X != prev.Coordinate.X && current.Coordinate.Y == prev.Coordinate.Y {
			// X ändert sich, Y bleibt gleich -> horizontale Bewegung
			change = "X"
		} else if current.Coordinate.X == prev.Coordinate.X && current.Coordinate.Y != prev.Coordinate.Y {
			// Y ändert sich, X bleibt gleich -> vertikale Bewegung
			change = "Y"
		}

		// Wenn sich die Koordinaten ändern und die Richtung (X oder Y) sich geändert hat, bedeutet das einen Richtungswechsel
		if change != "" {
			if lastChange == "" {
				// Beim ersten Schritt keine Richtung festgelegt
				lastChange = change
			} else if lastChange != change {
				// Richtungswechsel festgestellt
				fmt.Println("Richtungswechsel bei:", current.Coordinate)
				sides++             // Eine neue Seite beginnt
				lastChange = change // Die Richtung nach dem Wechsel festlegen
			}
		}
	}

	fmt.Printf("Region %v hat %v Seiten.\n", string(region.Rune), sides)
	return sides
}

func Day12(input string) (sum1, sum2 int) {

	matrix := fileops.FileToIntRune(input)
	fileops.PrintRuneMatrix(matrix)

	regions := findRegions(matrix)

	// fmt.Println(regions)

	for _, v := range regions {
		sum1 += v.Area * v.Perimeter
		sum2 += calculateSides(v) * v.Area
	}

	return sum1, sum2
}
