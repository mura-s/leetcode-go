package problem242

func isAnagram(s string, t string) bool {
	mp := map[rune]int{}
	for _, r := range s {
		mp[r]++
	}
	for _, r := range t {
		mp[r]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
