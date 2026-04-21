package threesum

import (
	"slices"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(threeSum(tt.nums))
			want := normalize(tt.want)
			if !slices.EqualFunc(got, want, func(a, b []int) bool {
				return slices.Equal(a, b)
			}) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

func normalize(triplets [][]int) [][]int {
	for _, t := range triplets {
		sort.Ints(t)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := range 3 {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
	return triplets
}
