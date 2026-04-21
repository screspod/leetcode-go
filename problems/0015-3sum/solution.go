package threesum

import (
	"slices"
)

// Time: O(n^2) — sorting is O(n log n), two-pointer scan is O(n) per element
// Memory: O(1) — no extra space beyond the output
// Pattern: Sort + Two Pointers — fix one element and use two pointers to find pairs summing to its negation; skip duplicates to avoid redundant triplets
func threeSum(nums []int) [][]int {
	result := [][]int{}

	slices.Sort(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Since the slice is sorted (ascending), we decrease right if we need
			// sum to be smaller and increase left if we need sum to be bigger
			switch {
			case sum > 0:
				right--
			case sum < 0:
				left++
			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				// Advance left if duplicate found
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Advance right if duplicate found
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

// Alternative that uses more memory
//
// Time: O(n²) outer loop O(n) × inner loop O(n) hash lookups
// Memory: O(n) hash set storing up to n elements per outer iteration
// Pattern: Sort + Hash Set — for each fixed pair (a,b), check if complement c exists in a set built from remaining elements
func _(nums []int) [][]int {
	result := [][]int{}

	slices.Sort(nums)

	for i := range len(nums) {
		// Skip duplicates for the fixed element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		seen := make(map[int]bool)

		for j := i + 1; j < len(nums); j++ {
			complement := -(nums[i] + nums[j])

			if seen[complement] {
				result = append(result, []int{nums[i], complement, nums[j]})

				// Skip duplicate nums[j] values to avoid duplicate triplets
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}

			seen[nums[j]] = true
		}
	}

	return result
}
