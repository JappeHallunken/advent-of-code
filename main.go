package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile() []byte {
	body, err := os.ReadFile("pairs.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return body
}

func makeSlices(body []byte) (firstElements, secondElements []int) {

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

func sortSlices(firstElements, secondElements []int) (sortedFirstElements, sortedSecondElements []int) {
	slices.Sort(firstElements)
	slices.Sort(secondElements)
	return firstElements, secondElements
}

func sumDiff(firstElements, secondElements []int) (sum int) {
	sum = 0
	for i := range firstElements {
		diff := firstElements[i] - secondElements[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}
	// fmt.Println(sum)
	return sum

}

func similarityScore(firstElements, secondElements []int) (score int) {
	score = 0
	for i := range firstElements {
		for j := range secondElements {
			multiplier := 0

			if firstElements[i] == secondElements[j] {
				multiplier++
				// fmt.Println(i, j, multiplier)
			}
			score = score + (firstElements[i] * multiplier)
			// fmt.Println(score)
		}
	}
	return score
}

func main() {

	body := readFile()
	firstElements, secondElements := makeSlices(body)
	sortedFirstElements, sortedSecondElements := sortSlices(firstElements, secondElements)
  
  diff := sumDiff(sortedFirstElements, sortedSecondElements)
  fmt.Println("First Puzzle: ", diff)
	score := similarityScore(firstElements, secondElements)
  fmt.Println("Second Puzzle: ", score)

}
