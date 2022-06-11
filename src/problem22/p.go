package problem22

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch b := s[i]; b {
		case '(':
			stack = append(stack, b)
		case ')':
			if len(stack) == 0 {
				return false
			}
			popB := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if popB != '(' {
				return false
			}
		}
	}
	return len(stack) == 0
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	for i := 0; i < (1 << (2 * n)); i++ {
		var s string
		v := i
		for j := 0; j < 2*n; j++ {
			if v%2 == 1 {
				s += "("
			} else {
				s += ")"
			}
			v /= 2
		}
		if isValid(s) {
			ans = append(ans, s)
		}
	}

	return ans
}
