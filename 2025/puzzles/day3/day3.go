package day3

import (
	"strings"
)

func Day3(input string) (int, int) {
	var result1, result2 int

	for bank := range strings.SplitSeq(input, "\n") {
		jolts := make([]int, len(bank))
		for i, r := range bank {
			jolts[i] = int(r - '0')
		}

		result1 += slidingWindowMax(jolts)
		result2 += maxNumberFromSequence(jolts, 12)
	}

	return result1, result2
}

func slidingWindowMax(jolts []int) int {
	k := 2
	max1 := jolts[0]
	var max2, max1Idx int

	for i := 1; i <= len(jolts)-k; i++ {
		if jolts[i] > max1 {
			max1 = jolts[i]
			max1Idx = i
		}
	}

	max2 = jolts[max1Idx+1]
	for i := max1Idx + 2; i < len(jolts); i++ {
		if jolts[i] > max2 {
			max2 = jolts[i]
		}
	}

	return max1*10 + max2
}

func maxNumberFromSequence(jolts []int, k int) int {
	stack := []int{}

	for i, d := range jolts {
		remaining := len(jolts) - i

		for len(stack) > 0 && stack[len(stack)-1] < d && len(stack)+remaining > k {
			stack = stack[:len(stack)-1] 
		}

		if len(stack) < k {
			stack = append(stack, d)
		}
	}

	num := 0
	for _, d := range stack[:k] {
		num = num*10 + d
	}
	return num
}
