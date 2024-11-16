package problem

// Example 1:

// Input: s = "egg", t = "add"

// Output: true

// Explanation:

// The strings s and t can be made identical by:

// Mapping 'e' to 'a'.
// Mapping 'g' to 'd'.
// Example 2:

// Input: s = "foo", t = "bar"

// Output: false

// Explanation:

// The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.

// Example 3:

// Input: s = "paper", t = "title"

// Output: true

func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if mappedChar, exists := sToT[charS]; exists && mappedChar != charT {
			return false
		}
		if mappedChar, exists := tToS[charT]; exists && mappedChar != charS {
			return false
		}

		sToT[charS] = charT
		tToS[charT] = charS
	}

	return true
}
