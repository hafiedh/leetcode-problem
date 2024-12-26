package hackkerankeasy

func HowManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var result int32 = 0
	for s >= p {
		result++
		s -= p
		p -= d
		if p < m {
			p = m
		}
	}
	return result
}
