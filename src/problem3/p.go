package problem3

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLongestSubstring(s string) int {
	foundSet := make(map[byte]bool, 0)
	maxLen := 0
	l, r := 0, 0
	for r < len(s) {
		for r < len(s) {
			if _, ok := foundSet[s[r]]; ok {
				break
			}
			foundSet[s[r]] = true
			r++
		}
		maxLen = max(maxLen, r-l)

		for r < len(s) {
			if _, ok := foundSet[s[r]]; !ok {
				break
			}
			delete(foundSet, s[l])
			l++
		}
	}
	return maxLen
}
