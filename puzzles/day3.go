package puzzles

import (
	"fmt"
	"github.com/JappeHallunken/advent-of-code/fileops"
	"regexp"
	"strconv"
)

func readFile() []byte {
	body := fileops.ReadFile("input/day3.txt")
	return body
}

func CalculateSum(body []byte) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	// Regex kompilieren
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Fehler beim Kompilieren des Regex:", err)
		return 0
	}

	matches := re.FindAllStringSubmatch(string(body), -1)
	// fmt.Printf("first match: %s, first number: %s, second number: %s\n", matches[0][0], matches[0][1], matches[0][2])

	//sum up
	sum := 0
	for _, match := range matches {
		number1, _ := strconv.Atoi(match[1])
		number2, _ := strconv.Atoi(match[2])
		sum += number1 * number2
	}
	fmt.Println("Day 3 puzzle 1: ", sum)
	return sum

}
func Day3() int {
	body := readFile()
	return CalculateSum(body)
}
