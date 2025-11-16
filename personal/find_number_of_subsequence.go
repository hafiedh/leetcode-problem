package personal

import "fmt"

// example:
// input a = "abcde", words = ["a", "bb", "acd", "ace"]
// expected output: 3
func NumberOfSubsequences(s string, words []string) int {
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}
	fmt.Println("wordMap:", wordMap)
	count := 0
	for word := range wordMap {
		if isSubsequence(s, word) {
			count++
		}
	}
	return count
}

func isSubsequence(s, word string) bool {
	sIndex, wIndex := 0, 0
	for sIndex < len(s) && wIndex < len(word) {
		if s[sIndex] == word[wIndex] {
			wIndex++
		}
		fmt.Println("sIndex:", sIndex, "wIndex:", wIndex)
		sIndex++
	}
	return wIndex == len(word)
}
