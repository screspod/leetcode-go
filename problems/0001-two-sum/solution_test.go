package twosum

import (
	"slices"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1 - has solution",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2 - no solution",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3 - duplicate values",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "target at end",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			want:   []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			sort.Ints(result)
			sort.Ints(tt.want)
			if !slices.Equal(result, tt.want) {
				t.Errorf("twoSum(%v, %v) = %v, want = %v", tt.nums, tt.target, result, tt.want)
			}
		})
	}
}
