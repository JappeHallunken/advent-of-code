package puzzles

import (
	"fmt"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func MakeArrays(body []byte) [][]rune {

	lines := strings.Split(strings.TrimSpace(string(body)), "\n")
	var xmasArray [][]rune

	for _, line := range lines {
		xmasArray = append(xmasArray, []rune(line))

	}
	return xmasArray
}

func FindMatch(array [][]rune) int {
	directions := [8][2]int{
		{0, 1},   //right
		{0, -1},  //left
		{1, 0},   //down
		{-1, 0},  //up
		{1, 1},   //down right
		{1, -1},  //down left
		{-1, 1},  //up right
		{-1, -1}, //up left
	}
	targets := []rune{'X', 'M', 'A', 'S'}
	wordLength := len(targets)
	counter := 0
	for i, row := range array {
		for j, rune := range row {
			if rune == targets[0] {
				for _, direction := range directions {
					x, y := i, j
					targetIndex := 1
					for {
						x += direction[0]
						y += direction[1]
						if x < 0 || x >= len(array) || y < 0 || y >= len(array[0]) {
							break
						}
						if targetIndex >= wordLength {
							break
						}
						if array[x][y] == targets[targetIndex] {
							targetIndex++
							if targetIndex == wordLength {
								counter++
							}
						} else {
							break
						}
					}

				}
			}
		}
	}

	return counter
}

func Day4(input string) {
	body, _ := fileops.ReadFile(input)
	xmasArray := MakeArrays(body)
	counter := FindMatch(xmasArray)
	fmt.Println("Day 4 puzzle 1: ", counter)
}
