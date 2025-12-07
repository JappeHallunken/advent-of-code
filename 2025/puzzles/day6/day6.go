package day6

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func P1(input string) int {
	var rows [][]string

	for r := range strings.SplitSeq(input, "\n") {

		fields := strings.Fields(r)
		rows = append(rows, fields)
	}

	var problems [][]string

	for c := range rows[0] {
		var column []string

		for r := range rows {
			v := rows[r][c]
			column = append(column, v)
		}
		problems = append(problems, column)
	}

	var result int
	rowLen := len(problems)

	for i := range rowLen {
		var subResult int
		colLen := len(problems[i])
		op := problems[i][colLen-1]
		// fmt.Printf("computing row %d: %v\n", i, problems[i])
		// fmt.Println("operation is ", op)
		subResult, err := strconv.Atoi(problems[i][0])
		if err != nil {
			fmt.Println(err)
			return -1
		}

		for j := 1; j < colLen-1; j++ {
			num, err := strconv.Atoi(problems[i][j])
			if err != nil {
				return -1
			}
			switch op {
			case "*":
				subResult *= num
			case "+":
				subResult += num
			default:
				// fmt.Println("unrecognized math symbol. doing nothing")
				continue
			}
		}
		result += subResult
	}
	return result
}

func P2(input string) int {

	// get width for each column
	type Column struct {
		Start int
		End   int
		Op    string
	}

	lines := strings.Split(input, "\n")
	opLine := lines[len(lines)-1]
	columns := []Column{}
	i := len(opLine) - 1

	var counter int
	for i >= 0 {
		counter++
		if opLine[i] == '+' || opLine[i] == '*' {
			columns = append(columns, Column{
				Start: i + counter - 1,
				End:   i,
				Op:    string(opLine[i]),
			})
			i -= 2
			counter = 0
		} else {
			i--
		}
	}
	slices.Reverse(columns)
	// fmt.Println(columns)

	nums := [][]string{}

	for _, col := range columns {
		var num []string
		num = append(num, col.Op)
		for i := col.Start; i >= col.End; i-- {
			var n string
			for _, line := range lines[:len(lines)-1] {
				n += string(line[i])
			}
			num = append(num, strings.TrimSpace(n))
			// fmt.Println(num)
		}
		nums = append(nums, num)
	}

	var result int
	for _, num := range nums {
		op := num[0]
		subResult, err := strconv.Atoi(num[1])
		if err != nil {
			fmt.Println(err)
			return -1
		}
		for k := 2; k < len(num); k++ {
			value, err := strconv.Atoi(num[k])
			if err != nil {
				fmt.Println(err)
				return -2
			}
			switch op {
			case "*":
				subResult *= value
			case "+":
				subResult += value
			default:
				fmt.Println("invalid operator")
			}
		}
		result += subResult
	}
	return result
}
