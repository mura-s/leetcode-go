package problem13

type Symbol struct {
	S   string
	Val int
}

var symbols = []Symbol{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func romanToInt(s string) int {
	ans := 0
	sIdx := 0
	sLen := len(s)
	for _, symbol := range symbols {
		symbolLen := len(symbol.S)
		for symbolLen <= sLen-sIdx && symbol.S == s[sIdx:sIdx+symbolLen] {
			ans += symbol.Val
			sIdx += symbolLen
		}
	}
	return ans
}
