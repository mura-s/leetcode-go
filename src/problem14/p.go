package problem14

import "strings"

func longestCommonPrefix(strs []string) string {
	var ans strings.Builder
	idx := 0

Loop:
	for {
		var c byte
		for _, s := range strs {
			if len(s) <= idx {
				break Loop
			}
			if c == 0 {
				c = s[idx]
			} else if c != s[idx] {
				break Loop
			}
		}
		ans.WriteByte(c)
		idx++
	}
	return ans.String()
}
