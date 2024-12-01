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
  fmt.Println("First Puzzle: ", diff)
	score := puzzles.SimilarityScore(firstElements, secondElements)
  fmt.Println("Second Puzzle: ", score)

}
