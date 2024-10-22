package hard

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).

func IsMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i := 1; i < len(p)+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(p)]
}
