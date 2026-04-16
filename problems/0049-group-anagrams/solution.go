package groupanagrams

import (
	"slices"
)

// Time: O(n * k log k) sorting each string of length k across n strings
// Memory: O(n * k) map storing all strings
// Pattern: sorting each string produces a canonical key — anagrams share the same sorted form
func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		key := string(b)
		group[key] = append(group[key], str)
	}
	result := [][]string{}
	for _, v := range group {
		result = append(result, v)
	}
	return result
}
