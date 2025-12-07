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

	type Column struct {
		Start int
		End   int
		Op    string
	}

	lines := strings.Split(input, "\n")
	opLine := lines[len(lines)-1]
	dataLines := lines[:len(lines)-1]

	// -----------------------------
	// 1. get column widths
	// -----------------------------

	cols := []Column{}
	counter := 0

	for i := len(opLine) - 1; i >= 0; i-- {
		counter++
		if opLine[i] == '+' || opLine[i] == '*' {
			cols = append(cols, Column{
				Start: i + counter - 1,
				End:   i,
				Op:    string(opLine[i]),
			})
			i -= 1
			counter = 0
		}
	}
	slices.Reverse(cols)
	// fmt.Println(cols)

	// -----------------------------
	// 2. extract values
	// -----------------------------

	extracted := make([][]string, len(cols))

	for idx, c := range cols {
		// operator first
		col := []string{c.Op}

		for pos := c.Start; pos >= c.End; pos-- {
			var buf strings.Builder
			for _, ln := range dataLines {
					buf.WriteByte(ln[pos])
			}
			col = append(col, strings.TrimSpace(buf.String()))
		}

		extracted[idx] = col
	}

	// -----------------------------
	// 3. compute
	// -----------------------------

	total := 0

	for _, col := range extracted {
		op := col[0]

		acc, _ := strconv.Atoi(col[1])

		for _, v := range col[2:] {
			n, _ := strconv.Atoi(v)
			if op == "*" {
				acc *= n
			} else {
				acc += n
			}
		}
		total += acc
	}
	return total
}
