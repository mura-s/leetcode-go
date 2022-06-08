package problem20

var (
	parenthesesPair = map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
)

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch b := s[i]; b {
		case '(', '{', '[':
			stack = append(stack, b)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			popB := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			pair := parenthesesPair[popB]
			if b != pair {
				return false
			}
		}
	}
	return len(stack) == 0
}
