package main

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
	"github.com/JappeHallunken/advent-of-code/puzzles"
)

func main() {
	/////////// Day 1
	//// puzzle 1
	body := fileops.ReadFile()
	firstElements, secondElements := fileops.MakeSlices(body)

	diff := puzzles.SumDiff(firstElements, secondElements)
	fmt.Println("Day  1 Puzzle 1: ", diff)

	//// puzzle 2

	score := puzzles.SimilarityScore(firstElements, secondElements)
	fmt.Println("Day  1 puzzle 2: ", score)

}
