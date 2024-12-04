package tests

import (
	"fmt"
	"testing"

	"github.com/JappeHallunken/advent-of-code/fileops"
	"github.com/JappeHallunken/advent-of-code/puzzles"
)

func TestFindMatch(t *testing.T) {
  t.Run("TestFindMatch", func(t *testing.T) {
    
  
    // Read the test input from a file
    inputFile := "../input/day4_test.txt"
    body, err := fileops.ReadFile(inputFile)
    if err != nil {
        t.Fatal(err)
    }

    // Convert the input data to a 2D array of runes
    xmasArray := puzzles.MakeArrays(body)

    // Call the findMatch function with the input data
    counter := puzzles.FindMatch(xmasArray)

    // Assert on the returned value
    expectedCounter := 18 // Replace with the expected counter value
    if counter != expectedCounter {
        t.Errorf("Expected counter to be %d, but got %d", expectedCounter, counter)
    }
  })
}

func TestFindMatch2(t *testing.T) {
  t.Run("TestFindMatch2", func(t *testing.T) {
    
  
    // Read the test input from a file
    inputFile := "../input/day4_test2.txt"
    body, err := fileops.ReadFile(inputFile)
    if err != nil {
        t.Fatal(err)
    }

    xmasArray := puzzles.MakeArrays(body)

    counter := puzzles.FindMatch2(xmasArray)

    expectedCounter := 9 // Replace with the expected counter value

    if counter != expectedCounter {
        t.Errorf("Expected counter to be %d, but got %d", expectedCounter, counter)
    }else{
      fmt.Println("Day 4 puzzle 2: ", counter)
    }
  })
}
