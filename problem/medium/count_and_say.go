package medium

import "fmt"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return say(CountAndSay(n - 1))
}

func say(s string) string {
	var result string
	var count int
	for i := 0; i < len(s); i++ {
		count++
		if i+1 == len(s) || s[i] != s[i+1] {
			result += fmt.Sprintf("%d", count)
			result += string(s[i])
			count = 0
		}
	}
	return result
}
