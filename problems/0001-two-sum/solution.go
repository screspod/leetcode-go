package twosum

// Time: O(n) single pass through the array
// Memory: O(n) hash map storing up to n elements
// Pattern: store complement (target - num) as key so each lookup checks if the current number completes a pair
func twoSum(nums []int, target int) []int {
	complement := make(map[int]int)
	for i, num := range nums {
		index, ok := complement[num]
		if ok {
			return []int{index, i}
		}
		complement[target-num] = i
	}
	return []int{}
}
