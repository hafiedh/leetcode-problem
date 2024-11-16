package problem

import "strings"

func WordPattern(pattern string, s string) bool {

	mPatternToS := make(map[byte]string)
	mSToPattern := make(map[string]byte)

	splittedString := strings.Split(s, " ")

	if len(pattern) != len(splittedString) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if mappedChar, exists := mPatternToS[pattern[i]]; exists && mappedChar != splittedString[i] {
			return false
		}

		if mappedPattern, exists := mSToPattern[splittedString[i]]; exists && mappedPattern != pattern[i] {
			return false
		}
		mPatternToS[pattern[i]] = splittedString[i]
		mSToPattern[splittedString[i]] = pattern[i]
	}

	return true
}
