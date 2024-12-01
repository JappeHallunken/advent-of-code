package main

import (
	"fmt"

  "github.com/JappeHallunken/advent-of-code/fileops"
  "github.com/JappeHallunken/advent-of-code/puzzles"
)
func main() {

	body := fileops.ReadFile()
	firstElements, secondElements := fileops.MakeSlices(body)
	sortedFirstElements, sortedSecondElements := fileops.SortSlices(firstElements, secondElements)
  
  diff := puzzles.SumDiff(sortedFirstElements, sortedSecondElements)
  fmt.Println("Day  1 Puzzle 1: ", diff)
	score := puzzles.SimilarityScore(firstElements, secondElements)
  fmt.Println("Day  1 puzzle 2: ", score)

}
