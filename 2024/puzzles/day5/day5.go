package day5

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JappeHallunken/advent-of-code/2024/fileops"
)

func splitAndMakeSlices(body []byte) (rules, pages [][]int) { // split data in rules and pages slices
	parts := bytes.Split(body, []byte("\n\n"))
	part1 := parts[0]
	part2 := parts[1]

	firstPartStr := string(part1)
	ruleLines := strings.Split(firstPartStr, "\n")
	for _, line := range ruleLines {
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
		rules = append(rules, row)
	}

	secondPartStr := strings.TrimSpace(string(part2)) //trimmspace bc it ends with \n
	pageLines := strings.Split(secondPartStr, "\n")
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
		pages = append(pages, row)
	}
	return rules, pages
}

func getIdxValidInvalid(orderRules, pageNumbers [][]int) (validPageNumberIdx, invalidPageNumberIdx []int) {

	for i := range pageNumbers { // get each array of page numbers,	check if rules is fulfilled
		isValid := true

		for j := range orderRules { //get page order from rule

			minRule := orderRules[j][0]
			maxRule := orderRules[j][1]

			minRuleIdx := slices.Index(pageNumbers[i], minRule)
			maxRuleIdx := slices.Index(pageNumbers[i], maxRule)

			if (minRuleIdx == -1) || (maxRuleIdx == -1) {
				continue
			}
			if minRuleIdx > maxRuleIdx { //invalid
				isValid = false
				break
			}
		}
		if isValid {
			validPageNumberIdx = append(validPageNumberIdx, i) //fullfilled: add to valid slice
		} else {
			invalidPageNumberIdx = append(invalidPageNumberIdx, i) // unfulfilled: add to invalid slice

		}
	}
	// fmt.Println("validPageNumberIdx: ", validPageNumberIdx)
	// fmt.Println("invalidPageNumberIdx: ", invalidPageNumberIdx)
	return validPageNumberIdx, invalidPageNumberIdx
}

func createNmbSlices(pageNumbers [][]int, invalidPageNumberIdx []int) [][]int { //take a slice and extract the given indexes to new slice

	invalidNumbers := make([][]int, len(invalidPageNumberIdx))
	for i := range invalidPageNumberIdx {
		invalidNumbers[i] = pageNumbers[invalidPageNumberIdx[i]]
	}
	return invalidNumbers
}

func fixOrder(rules, invalidNumbers [][]int) (fixedInvalidNumbers [][]int) { // fix the invalid

	changesMade := true // set to true, if it detects on rule break -> set to false and move to next line in invalidNumbers

	for changesMade { // loop until all rules fullfiled

		changesMade = false

		for i := range invalidNumbers { // get each array of page numbers, check if the pages from the rule are in there, if not check next line
			slice := invalidNumbers[i]

			for j := range rules {
				minRule := rules[j][0]
				maxRule := rules[j][1]

				minRuleIdx := slices.Index(slice, minRule) // get the indexes of the rules
				maxRuleIdx := slices.Index(slice, maxRule)

				if (minRuleIdx == -1) || (maxRuleIdx == -1) {
					continue
				}
				if minRuleIdx > maxRuleIdx { // if true then the order is incorrect
					invalidNumbers[i][minRuleIdx], invalidNumbers[i][maxRuleIdx] = //swap
						invalidNumbers[i][maxRuleIdx], invalidNumbers[i][minRuleIdx]
					changesMade = true
				}
			}
		}
	}
	// fmt.Println("Invalid numbers: ", invalidNumbers)
	return invalidNumbers

}

func findMiddleAndSum(pageNumbers [][]int) (sum int) { // find the middle number and sum it, asumes odd number
	for _, page := range pageNumbers {
		length := len(page)
		sum += page[length/2]
	}
	return sum
}

func Day5(input string) (sum, sum2 int) {

	var rules, pages [][]int
	body, err := fileops.ReadFile(input)

	if err != nil {
		fmt.Printf("Day 5: Error reading file: %v\n", err)
		return
	}
	rules, pages = splitAndMakeSlices(body)
	validUpdatesIdx, invalidUpdatesIdx := getIdxValidInvalid(rules, pages)

	invalidNumbers := createNmbSlices(pages, invalidUpdatesIdx)
	validNumbers := createNmbSlices(pages, validUpdatesIdx)
	fixedInvalidNumbers := fixOrder(rules, invalidNumbers)

	// fmt.Println("valid updates Index: ", validUpdatesIdx)
	// fmt.Println("Invalid Updates idx: ", invalidUpdatesIdx)
	// fmt.Println("validNumbers: ", validNumbers)
	// fmt.Println("invalidNumbers:        ", invalidNumbers)
	// fmt.Println("fixedInvalidNumbers:   ", fixedInvalidNumbers)

	sum = findMiddleAndSum(validNumbers)
	sum2 = findMiddleAndSum(fixedInvalidNumbers)

	// fmt.Println("Day 5 puzzle 1: ", sum)

	// sum2 = findMiddleAndSum(validUpdatesIdx, fixedInvalidUpdates)
	return sum, sum2
}
