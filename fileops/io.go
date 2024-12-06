package fileops

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	body, err := os.ReadFile(path)

	return  body, err
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

func MakeSlice(body []byte) (slice [][]rune) { //take the textfile and make a 2d array

	lines := strings.Split(strings.TrimSpace(string(body)), "\n")

	for _, line := range lines {
		slice = append(slice, []rune(line))
	}
	return slice
}
