package problem15

import "sort"

type cand struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	numCntMap := make(map[int]int, len(nums))
	for _, n := range nums {
		numCntMap[n]++
	}
	ansMap := make(map[cand]bool, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if c, ok := numCntMap[nums[i]]; !ok || c < 2 {
					continue
				}
			}
			target := -nums[i] - nums[j]
			needCnt := 1
			if target == nums[i] && target == nums[j] {
				needCnt = 3
			} else if target == nums[i] || target == nums[j] {
				needCnt = 2
			}
			if c, ok := numCntMap[target]; ok && c >= needCnt {
				cs := []int{nums[i], nums[j], target}
				sort.Ints(cs)
				can := cand{cs[0], cs[1], cs[2]}
				ansMap[can] = true
			}
		}
	}
	ans := make([][]int, 0, len(ansMap))
	for k := range ansMap {
		ans = append(ans, []int{k.a, k.b, k.c})
	}
	return ans
}
