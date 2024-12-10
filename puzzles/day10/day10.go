package day10

import (
	"fmt"
	// "time"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func findStarts(topoMap map[fileops.Coordinates]int) []fileops.Coordinates {
	var startMap []fileops.Coordinates
	for i := range topoMap {
		if topoMap[i] == 0 {
			startMap = append(startMap, i)
		}
	}
	fmt.Println("Startpunkte: ", startMap)
	return startMap
}

func isValidCoord(coord fileops.Coordinates, grid map[fileops.Coordinates]int) bool {
	_, exists := grid[coord]
	return exists
}


func searchPath(topoMap [][]int, startPos Coordinates) int {
	directions := []Coordinates{
		{X: 0, Y: 1},  // up
		{X: 0, Y: -1}, // down
		{X: 1, Y: 0},  // right
		{X: -1, Y: 0}, // left
	}

	// Map to track visited positions
	visited := make(map[Coordinates]bool)
	var dfs func(currentPos Coordinates, currentValue int) int

	dfs = func(currentPos Coordinates, currentValue int) int {
		// Check if out of bounds
		if !isValidCoord(currentPos.X, currentPos.Y, len(topoMap), len(topoMap[0])) {
			return 0
		}

		// Check if the current position has already been visited
		if visited[currentPos] {
			return 0
		}

		// Check if we've found the goal (value = 9)
		nextValue := topoMap[currentPos.X][currentPos.Y]
		if nextValue == 9 {
			return 1 // Found a valid path
		}

		// Mark this position as visited
		visited[currentPos] = true

		// Variable to count valid paths from the current position
		count := 0

		// Try all 4 possible directions
		for _, dir := range directions {
			nextPos := Coordinates{X: currentPos.X + dir.X, Y: currentPos.Y + dir.Y}
			if nextValue := topoMap[nextPos.X][nextPos.Y]; nextValue > currentValue {
				// Continue DFS recursively if nextValue is greater than current value
				count += dfs(nextPos, nextValue)
			}
		}

		// Backtrack: Unmark the position as visited to allow other paths
		visited[currentPos] = false

		return count
	}

	// Start DFS from the start position
	return dfs(startPos, topoMap[startPos.X][startPos.Y])
}
	// Iteriere über alle Startpunkte
	for _, startPos := range startCoords {
		// Initialisiere die besuchten Punkte
		visited := make(map[fileops.Coordinates]bool)

		// Starte die Suche für den aktuellen Startpunkt
		count := search(startPos, topoMap[startPos], visited)

		// Füge die Anzahl der gefundenen Pfade zum Resultat-Slice hinzu
		counts = append(counts, count)
	}

	return counts
}
func sumCounts(input []int) int {
	var sum int
	for i := range input {
		sum += input[i]
	}
	return sum
}

func Day10(input string) int {

	fmt.Println("hello day 10")
	topoMap := fileops.FileToIntMap(input)
	fmt.Println("topo map: ", topoMap)

	starts := findStarts(topoMap)
	counts1 := searchPath(topoMap, starts)
	fmt.Println("start points: ", starts)
	sum1 := sumCounts(counts1)
	return sum1
}
