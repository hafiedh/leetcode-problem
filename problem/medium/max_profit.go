package medium

func maxProfit(prices []int) int {
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// Find the maximum profit you can achieve. You may complete at most two transactions.

// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
func MaxProfitHard(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][3][2]int, n)
	dp[0][0][0] = 0
	dp[0][1][0] = 0
	dp[0][2][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
	}

	return max(dp[n-1][1][0], dp[n-1][2][0])
}
