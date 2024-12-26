package hackkerankeasy

import "sort"

func SherlockAndAnagrams(s string) int32 {
	var result int32
	var subStrMap = make(map[string]int32)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subStr := s[i:j]
			subStr = sortString(subStr)
			subStrMap[subStr]++
		}
	}

	for _, v := range subStrMap {
		result += v * (v - 1) / 2
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
