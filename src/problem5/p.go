package problem5

func findPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return s[l+1 : r]
}

func longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		// odd
		l, r := i, i
		oddPalindrome := findPalindrome(s, l, r)
		if len(result) < len(oddPalindrome) {
			result = oddPalindrome
		}

		// even
		l, r = i, i+1
		evenPalindrome := findPalindrome(s, l, r)
		if len(result) < len(evenPalindrome) {
			result = evenPalindrome
		}
	}
	return result
}
