package day5

import (
	"fmt"
	"sort"
	"strings"
)

func P1(input string) (int, int) {

	split := strings.Split(input, "\n\n")
	ranges, fresh := split[0], split[1]

	rng := [][]int{}

	for r := range strings.SplitSeq(ranges, "\n") {
		var start, end int
		fmt.Sscanf(r, "%d-%d\n", &start, &end)
		rng = append(rng, []int{start, end})
	}

	var counter int
	for f := range strings.SplitSeq(fresh, "\n") {
		var id int
		fmt.Sscanf(f, "%d\n", &id)
		for _, r := range rng {
			min, max := r[0], r[1]
			// fmt.Printf("Is %v in the range %d-%d? %v\n", id, min, max, (id >= min && id <= max) )
			if id >= min && id <= max {
				counter++
				break
			}
		}
	}

	// PART 2

	sort.Slice(rng, func(i, j int) bool {
		return rng[i][0] < rng[j][0]
	})

	merged := [][]int{rng[0]}

	for _, cur := range rng[1:] {
		last := merged[len(merged)-1]

		if cur[0] <= last[1]+1 {
			// merge
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
			merged[len(merged)-1] = last
		} else {
			merged = append(merged, cur)
		}
	}

	var total int
	for _, m := range merged {
		total += m[1] - m[0] + 1

	}

	return counter, total
}
