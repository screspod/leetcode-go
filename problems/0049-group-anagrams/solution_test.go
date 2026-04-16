package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "example 2 - empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "example 3 - single char",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			if !equalGroups(result, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, result, tt.want)
			}
		})
	}
}

// equalGroups compares two [][]string regardless of group order or element order within groups.
func equalGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	normalize := func(groups [][]string) [][]string {
		result := make([][]string, len(groups))
		for i, g := range groups {
			sorted := make([]string, len(g))
			copy(sorted, g)
			sort.Strings(sorted)
			result[i] = sorted
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0]
		})
		return result
	}
	return reflect.DeepEqual(normalize(a), normalize(b))
}
