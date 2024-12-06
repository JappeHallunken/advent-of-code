package puzzles

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/puzzles/day1"
	"github.com/JappeHallunken/advent-of-code/puzzles/day2"
	"github.com/JappeHallunken/advent-of-code/puzzles/day3"
	"github.com/JappeHallunken/advent-of-code/puzzles/day4"
	"github.com/JappeHallunken/advent-of-code/puzzles/day5"
	"github.com/JappeHallunken/advent-of-code/puzzles/day6"
)

func SolveAll() {
	d1p1, d1p2 := day1.Day1("./input/day1.txt")
	fmt.Printf("\nDay 1 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d1p1, d1p2)

	d2p1, d2p2 := day2.Day2("./input/day2.txt")
	fmt.Printf("Day 2 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d2p1, d2p2)

	d3p1, d3p2 := day3.Day3("./input/day3.txt")
	fmt.Printf("Day 3 puzzle 1: %v\n      puzzle 2: %v \n------------------------------\n\n", d3p1, d3p2)

	d4p1, d4p2 := day4.Day4("./input/day4.txt")
	fmt.Printf("Day 4 puzzle 1: %v\nDay 4 puzzle 2: %v \n------------------------------\n\n", d4p1, d4p2)

  d5p1, d5p2 := day5.Day5("./input/day5.txt")
	fmt.Printf("Day 5 puzzle 1: %v\nDay 5 puzzle 2: %v \n------------------------------\n\n", d5p1, d5p2)

  d6p1, d6p2 := day6.Day6("./input/day6.txt")
	fmt.Printf("Day 5 puzzle 1: %v\nDay 5 puzzle 2: TBD \n------------------------------\n\n", d6p1, d6p2 )

}
