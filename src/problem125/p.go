package problem125

import (
	"unicode"
)

func isAlphaNumeric(b byte) bool {
	if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		for l <= r && !isAlphaNumeric(s[l]) {
			l++
		}
		for l <= r && !isAlphaNumeric(s[r]) {
			r--
		}
		if l >= r {
			break
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}
