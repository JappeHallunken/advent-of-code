package day10

import (
	"fmt"
	"time"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

// Coordinates struct to store the X and Y positions on the map.
type Coordinates struct {
	X, Y int
}

// Helper function to determine if a coordinate is within bounds.
func isValidCoord(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// DFS function to find all paths to 9 starting from a given position.
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
			fmt.Println("out of bounds")
			return 0
		}
		fmt.Println("valid position")
		// Check if the current position has already been visited
		if visited[currentPos] {
			fmt.Println("visited")
			return 0
		}
		// Check if we've found the goal (value = 9)
		if currentValue == 9 {
			fmt.Println("goal")
			fmt.Printf("Zählung! currentPos: %v, currentValue: %v\n", currentPos, currentValue)
			return 1 // Found a valid path
		}

		// Mark this position as visited
		visited[currentPos] = true
		fmt.Println("unvisited ", currentPos)

		// Variable to count valid paths from the current position
		count := 0

		// Try all 4 possible directions
		for _, dir := range directions {

			nextValue := topoMap[currentPos.X][currentPos.Y]
			nextPos := Coordinates{X: currentPos.X + dir.X, Y: currentPos.Y + dir.Y}
			// fmt.Println("current pos: ", currentPos, currentValue)
			// fmt.Println("next pos: ", nextPos, nextValue)

			printMatrixColored(topoMap, currentPos.X, currentPos.Y, nextPos.X, nextPos.Y)
			fmt.Println("current pos: ", currentPos, currentValue)
			fmt.Println("next pos: ", nextPos, nextValue)
			fmt.Println()
			if isValidCoord(nextPos.X, nextPos.Y, len(topoMap), len(topoMap[0])) {

				if nextValue := topoMap[nextPos.X][nextPos.Y]; nextValue > currentValue {

					// Continue DFS recursively if nextValue is greater than current value
					count += dfs(nextPos, nextValue)
					fmt.Println("count", count, currentPos, currentValue, nextPos, nextValue)
				}
			}
		}

		// Backtrack: Unmark the position as visited to allow other paths
		defer func() { visited[currentPos] = false }()

		return count
	}

	// Start DFS from the start position
	return dfs(startPos, topoMap[startPos.X][startPos.Y])
}

func printMatrixColored(matrix [][]int, currentX, currentY, nextX, nextY int) {
	// Define ANSI color codes
	highlightCurrent := "\033[32m" // Green for current value
	highlightNext := "\033[34m"    // Blue for next value
	reset := "\033[0m"             // Reset to default color

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == currentX && j == currentY {
				// Highlight current value
				fmt.Printf("%s%2d%s ", highlightCurrent, matrix[i][j], reset)
			} else if i == nextX && j == nextY {
				// Highlight next value
				fmt.Printf("%s%2d%s ", highlightNext, matrix[i][j], reset)
			} else {
				// Regular value
				fmt.Printf("%2d ", matrix[i][j])
			}
		}
		fmt.Println() // Newline after each row
	}
}
func findStartPositions(topoMap [][]int) []Coordinates {
	startPositions := []Coordinates{}
	for i := range topoMap {
		for j := range topoMap[i] {
			if topoMap[i][j] == 0 {
				startPositions = append(startPositions, Coordinates{i, j})
			}
		}
	}
	return startPositions
}

func Day10(input string) int {
	topoMap := fileops.FileToIntInt(input)
	fileops.PrintMatrix(topoMap)
	// List of start positions (trailheads)
	starts := findStartPositions(topoMap)
	fmt.Println(starts)

	totalScore := []int{}
	for _, startPos := range starts {
		fmt.Println("CURRENT START: ", startPos)
		time.Sleep(1 * time.Second)
		// For each trailhead, call the searchPath function
		subScore := searchPath(topoMap, startPos)
		totalScore = append(totalScore, subScore)

	}
	fmt.Println(totalScore)
	totalScoreSum := 0
	for j := range totalScore {
		totalScoreSum += totalScore[j]
	}

	// Output the total score (sum of all found paths)
	fmt.Println("Total Score:", totalScore)
	return totalScoreSum
}
