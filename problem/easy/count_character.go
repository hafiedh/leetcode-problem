package problem

func CountCharacters(words []string, chars string) int {
	charCount := make(map[rune]int)
	for _, char := range chars {
		charCount[char]++
	}

	var result int
	for _, word := range words {
		if isGoodWord(word, charCount) {
			result += len(word)
		}
	}
	return result
}

func isGoodWord(word string, charCount map[rune]int) bool {
	wordCount := make(map[rune]int)
	for _, char := range word {
		wordCount[char]++
		if wordCount[char] > charCount[char] {
			return false
		}
	}
	return true
}
