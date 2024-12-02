package puzzles

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func readFile() []byte {
	body, _ := fileops.ReadFile("input/day3.txt")
	return body
}

// /puzzle 1
func CalculateSum(body []byte) (int, error) {
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
	return sum, nil

}

//puzzle 2


func MakeString(body []byte) string {
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

func Day3() {
	// Read the input file
	body := readFile()

	// Puzzle 1: Calculate the sum
	sum, err := CalculateSum(body)
	if err != nil {
		fmt.Println("Error in Puzzle 1:", err)
		return
	}
	fmt.Println("Day 3 puzzle 1:", sum)

	// Puzzle 2: Transform the string and calculate the sum again
	transformedString := MakeString(body)
	sum2, err2 := CalculateSum([]byte(transformedString))
	if err2 != nil {
		fmt.Println("Error in Puzzle 2:", err2)
		return
	}
	fmt.Println("Day 3 puzzle 2:", sum2)
}

