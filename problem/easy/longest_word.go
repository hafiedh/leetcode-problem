package problem

import (
	"fmt"
	"regexp"
	"strings"
)

func LongestWord(words string) string {
	// remove punctuation marks with regex
	var result string
	re := regexp.MustCompile(`[^\w\s]`)
	words = re.ReplaceAllString(words, "")

	fmt.Println(words)

	// split words into array
	wordsArr := strings.Split(words, " ")

	for _, word := range wordsArr {
		if len(word) > len(result) {
			result = word
		}
	}
	return result
}
