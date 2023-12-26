package leetWordBreak

// f(n) = f(n-k) && inDict(substr(n-k, n))
// k - min word length
func wordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for r := 1; r <= n; r++ {
		for l := 0; l < r; l++ {
			dp[r] = dp[l] && hash[s[l:r]]
			if dp[r] {
				break
			}
		}
	}
	return dp[n]
}
