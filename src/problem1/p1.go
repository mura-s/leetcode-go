package problem1

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int, len(nums))
	for i, v := range nums {
		targetMap[target-v] = i
	}
	for i, v := range nums {
		if j, ok := targetMap[v]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}
