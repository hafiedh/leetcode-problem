package hackkerankeasy

import "math"

func MinimumGroups(predators []int32) int32 {
	n := len(predators)
	adjList := make([][]int32, n)
	for i, p := range predators {
		if p != -1 {
			adjList[p] = append(adjList[p], int32(i))
		}
	}

	visited := make([]bool, n)
	maxDepth := int32(0)

	var dfs func(int32, int32)
	dfs = func(node, depth int32) {
		visited[node] = true
		maxDepth = int32(math.Max(float64(maxDepth), float64(depth)))
		for _, v := range adjList[node] {
			if !visited[v] {
				dfs(v, depth+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] && predators[i] == -1 {
			dfs(int32(i), 1)
		}
	}

	return maxDepth
}
