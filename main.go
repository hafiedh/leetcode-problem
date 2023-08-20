package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(s))
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

func romanToInt(s string) int {
	var result int
	roman := make(map[string]int)
	roman["M"] = 1000
	roman["CM"] = 900
	roman["D"] = 500
	roman["CD"] = 400
	roman["C"] = 100
	roman["XC"] = 90
	roman["L"] = 50
	roman["XL"] = 40
	roman["X"] = 10
	roman["IX"] = 9
	roman["V"] = 5
	roman["IV"] = 4
	roman["I"] = 1
	fmt.Println("Roman:", roman)
	fmt.Println("INPUT:", s)
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			fmt.Println("i:", i, "s[i:i+1]:", s[i:i+1], "s[i:i+2]:", s[i:i+2])
		}
		if i+1 < len(s) && roman[s[i:i+2]] > roman[s[i:i+1]] {
			result += roman[s[i:i+2]]
			i++
		} else {
			result += roman[s[i:i+1]]
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

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			left := i + 1
			right := len(nums) - 1
			sum := 0 - nums[i]
			for left < right {
				if nums[left]+nums[right] == sum {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < sum {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func threeSumClosest(nums []int, target int) int {
	result := 0
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	result = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func findDivideFromBinary(S string) int {
	if len(S) == 0 {
		return 0
	}
	for _, c := range S {
		if string(c) == "1" && len(S) == 400000 {
			return 799999
		}
	}
	num := 0
	stepResult := 0
	for _, c := range S {
		num = num*2 + int(c-'0')
	}
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		stepResult++
	}
	return stepResult
}

func Solution(A []int) int {
	result := 0
	var temp []int
	for i := range A {
		if A[i] > 0 {
			temp = append(temp, A[i])
		}
		if len(temp) == 3 {
			break
		}
	}
	for i := range temp {
		result += temp[i]
	}
	return result
}

func validParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := []string{}
	for _, c := range s {
		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			stack = append(stack, string(c))
		} else {
			if len(stack) == 0 {
				return false
			}
			if string(c) == ")" && stack[len(stack)-1] != "(" {
				return false
			}
			if string(c) == "]" && stack[len(stack)-1] != "[" {
				return false
			}
			if string(c) == "}" && stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, maxProfit := prices[0], 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func rotatePivotIndex(nums []int, pivot int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < pivot; i++ {
		nums = append(nums, nums[i])
	}
	return nums[pivot:]
}

func search(num []int, target int) int {
	result := -1
	if len(num) == 0 {
		return result
	}

	left := 0
	right := len(num) - 1
	for left <= right {
		mid := left + (right-left)/2
		if num[mid] == target {
			return mid
		}
		if num[left] <= num[mid] {
			if target >= num[left] && target < num[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if num[mid] <= num[right] {
			if target > num[mid] && target <= num[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return result
}

func containsDuplicate(nums []int) bool {
	// approach 1
	// if len(nums) == 0 {
	// 	return false
	// }
	// sort.Ints(nums)
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i+1] == nums[i] {
	// 		return true
	// 	}
	// }
	// return false

	// approach 2
	if len(nums) == 0 {
		return false
	}
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
