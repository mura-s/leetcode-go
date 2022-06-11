package problem26

func removeDuplicates(nums []int) int {
	var size, cursor int
	for cursor < len(nums) {
		num := nums[cursor]
		cursor++
		nums[size] = num
		size++
		for cursor < len(nums) && num == nums[cursor] {
			cursor++
		}
	}
	return size
}
