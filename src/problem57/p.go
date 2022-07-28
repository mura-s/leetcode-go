package problemw57

import (
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	newS, newE := newInterval[0], newInterval[1]
	setinelIntervals := make([][]int, 0)
	setinelIntervals = append(setinelIntervals, intervals...)
	setinelIntervals = append(setinelIntervals, []int{math.MaxInt, math.MaxInt})
	ans := make([][]int, 0)

	usedNew := false
	i := 0
	for i < len(setinelIntervals) {
		s, e := setinelIntervals[i][0], setinelIntervals[i][1]
		if newE < s {
			if !usedNew {
				ans = append(ans, []int{newS, newE})
				usedNew = true
			}
			ans = append(ans, []int{s, e})
		} else if newS <= e {
			overlapS := min(newS, s)
			for e < newE {
				i++
				s, e = setinelIntervals[i][0], setinelIntervals[i][1]
			}
			if newE < s {
				ans = append(ans, []int{overlapS, newE})
				ans = append(ans, []int{s, e})
			} else {
				ans = append(ans, []int{overlapS, e})
			}
			usedNew = true
		} else {
			ans = append(ans, []int{s, e})
		}
		i++
	}

	return ans[:len(ans)-1]
}
