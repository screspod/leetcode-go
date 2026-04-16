package containsduplicate

// containsDuplicate returns true if the array contains any duplicate values.
// Uses a hash set to track seen values with O(n) time and O(n) space complexity.
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true

		}
		seen[num] = struct{}{}
	}
	return false
}
