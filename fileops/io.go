package fileops

import (
  "strconv"
  "strings"
  "os"
  "log"
)

func ReadFile(path string) []byte {
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return body
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
