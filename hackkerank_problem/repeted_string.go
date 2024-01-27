package hackkerankproblem

func RepetedString(s string, n int64) int64 {
	if len(s) == 1 && s == "a" {
		return n
	}

	var result int64 = 0

	for _, v := range s {
		if string(v) == "a" {
			result++
		}
	}

	var multiplier int64 = n / int64(len(s)) // 10 / 3 = 3
	result *= multiplier                     // 2 * 3 = 6

	var remainder int64 = n % int64(len(s)) // 10 % 3 = 1
	for i := int64(0); i < remainder; i++ { // 0 < 1
		if string(s[i]) == "a" { // s[0] = a
			result++ // 6 + 1 = 7
		}
	}

	return result
}
