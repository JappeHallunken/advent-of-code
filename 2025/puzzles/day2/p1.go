package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ranges []string

func getRanges(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file, %v", err)
	}

	line := string(content)
	ranges = strings.Split(line, ",")
	return ranges, nil
}

func P1(path string) (int, error) {
	ranges, err := getRanges(path)
	if err != nil {
		return -1, fmt.Errorf("error retreaving ranges from %s: %v", path, err)
	}

	var start, end int
	result := 0

	for _, r := range ranges {
		// fmt.Println("checking range ", r)
		_, err := fmt.Sscanf(r, "%d-%d", &start, &end)
		if err != nil {
			return -1, fmt.Errorf("error scanning range %s: %v", r, err)
		}

		for i := start; i <= end; i++ {
			st := strconv.Itoa(i)
			if len(st)%2 != 0 {
				continue
			}
			sub1 := st[:(len(st) / 2)]
			sub2 := st[(len(st) / 2):]
			if sub1 == sub2 {
				result += i
			}
		}
	}
	return result, nil
}

func P2(path string) (int, error) {
	ranges, err := getRanges(path)
	if err != nil {
		return -1, fmt.Errorf("error retreaving ranges from %s: %v", path, err)
	}
	var start, end int
	result := 0
	for _, r := range ranges {
		// fmt.Println("checking range ", r)
		_, err := fmt.Sscanf(r, "%d-%d", &start, &end)
		if err != nil {
			return -1, fmt.Errorf("error scanning range %s: %v", r, err)
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			s2 := s + s
			if strings.Contains(s2[1:len(s2)-1], s) {
				result += i
			}
		}
	}
	return result, nil
}
