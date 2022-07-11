package problem169

func majorityElement(nums []int) int {
	mp := make(map[int]int, 0)
	for _, n := range nums {
		mp[n]++
		if mp[n] > (len(nums) / 2) {
			return n
		}
	}
	return -1
}
