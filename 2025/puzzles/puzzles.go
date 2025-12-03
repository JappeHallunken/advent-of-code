package puzzles

import (
	// "advent-of-code-2025/puzzles/day1"
	// "advent-of-code-2025/puzzles/day2"
	"advent-of-code-2025/puzzles/day3"
	"fmt"
)

func SolveAll() {

	// d1p1, _ := day1.P1("./puzzles/day1/input.txt")
	// d1p2, _ := day1.P2("./puzzles/day1/input.txt")
	// fmt.Printf("----- Day 1 -----\nPuzzle 1: %v ||| Puzzle 2: %v\n\n", d1p1, d1p2)
	//
	// d2p1, err := day2.P1("./puzzles/day2/input.txt")
	// if err != nil {
	// 	fmt.Println("D2P1: ", err)
	// 	return
	// }
	// d2p2, err := day2.P2("./puzzles/day2/input.txt")
	// if err != nil {
	// 	fmt.Println("D2P2: ", err)
	// }
	// fmt.Printf("----- Day 2 -----\nPuzzle 1: %v ||| Puzzle 2: %v\n\n", d2p1, d2p2)

	d3p1 := day3.Day3(day3.Input)
	fmt.Printf("----- Day 3 -----\nPuzzle 1: %v ||| Puzzle 2: TBD\n\n", d3p1)
}
