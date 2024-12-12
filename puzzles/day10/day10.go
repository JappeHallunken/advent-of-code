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
		{X: 1, Y: 0},  // right
		{X: 0, Y: -1}, // down
		{X: -1, Y: 0}, // left
		{X: 0, Y: 1},  // up
	}

	// Map to track visited positions
	visited := make(map[Coordinates]bool)
	var dfs func(currentPos Coordinates, currentValue int) int

	dfs = func(currentPos Coordinates, currentValue int) int {
		time.Sleep(10 * time.Millisecond)
		// Check if out of bounds
		if !isValidCoord(currentPos.X, currentPos.Y, len(topoMap), len(topoMap[0])) {
			return 0
		}
		// Check if the current position has already been visited
		if visited[currentPos] {
			return 0
		}
		// Check if we've found the goal (value = 9)
		if currentValue == 9 {
			// fmt.Printf("Found 9! currentPos: %v", currentPos)
			// mark this position as visited
			visited[currentPos] = true
			return 1 // Found a valid path
		}

		// Variable to count valid paths from the current position
		count := 0

		// Try all 4 possible directions
		for _, dir := range directions {
			currentValue := topoMap[currentPos.X][currentPos.Y]
			nextPos := Coordinates{X: currentPos.X + dir.X, Y: currentPos.Y + dir.Y}

			if isValidCoord(nextPos.X, nextPos.Y, len(topoMap), len(topoMap[0])) {
				nextValue := topoMap[nextPos.X][nextPos.Y]
				if nextValue == currentValue+1 {
					// Continue DFS recursively if nextValue is greater than current value
					count += dfs(nextPos, nextValue)
				} //else {
				// fmt.Println("next pos not valid ", nextValue)
				// fmt.Println("change dir")
				// Skip this direction if nextValue is not greater than current value
				// }
			}
		}
		return count
	}

	// Start DFS from the start position
	return dfs(startPos, topoMap[startPos.X][startPos.Y])
}

func printMatrixColored(matrix [][]int, currentX, currentY, nextX, nextY int) {
	// Define ANSI color codes
	highlightCurrent := "\033[31m" // Green for current value
	highlightNext := "\033[36m"    // Blue for next value
	reset := "\033[0m"             // Reset to default color

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == currentX && j == currentY {
				// Highlight current value
				fmt.Printf("%s%2d%s", highlightCurrent, matrix[i][j], reset)
			} else if i == nextX && j == nextY {
				// Highlight next value
				fmt.Printf("%s%2d%s", highlightNext, matrix[i][j], reset)
			} else {
				// Regular value
				fmt.Printf("%2d", matrix[i][j])
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
	// starts := []Coordinates{{X: 0, Y: 4}}
	fmt.Println(starts)

	totalScore := []int{}
	for _, startPos := range starts {
		fmt.Println("CURRENT START: ", startPos)
		// For each trailhead, call the searchPath function
		subScore := searchPath(topoMap, startPos)
		totalScore = append(totalScore, subScore)

	}
	totalScoreSum := 0
	for j := range totalScore {
		totalScoreSum += totalScore[j]
	}

	// Output the total score (sum of all found paths)
	fmt.Println("Total Score:", totalScoreSum)
	return totalScoreSum
}
