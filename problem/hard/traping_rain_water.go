package hard

func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, trap := 0, 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			trap += leftMax - height[left]
			left++
		} else {
			trap += rightMax - height[right]
			right--
		}
	}
	return trap
}
