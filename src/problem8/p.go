package problem8

const (
	max = 2147483647
	min = -2147483648
)

func parseNumberString(s string) string {
	if len(s) == 0 {
		return ""
	}

	l := 0
	for ; l < len(s) && s[l] == ' '; l++ {
	}
	if l == len(s) {
		return ""
	}
	if !(s[l] == '+' || s[l] == '-' || (s[l] >= '0' && s[l] <= '9')) {
		return ""
	}

	r := l + 1
	for ; r < len(s) && s[r] >= '0' && s[r] <= '9'; r++ {
	}
	if r-l == 1 && s[l] == '-' {
		return ""
	}
	if s[l] == '+' {
		l++
	}
	return s[l:r]
}

func myAtoi(s string) int {
	pns := parseNumberString(s)
	if len(pns) == 0 {
		return 0
	}

	isMinus := false
	if pns[0] == '-' {
		isMinus = true
		pns = pns[1:]
	}

	idx := 0
	for ; idx < len(pns) && pns[idx] == '0'; idx++ {
	}
	pns = pns[idx:]

	if len(pns) > 10 {
		if isMinus {
			return min
		}
		return max
	}

	ans := 0
	for i := 0; i < len(pns); i++ {
		ans = 10*ans + int(pns[i]-'0')
	}
	if isMinus {
		ans = -ans
	}
	if ans > max {
		return max
	}
	if ans < min {
		return min
	}
	return ans
}
