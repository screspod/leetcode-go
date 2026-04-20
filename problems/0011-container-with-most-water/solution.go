package containerwithmostwater

// Time: O(n) single pass with two pointers converging from both ends
// Memory: O(1) only a constant number of variables
// Pattern: Two pointers; always move the pointer at the shorter line inward to maximize potential area
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	best := 0

	for left < right {
		best = max(best, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return best
}
