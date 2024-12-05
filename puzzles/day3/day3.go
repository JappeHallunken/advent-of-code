package day3

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

// /puzzle 1
func calculateSum(body []byte) (int, error) {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	// Regex kompilieren
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Fehler beim Kompilieren des Regex:", err)
		return 0, fmt.Errorf("Fehler beim Kompilieren des Regex: %v", err)
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
	// fmt.Println("Day 3 puzzle 1: ", sum)
	return sum, nil

}

//puzzle 2

func makeString(body []byte) string {
	// Create a buffer to store the resulting string
	var result bytes.Buffer
	// Start iterating using a while-like loop
	i := 0
	inCapture := true // Start capturing immediately
	for i < len(body) {
		if i+6 < len(body) && string(body[i:i+7]) == "don't()" {
			// Stop capturing when "don't()" is encountered
			inCapture = false
			i += 7 // Skip the "don't()" string
		} else if i+4 < len(body) && string(body[i:i+4]) == "do()" {
			// Resume capturing when "do()" is encountered
			inCapture = true
			i += 4 // Skip the "do()" string
		} else {
			// Append the character to the result if in capture mode
			if inCapture {
				result.WriteByte(body[i])
			}
			i++
		}
	}
	// Debugging output
	// fmt.Println("Final string:", result.String())
	return result.String()
}

// ///"main"

func Day3(input string) (sum1, sum2 int) {
	// Read the input file
	body, err := fileops.ReadFile(input)
	if err != nil {
		fmt.Printf("Day 3: Error reading file: %v\n", err)
		return 0, 0 // Falls ein Fehler auftritt, Rückgabe von Standardwerten
	}

	// Puzzle 1: Calculate the sum
	sum1, err = calculateSum(body)
	if err != nil {
		fmt.Println("Error in Puzzle 1: ", err)
		return 0, 0 // Fehlerbehandlung, Rückgabe von Standardwerten
	}
	// fmt.Println("Day 3 puzzle 1: ", sum1)

	// Puzzle 2: Transform the string and calculate the sum again
	transformedString := makeString(body)
	sum2, err = calculateSum([]byte(transformedString))
	if err != nil {
		fmt.Println("Error in Puzzle 2: ", err)
		return sum1, 0 // Rückgabe von sum1 und Standardwert für sum2
	}
	// fmt.Println("Day 3 puzzle 2: ", sum2)

	// Am Ende werden die aktualisierten Werte von sum1 und sum2 zurückgegeben
	return sum1, sum2
}
