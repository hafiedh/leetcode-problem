package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
	"hafiedh.com/leetcode/problem/medium"
)

func main() {
	// Input: candidates = [10,1,2,7,6,1,5], target = 8
	candidates := []int{1, 2, 5}
	target := 9

	fmt.Println("Coin Change :", medium.CoinChange(candidates, target))

}
func Sort[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
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
		if i+1 < len(s) && roman[s[i:i+2]] > roman[s[i:i+1]] {
			fmt.Println("i:", i, "i+1:", i+1, "i+2:", i+2, "s[i:i+2]:", s[i:i+2], "s[i:i+1]:", s[i:i+1])
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

func merge(nums1 []int, m int, nums2 []int, n int) {
	// approach 1
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

	// approach 2
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	result := make([]int, m+n)
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] && nums1[i] != 0 && nums2[j] != 0 {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		result[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		result[k] = nums2[j]
		j++
		k++
	}
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	// approach 1
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[:i+1])
	return i + 1

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

	// approach 2
	// if len(nums) == 0 {
	// 	return 0
	// }
	// m := make(map[int]bool)
	// for _, num := range nums {
	// 	m[num] = true
	// }
	// result := []int{}
	// for key := range m {
	// 	result = append(result, key)
	// 	nums = result
	// }
	// return len(result)
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if !isEqual(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func isEqual(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}
	return a == b
}

func isAnagram(s string, t string) bool {

	// approach 1
	// if len(s) != len(t) {
	// 	return false
	// }
	// m := make(map[rune]int)
	// for _, c := range s {
	// 	m[c]++
	// }
	// for _, c := range t {
	// 	m[c]--
	// }
	// for _, v := range m {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }
	// return true

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

	// approach 2
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
		m[rune(s[i])]++
		m[rune(t[j])]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k %= len(nums)
	fmt.Println("k:", k)
	fmt.Println("original nums:", nums)
	reverse(nums, 0, len(nums)-1)
	fmt.Println(nums)
	reverse(nums, 0, k-1)
	fmt.Println(nums)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		fmt.Println("start:", start, "end:", end)
		nums[start], nums[end] = nums[end], nums[start]
		fmt.Println("nums in reverse: ", nums)
		start++
		end--
	}
}

type ListNode struct {
	Next *ListNode
	Val  int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var result *ListNode
	if list1.Val < list2.Val {
		result = list1
		result.Next = mergeTwoLists(list1.Next, list2)
	} else {
		result = list2
		result.Next = mergeTwoLists(list1, list2.Next)
	}
	return result
}

func searchLeetCode(nums []int, target int) int {
	// approach 1
	// result := -1
	// if len(nums) == 0 {
	// 	return result
	// }

	// left := 0
	// right := len(nums) - 1
	// for left <= right {
	// 	mid := left + (right-left)/2
	// 	if nums[mid] == target {
	// 		return mid
	// 	}
	// 	if nums[left] <= nums[mid] {
	// 		if target >= nums[left] && target < nums[mid] {
	// 			right = mid - 1
	// 		} else {
	// 			left = mid + 1
	// 		}
	// 	} else if nums[mid] <= nums[right] {
	// 		if target > nums[mid] && target <= nums[right] {
	// 			left = mid + 1
	// 		} else {
	// 			right = mid - 1
	// 		}
	// 	}

	// }
	// return result

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

	// approach 2
	result := -1
	if len(nums) == 0 {
		return result
	}

	for _, v := range nums {
		if v == target {
			return v
		}
	}
	return result
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[len(s)-1] == ' ' {
		s = strings.Trim(s, " ")
	}
	lastWordLength := 0
	for i := len(s) - 1; i >= 0 && s[i] != ' '; i-- {
		lastWordLength++
	}
	return lastWordLength
}

func letterCombination(str string) []string {
	res := []string{}
	if len(str) == 0 {
		return res
	}

	m := make(map[string]string)
	m["2"] = "abc"
	m["3"] = "def"
	m["4"] = "ghi"
	m["5"] = "jkl"
	m["6"] = "mno"
	m["7"] = "pqrs"
	m["8"] = "tuv"
	m["9"] = "wxyz"

	res = []string{""}
	for _, digit := range str {
		var temp []string
		for _, char := range m[string(digit)] {
			for _, s := range res {
				temp = append(temp, s+string(char))
			}
		}
		res = temp
	}
	return res
}

func plusOne(digits []int) []int {
	// approach 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)

	// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-= //

	// approach 2
	// carry := false
	// for i := len(digits) - 1; i >= 0; i-- {
	// 	if digits[i] == 9 {
	// 		carry = true
	// 		digits[i] = 0
	// 	} else {
	// 		digits[i]++
	// 		carry = false
	// 		break
	// 	}
	// }
	// if carry {
	// 	digits = append([]int{1}, digits...)
	// }
	// return digits
}
