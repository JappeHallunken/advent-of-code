package day8

import (
	"fmt"

	"github.com/JappeHallunken/advent-of-code/fileops"
)







func Day8(input string) int {
	fmt.Println("day8:")

  body, err := fileops.ReadFile(input)
  if err != nil {
    fmt.Println(err)
    return -1
  }

  // coordMap := fileops.ByteToCoordRuneMap(body)
  // fmt.Println(coordMap)
  

  slices := fileops.MakeStringSlice(body)
fileops.PrintMap(slices)


	return 0

}
