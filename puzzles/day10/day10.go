package day10

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func Day10(input string) {
  fmt.Println("hello day 10")
  topoMap := fileops.FileToIntMap(input)
  fmt.Println(topoMap)
}
  
