package containsduplicate

// Time: O(n) single pass through the array
// Memory: O(n) hash set storing up to n elements
// Pattern: hash set gives O(1) existence checks, turning a naive O(n²) scan into a single pass
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
