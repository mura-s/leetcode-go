package problem11

func calcArea(height []int, i, j int) int {
	if height[i] < height[j] {
		return height[i] * (j - i)
	}
	return height[j] * (j - i)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0

	for i < j {
		ans = max(ans, calcArea(height, i, j))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}
