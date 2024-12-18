package hackkerankeasy

func TwoStrings(s1 string, s2 string) string {
	m := make(map[rune]bool)
	for _, c := range s1 {
		m[c] = true
	}
	for _, c := range s2 {
		if m[c] {
			return "YES"
		}
	}
	return "NO"
}
