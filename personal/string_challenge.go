package personal

import (
	"math"
)

// example input coderbyte
// example output etybredoc
// reverse the string

func StringChallenge(s string) string {
	clToken := "g6voq38ab"
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}
	// concat with clToken and replace every fourth character with underscore
	result := string(runes) + clToken
	resultRunes := []rune(result)
	for i := 3; i < len(resultRunes); i += 4 {
		resultRunes[i] = '_'
	}
	return string(resultRunes)
}
func IsFibonacci(num int) string {
	// check if num is in Fibonacci sequence
	if num < 0 {
		return "no"
	}
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	result := "no"
	if b == num || num == 0 {
		result = "yes"
	}
	return result
}

func SearchingNumberInString(s string) int {
	// example input: "hello123world456"
	// example output: 1 + 2 + 3 + 4 + 5 + 6 / (divide by letters count) 10 = 3 (rounded)

	sum := 0
	lettersCount := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			sum += int(char - '0')
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			lettersCount++
		}
	}
	if lettersCount == 0 {
		return 0
	}
	average := float64(sum) / float64(lettersCount)
	return int(math.Round(average))
}
