package twosum

// Time: O(n) single pass through the array
// Memory: O(n) hash map storing up to n elements
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		index, ok := seen[num]
		if ok {
			return []int{index, i}
		}
		seen[target-num] = i
	}
	return []int{}
}
