package day1

import (
	"advent-of-code-2025/helpers"
	"fmt"
)

func P1(input string) (int, error) {
	lines, err := helpers.ReadFileToString(input)
	if err != nil {
		return -1, err
	}

	size := 100
	current := 50
	counter := 0
	for _, line := range lines {
		// fmt.Println("old value: ", current)
		var prefix rune
		var num int

		_, err := fmt.Sscanf(line, "%c%d", &prefix, &num)
		if err != nil {
			return -1, fmt.Errorf("error parsing line, incorrect format")
		}

		switch prefix {
		case 'R':
			current += num
			current %= size

		case 'L':
			current = (current - num) % size
			if current < 0 {
				current += size
			}

		default:
			return -1, fmt.Errorf("invalid instruction")
		}
		// fmt.Println("instruction: ", line)
		// fmt.Println("new value: ", current)
		if current == 0 {
			counter++
		}
		// fmt.Println("counter: ", counter)
	}
	return counter, nil
}
