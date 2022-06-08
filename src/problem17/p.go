package problem17

var mapping = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	dp := make([][]string, len(digits)+1)
	dp[0] = []string{}
	for i := 0; i < len(digits); i++ {
		cands := mapping[string(digits[i])]
		if len(dp[i]) == 0 {
			dp[i+1] = cands
			continue
		}
		combs := make([]string, 0, len(cands)*len(dp[i]))
		for _, c := range cands {
			for _, prev := range dp[i] {
				combs = append(combs, prev+c)
			}
		}
		dp[i+1] = combs
	}
	return dp[len(digits)]
}
