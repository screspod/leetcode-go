package topkfrequentelements

import (
	"slices"
)

// Time: O(n) three linear passes (frequency map, bucket fill, bucket scan)
// Nested loops are still O(n) total because each element is touched once across all buckets
//
// Memory: O(n) frequency map and buckets slice both bounded by input size
// The 2D buckets slice holds at most n elements total across all inner slices
//
// Pattern: Hash Map + Bucket Sort — count frequencies with a map, use frequency as bucket index (range [1,n]) to rank without sorting, scan buckets in reverse for top-k
func topKFrequent(nums []int, k int) []int {
	// Maps each number to a counter of frequency
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// Sort each number into buckets where the index of the bucket represents the frequency
	// Example with [1,1,1,2,2,3]:
	// freq map: {1→3, 2→2, 3→1}
	//
	// buckets:
	// index 0: []
	// index 1: [3]      ← element 3 appears 1 time
	// index 2: [2]      ← element 2 appears 2 times
	// index 3: [1]      ← element 1 appears 3 times
	buckets := make([][]int, len(nums)+1)
	for num, freq := range frequency {
		buckets[freq] = append(buckets[freq], num)
	}

	// Iterate through all buckets in reverse to get top k numbers
	result := make([]int, 0, k)
	for _, bucket := range slices.Backward(buckets) {
		for _, num := range bucket {
			result = append(result, num)
			if len(result) >= k {
				return result
			}
		}
	}

	return result
}
