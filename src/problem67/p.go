package problem67

func addBinary(a string, b string) string {
	ans := ""
	var carry byte
	for i := 0; i < len(a) || i < len(b); i++ {
		var ai, bi byte
		if i < len(a) {
			ai = a[len(a)-1-i] - '0'
		}
		if i < len(b) {
			bi = b[len(b)-1-i] - '0'
		}

		sum := ai + bi + carry
		if sum%2 == 1 {
			ans = "1" + ans
		} else {
			ans = "0" + ans
		}
		carry = sum / 2
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}
