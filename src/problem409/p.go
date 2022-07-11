package problem409

func longestPalindrome(s string) int {
	mp := make(map[byte]int, 0)
	ans := 0
	for i := 0; i < len(s); i++ {
		if cnt := mp[s[i]]; cnt > 0 {
			delete(mp, s[i])
			ans += 2
		} else {
			mp[s[i]] = 1
		}
	}
	if len(mp) > 0 {
		ans++
	}
	return ans
}
