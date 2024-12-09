package fileops

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func ReadFile(path string) ([]byte, error) {
	body, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, err
}

func MakeSlices(body []byte) (firstElements, secondElements []int) {

	content := string(body)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		elements := strings.Fields(line)
		if len(elements) == 2 {
			num1, err1 := strconv.Atoi(strings.TrimSpace(elements[0]))
			num2, err2 := strconv.Atoi(strings.TrimSpace(elements[1]))
			if err1 == nil && err2 == nil {
				firstElements = append(firstElements, num1)
				secondElements = append(secondElements, num2)
			}
		}
	}
	return firstElements, secondElements
}

func MakeStringSlice(body []byte) (slice []string) {
	content := string(body)
	lines := strings.Split(content, "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) == "" {
			lines = lines[:i]
		} else {
			break
		}
	}
	for _, line := range lines {
		slice = append(slice, strings.TrimSpace(line))
	}
	return slice
}

func MakeSlice(body []byte) (slice [][]rune) { //take the textfile and make a 2d array

	lines := strings.Split(strings.TrimSpace(string(body)), "\n")

	for _, line := range lines {
		slice = append(slice, []rune(line))
	}
	return slice
}

func ByteToCoordRuneMap(body []byte) (points map[Coordinates]rune) {

	points = make(map[Coordinates]rune)

	lines := strings.Split(strings.TrimSpace(string(body)), "\n")

	for i := range lines {
		for j := range lines[i] {
			points[Coordinates{i, j}] = rune(lines[i][j])
		}
	}
	return points
}

func PrintMap(body []string) {
	// we asume every line has same width

	for i := range body {
		for j := range body[i] {
			fmt.Print(string(body[i][j]), " ")
		}
		fmt.Println()
	}
	// for i := range body {
	// 	for j := range body[i] {
	// 		fmt.Print(i,".", j, " ")
	// 	}
	// 	fmt.Println()
	// }
}
