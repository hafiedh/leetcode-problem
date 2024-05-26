package medium

// 1992. Find All Groups of Farmland

// You are given a 0-indexed m x n binary matrix land where a 0 represents a hectare of forested land and a 1 represents a hectare of farmland.

// To keep the land organized, there are designated rectangular areas of hectares that consist entirely of farmland. These rectangular areas are called groups. No two groups are adjacent, meaning farmland in one group is not four-directionally adjacent to another farmland in a different group.

// land can be represented by a coordinate system where the top left corner of land is (0, 0) and the bottom right corner of land is (m-1, n-1). Find the coordinates of the top left and bottom right corner of each group of farmland. A group of farmland with a top left corner at (r1, c1) and a bottom right corner at (r2, c2) is represented by the 4-length array [r1, c1, r2, c2].

// Return a 2D array containing the 4-length arrays described above for each group of farmland in land. If there are no groups of farmland, return an empty array. You may return the answer in any order.

// Example 1:
// Input: land = [[1,0,0],[0,1,1],[0,1,1]]
// Output: [[0,0,0,0],[1,1,2,2]]
// Explanation:
// The first group has top left coordinate (0, 0) and bottom right coordinate (0, 0).
// The second group has top left coordinate (1, 1) and bottom right coordinate (2, 2).

// Example 2:
// Input: land = [[1,1],[1,1]]
// Output: [[0,0,1,1]]
// Explanation:
// The first group has top left coordinate (0, 0) and bottom right coordinate (1, 1).

// Example 3:
// Input: land = [[0]]
// Output: []

func FindFarmland(land [][]int) [][]int {
	var farmlands [][]int

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j] == 1 {
				i1, j1 := i, j
				for i1 < len(land) && land[i1][j] == 1 {
					for j1 < len(land[i1]) && land[i1][j1] == 1 {
						j1++
					}
					i1++
				}
				// mark the farmland as visited
				for i2 := i; i2 < i1; i2++ {
					for j2 := j; j2 < j1; j2++ {
						land[i2][j2] = 0
					}
				}
				// append the result
				farmlands = append(farmlands, []int{i, j, i1 - 1, j1 - 1})
			}
		}
	}

	return farmlands
}
