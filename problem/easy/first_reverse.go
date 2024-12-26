package easy

import (
	"fmt"
	"strings"
)

// FirstReverse takes a string and returns the string in reverse order
func FirstReverse(str string) string {
	var result []rune
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result = append(result, strArr[i])
	}
	return string(result)
}

func WordReverseInPlace(str string) string {
	split := strings.Split(str, " ")
	for i, word := range split {
		split[i] = reverse(word)
		fmt.Println(split[i])
	}
	return strings.Join(split, " ")
}

func reverse(str string) string {
	var result []rune
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result = append(result, strArr[i])
	}
	return string(result)
}
