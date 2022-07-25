package problem217

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, 0)
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = true
	}
	return false
}
