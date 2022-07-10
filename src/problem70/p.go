package problem70

func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[2] = 1
	for i := 1; i < n; i++ {
		dp[i+1] += dp[i]
		dp[i+2] += dp[i]
	}
	return dp[n]
}
