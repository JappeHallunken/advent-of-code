package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/JappeHallunken/advent-of-code/fileops"
)

func makeSlice(input string) (twoDIntSlice [][]uint64) {

	byteSlice, err := fileops.ReadFile(input)
	if err != nil {
		panic(err)
	}

	content := string(byteSlice)
	content = strings.TrimSpace(content)

	//remove the colon
	cleanedContent := strings.ReplaceAll(content, ":", "")

	lines := strings.Split(cleanedContent, "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		var intSlice []uint64
		for _, field := range fields {
			number, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				return
			}
			intSlice = append(intSlice, uint64(number))
		}
		twoDIntSlice = append(twoDIntSlice, intSlice)
	}
	fmt.Println(twoDIntSlice)
	return twoDIntSlice
}

func findOperators(equations [][]uint64) (pass2Ops, pass3Ops, combinedLines []int) { //finds the valid lines and returns the idx of the lines
	for r := range equations {
		fmt.Printf("%v of %v\n", r+1, len(equations))

		if tryOperators(equations[r]) { // 2 operators
			pass2Ops = append(pass2Ops, r)
			combinedLines = append(combinedLines, r)

		} else {
			if (findCombinations(equations[r])) { // 3 operators
				// fmt.Println("check valid Invald: ", equations[r])
				pass3Ops = append(pass3Ops, r)
        fmt.Println("add line: ", r)
				combinedLines = append(combinedLines, r)
			}
		}

	}
	return pass2Ops, pass3Ops, combinedLines
}

func tryOperators(nums []uint64) bool {
	numbers := nums[1:]
	target := nums[0]
	// fmt.Println()
	// fmt.Println("target: ", target)

	if len(numbers) == 1 {
		return numbers[0] == target
	}

	operatorCount := len(numbers) - 1
	for i := 0; i < (1 << operatorCount); i++ {
		operators := make([]rune, operatorCount)
		for j := 0; j < operatorCount; j++ {
			if (i>>j)&1 == 1 {
				operators[j] = '*'
			} else {
				operators[j] = '+'
			}
		}
		result := calculate(nums, operators)
		if result == target {
			fmt.Printf("true: ")
			fmt.Println("result: ", result)
			fmt.Println("nums: ", nums)
			fmt.Println("operators: ", string(operators))
			printEquation(nums, operators, result)
			return true
		}
		// printEquation(nums, operators, result)
	}
	return false
}

func calculate(nums []uint64, operators []rune) (result uint64) {
	result = nums[1]
	for i, operator := range operators {
		switch operator {
		case '+':
			result += nums[i+2]
		case '*':
			result *= nums[i+2]
		}
	}
	return result
}


func printEquation(nums []uint64, operators []rune, result uint64) {
	equation := strconv.FormatUint(nums[1], 10)
	for i, operator := range operators {
		equation += string(operator) + strconv.FormatUint(nums[i+2], 10)
	}
	fmt.Printf("Equation: %v = %v\n", strconv.FormatUint(result, 10), equation)
}

func calculateSum(validLines []int, equations [][]uint64) uint64 {
	var sum uint64
	sum = 0
	for _, line := range validLines {
		sum += equations[line][0]
	}
	return sum
}

func evaluate(numbers []uint64, operators []rune) uint64 {
	result := numbers[0]

	for i, operator := range operators {
		switch operator {
		case '+':
			result += numbers[i+1]
		case '*':
			result *= numbers[i+1]
		case '|':
			// Concatenate numbers as a string and convert back to uint64
			combined, err := strconv.ParseUint(fmt.Sprintf("%d%d", result, numbers[i+1]), 10, 64)
			if err != nil {
				fmt.Printf("Error concatenating numbers: %v\n", err)
				return 0
			}
			result = combined
		default:
			fmt.Printf("Unknown operator: %c\n", operator)
			return 0
		}
	}
	return result
}

func findCombinations(numbers []uint64) bool {
	operators := []rune{'+', '*', '|'}
  target := numbers[0]
  n := len(numbers[1:]) - 1
	totalCombinations := int(math.Pow(float64(len(operators)), float64(n)))
	// for i := 0; i < n; i++ {
	// 	totalCombinations *= len(operators)
	// }

	// Iterate through all combinations
	for i := 0; i < totalCombinations; i++ {
		comb := make([]rune, n)
		tmp := i
		for j := 0; j < n; j++ {
			comb[j] = operators[tmp%len(operators)]
			tmp /= len(operators)
		}
		// Evaluate and compare
    result := evaluate(numbers[1:], comb)
		if result == target {
			fmt.Printf("Match found with operators: %v\n", string(comb))
			fmt.Printf("%v = %v with operators %v\n", target, numbers, string(comb))
			return true
		} else {
			// fmt.Printf("No match for combination: %v, result: %v\n", string(comb), result)
		}
	}
	return false
}

func Day7(input string) (uint64, uint64, uint64) {

	// puzzle 1
	slice := makeSlice(input)
	pass2, pass3, combined := findOperators(slice)
	fmt.Println("valid lines: ", pass2)
	fmt.Println("valid invlid lines: ", pass3)

	sum := calculateSum(pass2, slice) // with 2 operators

	// puzzle 2

	sum2 := calculateSum(pass3, slice)    // pass 3 Operators
	sum3 := calculateSum(combined, slice) // pass 2 and 3

	fmt.Println("sum: ", sum)
	fmt.Println("sum2: ", sum2)
	fmt.Println("sum3: ", sum3)
	fmt.Println(sum+sum2 == sum3)

	return sum, sum2, sum3
}
