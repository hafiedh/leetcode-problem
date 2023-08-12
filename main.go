package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	x := []string{"florida", "flores", "flrid", "florf"}
	fmt.Println(longestCommonPrefix(x))
}

func findDuplicateWords(input string) []string {
	words := strings.Fields(input)
	fmt.Println("Words in the input string:", words)
	wordCount := make(map[string]int)

	// Count the occurrences of each word
	for _, word := range words {
		wordCount[word]++
	}

	// Collect the duplicate words
	var duplicateWords []string
	for word, count := range wordCount {
		if count > 1 {
			duplicateWords = append(duplicateWords, word)
		}
	}

	return duplicateWords
}

func reverseOrder(words []string) []string {
	length := len(words)
	reverseWords := make([]string, length)

	for i, word := range words {
		fmt.Println("Word:", word, "Index:", i)
		reverseWords[length-i-1] = word
		fmt.Println("Reverse word:", reverseWords[length-i-1])
	}

	return reverseWords
}

func printWords(words []string) {
	fmt.Println("Duplicate words in reverse order:")
	for _, word := range words {
		reverseWord := reverseString(word)
		fmt.Println(reverseWord)
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}
	return string(runes)
}

func zigzagPattern(s string, numrows int) string {
	if numrows == 1 {
		return s
	}
	rows := make([]string, numrows)
	row := 0
	step := 1
	for _, c := range s {
		rows[row] += string(c)
		if row == 0 {
			step = 1
		} else if row == numrows-1 {
			step = -1
		}
		row += step
	}
	return strings.Join(rows, "")
}

func reverseInteger(x int) int {
	overflow := false
	underflow := false
	if x < 0 {
		underflow = true
		x = -x
	}
	var result int
	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x /= 10
	}
	if underflow {
		result = -result
	}
	if result > 2147483647 || result < -2147483648 {
		overflow = true
	}
	if overflow {
		return 0
	}
	return result
}

func myAtoi(str string) int {
	var result int
	var sign int
	var overflow bool
	var underflow bool
	var start bool
	for _, c := range str {
		if c == ' ' && !start {
			continue
		}
		if c == '-' && !start {
			sign = -1
			start = true
			continue
		}
		if c == '+' && !start {
			sign = 1
			start = true
			continue
		}
		if c >= '0' && c <= '9' {
			start = true
			if sign == 0 {
				sign = 1
			}
			result = result*10 + int(c-'0')
			if sign*result > math.MaxInt32 {
				overflow = true
				break
			}
			if sign*result < math.MinInt32 {
				underflow = true
				break
			}
		} else {
			break
		}
	}
	if overflow {
		return math.MaxInt32
	}
	if underflow {
		return math.MinInt32
	}
	return sign * result
}

func palindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	var result int
	original := x
	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x /= 10
	}
	return result == original
}

func isMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func maxArea(height []int) int {
	var max int
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			max = int(math.Max(float64(max), float64(height[left]*(right-left))))
			left++
		} else {
			max = int(math.Max(float64(max), float64(height[right]*(right-left))))
			right--
		}
	}
	return max
}

func intToRoman(num int) string {
	var result string
	var roman []string
	roman = append(roman, "M", "CM", "D", "CD", "C", "XC", "L",
		"XL", "X", "IX", "V", "IV", "I")
	integer := []int{1000, 900, 500, 400, 100, 90, 50,
		40, 10, 9, 5, 4, 1}
	for i := 0; i < len(integer); i++ {
		for num >= integer[i] {
			num -= integer[i]
			result += roman[i]
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}
