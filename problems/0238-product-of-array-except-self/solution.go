package productofarrayexceptself

// Time: O(n) two passes over nums
// Memory: O(1) only two scalar variables, excluding output array
// Pattern: Prefix/suffix products — fill answer with left products in a forward pass, then multiply in right products in-place with a backward pass
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	left := 1
	for i := range nums {
		answer[i] = left
		left = left * nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right = right * nums[i]
	}

	return answer
}
