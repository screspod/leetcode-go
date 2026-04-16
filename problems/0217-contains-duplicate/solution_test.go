package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1 - has duplicates",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "example 2 - no duplicated",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "example 3 - multiple duplicates",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "no elements",
			nums: []int{},
			want: false,
		},
		{
			name: "two identical elements",
			nums: []int{2, 2},
			want: true,
		},
		{
			name: "duplicates at the end",
			nums: []int{1, 2, 3, 4, 5, 5},
			want: true,
		},
		{
			name: "negative duplicates",
			nums: []int{1, 2, -3, 4, -3, 5},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.want {
				t.Logf("name: %v", tt.name)
				t.Errorf("containsDuplicate(%v) = %v, want = %v", tt.nums, result, tt.want)
			}
		})
	}
}
