package day5

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func splitAndMakeSlices(body []byte) (orderRules, pageNumbers [][]int) {
	parts := bytes.Split(body, []byte("\n\n"))
	part1 := parts[0]
	part2 := parts[1]

	firstPartStr := string(part1)
	orderLines := strings.Split(firstPartStr, "\n")
	var data [][]int
	for _, line := range orderLines {
		values := strings.Split(line, "|")
		var row []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	secondPartStr := strings.TrimSpace(string(part2))
	pageLines := strings.Split(secondPartStr, "\n")
	var data2 [][]int
	for _, line := range pageLines {
		values := strings.Split(line, ",")
		var row []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			row = append(row, num)
		}
		data2 = append(data2, row)
	}
	// fmt.Println(data)
	return data, data2
}

func checkForRightOrder(orderRules, pageNumbers [][]int) (validPageNumberIdx []int) {

	for numberLine := range pageNumbers { // get each array of page numbers, check if the pages from the rule are in there, if not check next line

		rightOrder := true                   // set to true, if it detects on rule break -> set to false and move to next line in pageNumbers
		for orderLines := range orderRules { //get page order from rule

			beforePageRule := orderRules[orderLines][0]
			afterPageRule := orderRules[orderLines][1]

			beforePageIndex := slices.Index(pageNumbers[numberLine], beforePageRule)
			if beforePageIndex == -1 {
				continue
			}
			afterPageIndex := slices.Index(pageNumbers[numberLine], afterPageRule)
			if afterPageIndex == -1 {
				continue
			}
			if beforePageIndex > afterPageIndex {
				rightOrder = false
				break
			}

		}
		if rightOrder {
			validPageNumberIdx = append(validPageNumberIdx, numberLine)
		}
	}
	// fmt.Println(validPageNumberIdx)
	return validPageNumberIdx
}

func findMiddleAndSum(validPageNumberIdx []int, pageNumbers [][]int) (sum int) {
  for _, i := range validPageNumberIdx {
    length := len(pageNumbers[i])
    sum += pageNumbers[i][length/2]
  }

  return sum
}


func Day5(input string) (sum int){

	var orderRules, pageNumbers [][]int
	body, err := fileops.ReadFile(input)

	if err != nil {
		fmt.Printf("Day 5: Error reading file: %v\n", err)
		return
	}
	orderRules, pageNumbers = splitAndMakeSlices(body)

	validUpdates := checkForRightOrder(orderRules, pageNumbers)
  // fmt.Println("Day 5 puzzle 1: ", validUpdates)

  sum = findMiddleAndSum(validUpdates, pageNumbers)
  // fmt.Println("Day 5 puzzle 2: ", sum)
  return sum

}
