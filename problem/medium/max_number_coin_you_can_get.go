package medium

// There are 3n piles of coins of varying size, you and your friends will take piles of coins as follows:

// In each step, you will choose any 3 piles of coins (not necessarily consecutive).
// Of your choice, Alice will pick the pile with the maximum number of coins.
// You will pick the next pile with the maximum number of coins.
// Your friend Bob will pick the last pile.
// Repeat until there are no more piles of coins.
// Given an array of integers piles where piles[i] is the number of coins in the ith pile.

// Return the maximum number of coins that you can have.

// Example 1:
// Input: piles = [2,4,1,2,7,8]
// Output: 9
// explanation: [2,4,1] -> [2,7,8] -> [4,1,2] -> [7,8]

// Explanation: Choose the triplet (2, 7, 8), Alice Pick the pile with 8 coins, you the pile with 7 coins and Bob the last one.
// Choose the triplet (1, 2, 4), Alice Pick the pile with 4 coins, you the pile with 2 coins and Bob the last one.
// The maximum number of coins which you can have are: 7 + 2 = 9.
// On the other hand if we choose this arrangement (2, 4, 7), (8, 1, 2) you only get 2 + 1 = 3 coins which is not optimal.

// Example 2:
// Input: piles = [2,4,5]
// Output: 4

// Example 3:
// Input: piles = [9,8,7,6,5,1,2,3,4]
// Output: 18
// explanation: [9,8,1] -> [7,6,3] -> [5,4,2]

// MaxCoins return the maximum number of coins that you can have.
// steps:
// 1. sort the piles
// 2. iterate the piles from the end and pick the second element in each iteration
func MaxCoins(piles []int) int {
	quickSort(piles, 0, len(piles)-1)
	var result int
	for i := len(piles) / 3; i < len(piles); i += 2 {
		result += piles[i]
	}
	return result
}

func quickSort(piles []int, low, high int) {
	if low < high {
		p := partition(piles, low, high)
		quickSort(piles, low, p-1)
		quickSort(piles, p+1, high)
	}
}

func partition(piles []int, low, high int) int {
	pivot := piles[high]
	i := low
	for j := low; j < high; j++ {
		if piles[j] < pivot {
			piles[i], piles[j] = piles[j], piles[i]
			i++
		}
	}
	piles[i], piles[high] = piles[high], piles[i]
	return i
}
