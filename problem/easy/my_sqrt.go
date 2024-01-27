package problem

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2 // 1 + (8-1)/2 = 4
		if mid*mid == x {            // 4*4 = 16
			return mid // 4
		} else if mid*mid < x { // 4*4 < 16
			left = mid + 1 // 4 + 1 = 5
		} else {
			right = mid - 1 // 4 - 1 = 3
		}
	}
	return right
}
