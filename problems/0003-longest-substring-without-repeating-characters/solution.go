package longestsubstringwithoutrepeatingcharacters

// Time: O(n) single pass through the string
// Memory: O(min(n, m)) hash map stores at most m unique characters (m = charset size)
// Pattern: Sliding window with a hash map to track last seen index; shrink window left boundary on duplicate
func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	left := 0
	best := 0

	for right := range len(s) {
		if i, ok := seen[s[right]]; ok {
			left = max(left, i+1)
		}
		seen[s[right]] = right
		best = max(best, right-left+1)
	}

	return best
}
