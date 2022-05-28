package problem7

func reverse(x int) int {
	max := 2147483647
	min := -2147483648

	var output int
	for x != 0 {
		d := x % 10
		x /= 10
		if output > max/10 || (output == (max/10) && d > (max%10)) {
			output = 0
			break
		}
		if output < min/10 || (output == (min/10) && d < (min%10)) {
			output = 0
			break
		}
		output = 10*output + d
	}
	return output
}
