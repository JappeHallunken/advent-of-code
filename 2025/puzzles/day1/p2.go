package day1

import (
	"advent-of-code-2025/helpers"
	"fmt"
	"strconv"
)

func P2(input string) (int, error) {
	lines, err := helpers.ReadFileToString(input)
	if err != nil {
		return -1, err
	}

	size := 100
	current := 50
	counter := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fmt.Println("start: ", current)
		fmt.Println("rotates: ", line)

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return -1, fmt.Errorf("invalid number in line '%s': %v", line, err)
		}
		l := []rune(line)
		dir := l[0]

		switch dir {
		case 'R':
			// edge case to count not again when it landed on 100 the last round
			if current == 100 {
				current = 0
			}

			counter += (num / size)
			rest := num % size
			if rest+current >= size {
				counter++
				current = current + rest - size
			} else {
				current = current + rest
			}

		case 'L':
			// edge case to count not again when it landed on 0 the last round
			if current == 0 {
				current = 100
			}

			counter += num / size
			rest := num % size
			if current-rest <= 0 {
				counter++
				current = current - rest + size
			} else {
				current = current - rest
			}

		default:
			return -1, fmt.Errorf("invalid direction '%c'", dir)
		}

		fmt.Println("new point: ", current)
		fmt.Println("crossings so far: ", counter)
	}
	return counter, nil
}
