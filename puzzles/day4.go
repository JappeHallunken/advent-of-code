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
			if rune == targets[0] { //Find the 'X'
				for _, direction := range directions { //iterate over all directions

					x, y := i, j
					targetIndex := 1 //set target to 'M'

					// Check if the remaining word fits in the current direction, if not continue
					if x+direction[0]*(wordLength-1) < 0 ||
						x+direction[0]*(wordLength-1) >= len(array) ||
						y+direction[1]*(wordLength-1) < 0 ||
						y+direction[1]*(wordLength-1) >= len(array[0]) {
						continue
					}

					for {
						x += direction[0] // continue search in the current direction
						y += direction[1]
						if targetIndex >= wordLength { // check if we are at 'S'
							break
						}

						if array[x][y] == targets[targetIndex] { //If we found the next rune, set index to the next one
							targetIndex++
							if targetIndex == wordLength { //if we are at the end of our target array we know, we found every target, so it's a match!
								counter++
							}
						} else {
							break
						}
						// Check if the word would fit in the remaining space
						if x+(wordLength-targetIndex)*direction[0] >= len(array) ||
							y+(wordLength-targetIndex)*direction[1] >= len(array[0]) {
							break
						}
					}

				}
			}
		}
	}

	return counter
}

func FindMatch2(array [][]rune) int {
	counter := 0
	for i := 1; i < len(array)-1; i++ { // Skip the first and last row
		row := array[i]
		for j := 1; j < len(row)-1; j++ { // Skip the first and last elements
			rune := row[j]
			if rune == 'A' {

				if array[i-1][j-1] == 'S' && array[i+1][j+1] == 'M' &&
					array[i-1][j+1] == 'M' && array[i+1][j-1] == 'S' {
					counter++
					continue
				}
				if array[i-1][j-1] == 'S' && array[i+1][j+1] == 'M' &&
					array[i-1][j+1] == 'S' && array[i+1][j-1] == 'M' {
					counter++
					continue
				}
				if array[i-1][j-1] == 'M' && array[i+1][j+1] == 'S' &&
					array[i-1][j+1] == 'M' && array[i+1][j-1] == 'S' {
					counter++
					continue
				}
				if array[i-1][j-1] == 'M' && array[i+1][j+1] == 'S' &&
					array[i-1][j+1] == 'S' && array[i+1][j-1] == 'M' {
					counter++
					continue
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

	counter2 := FindMatch2(xmasArray)
	fmt.Println("Day 4 puzzle 2: ", counter2)
}
