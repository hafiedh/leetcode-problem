package problem

func TailFact(n int) int {
	return factT(n-1, n)
}

func factT(n, current int) int {
	if n <= 1 {
		return current
	}
	return factT(n-1, n*current)
}
