package medium

// example input [1, 2, 5], 7 => 2
// example input [1, 2, 5], 9 => 3
// example input [1, 2, 5], 10 => 2
// example input [1, 2, 5], 11 => 3
// example input [2], 3 => -1
// example input [1], 0 => 0

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func CoinChangeRecursive(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	return coinChangeRecursive(coins, amount, make(map[int]int))
}

func coinChangeRecursive(coins []int, amount int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if v, ok := memo[amount]; ok {
		return v
	}

	minCoins := amount + 1
	for _, coin := range coins {
		res := coinChangeRecursive(coins, amount-coin, memo)
		if res == -1 {
			continue
		}
		minCoins = min(minCoins, res+1)
	}

	if minCoins == amount+1 {
		memo[amount] = -1
		return -1
	}

	memo[amount] = minCoins
	return minCoins
}
