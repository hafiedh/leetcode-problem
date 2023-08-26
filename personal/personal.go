package personal

import "fmt"

// i have array [1,2,3] then print 123,
// i have array [8,9] then print 90,
// i have array [9,9] then print 100,
// i have array [9,9,9] then print 1000,
// i have array [1,0,0,0] then print 1000,

func PlusOne(digits []int) []int {
	// approach 1
	// for i := len(digits) - 1; i >= 0; i-- {
	// 	if digits[i] < 9 {
	// 		digits[i]++
	// 		return digits
	// 	}
	// 	digits[i] = 0
	// }
	// return append([]int{1}, digits...)

	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			carry = true
			digits[i] = 0
			fmt.Println(digits)
		} else {
			digits[i]++
			carry = false
			fmt.Println(digits)
			break
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
