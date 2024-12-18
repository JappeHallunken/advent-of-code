package day16

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func Day16(input string) (int, int) {
	grid := fileops.FileToIntRune(input)
	fileops.PrintRuneMatrix(grid)

	rows, cols := len(grid), len(grid[0])
	row, col, found := findStart(grid, 'S')
	if !found {
		fmt.Println("start not found")
		return -1, -1
	} else {
		fmt.Printf("start: (%d, %d)\n", row, col)
	}

	nextNeighbor := neighborIterator(row, col, rows, cols)
	for {
		neighbor, ok := nextNeighbor()
		if !ok {
			break
		}
		// fmt.Printf("valid neighbor: (%d, %d)\n", neighbor[0], neighbor[1])
	}

	return 0, 0
}

func findStart(grid [][]rune, target rune) (int, int, bool) {
	rows, cols := len(grid), len(grid[0])
	iterator := gridIterator(rows, cols)

	for {
		pos, ok := iterator()
		if !ok {
			break
		}

		row, col := pos[0], pos[1]
		if grid[row][col] == target {
			return row, col, true
		}
	}

	return -1, -1, false
}

func gridIterator(rows, cols int) func() ([2]int, bool) {
	row, col := 0, 0

	return func() ([2]int, bool) {
		if row >= rows {
			return [2]int{}, false
		}

		pos := [2]int{row, col}
		col++
		if col >= cols {
			col = 0
			row++
		}
		return pos, true
	}
}

func neighborIterator(row, col, rows, cols int) func() ([2]int, bool) {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	i := 0
	return func() ([2]int, bool) {
		for i < len(directions) {
			nr, nc := row+directions[i][0], col+directions[i][1]
			i++
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				return [2]int{nr, nc}, true
			}
		}
		return [2]int{}, false
	}
}
