package problem242

import "sort"

func search(nums []int, target int) int {
	idx := sort.SearchInts(nums, target)
	if idx >= len(nums) {
		return -1
	}
	if nums[idx] != target {
		return -1
	}
	return idx
}
