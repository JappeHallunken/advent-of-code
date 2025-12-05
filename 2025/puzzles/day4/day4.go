package day4

import (
	"strings"
)

var directions = [][]int{
	{1, 0},   // right
	{1, 1},   // right down
	{0, 1},   // down
	{-1, 1},  // left down
	{-1, 0},  // left
	{-1, -1}, // left up
	{0, -1},  // up
	{1, -1},  // up right
}

func P1(input string) (int, int) {
	var result1, result2 int
	var grid [][]rune

	for s := range strings.SplitSeq(input, "\n") {
		grid = append(grid, []rune(s))
	}
	_, result1 = checkGrid(grid)

	for {
		var removed int
		grid, removed = checkGrid(grid)
		
		if removed == 0 {
			break
		}
		// fmt.Printf("Removed %d rolls of paper:\n", removed)
		// for row := range grid {
		// 	fmt.Println(string(grid[row]))
		// }
		// fmt.Println()
		result2 += removed
	}

	return result1, result2
}

func checkGrid(grid [][]rune) ([][]rune, int) {

	var result int

	rows := len(grid)
	cols := len(grid[0])
	resGrid := make([][]rune, rows)

	for i := range resGrid {
		resGrid[i] = make([]rune, cols)
		copy(resGrid[i], grid[i]) // initialisiere mit Original
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '@' {
				continue
			}

			counter := 0

			for _, dir := range directions {
				ni := i + dir[0]
				nj := j + dir[1]

				if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if grid[ni][nj] == '@' {
						counter++
					}
				}
			}
			if counter < 4 {
				resGrid[i][j] = 'X'
				result++
			}
		}
	}
	return resGrid, result
}
