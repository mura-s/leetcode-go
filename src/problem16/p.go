package problem16

import "sort"

type cand struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	ansMap := make(map[cand]bool, 0)
	for i := 0; i < len(nums); i++ {
		numCntMap := make(map[int]int, 0)
		for j := i + 1; j < len(nums); j++ {
			numCntMap[nums[j]]++
		}
		for j := i + 1; j < len(nums); j++ {
			target := -nums[i] - nums[j]
			needCnt := 1
			if target == nums[j] {
				needCnt = 2
			}
			if cnt, ok := numCntMap[target]; ok && cnt >= needCnt {
				cs := []int{nums[i], nums[j], target}
				sort.Ints(cs)
				c := cand{cs[0], cs[1], cs[2]}
				ansMap[c] = true
			}
		}
	}
	ans := make([][]int, 0, len(ansMap))
	for k := range ansMap {
		ans = append(ans, []int{k.a, k.b, k.c})
	}
	return ans
}
