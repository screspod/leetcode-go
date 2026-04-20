package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "abcabcbb", want: 3},
		{name: "example 2", s: "bbbbb", want: 1},
		{name: "example 3", s: "pwwkew", want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
