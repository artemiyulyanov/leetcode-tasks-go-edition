package mergeIntervals

import (
	"cmp"
	"slices"
)

func MergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	start, end := -1, -1

	new_intervals := [][]int{}

	for i := 0; i < len(intervals); i++ {
		l, r := intervals[i][0], intervals[i][1]

		if start == -1 && end == -1 {
			start, end = l, r
		}

		if l > end {
			new_intervals = append(new_intervals, []int{start, end})
			start, end = l, r
		} else {
			end = max(end, r)
		}
	}

	new_intervals = append(new_intervals, []int{start, end})

	return new_intervals
}
