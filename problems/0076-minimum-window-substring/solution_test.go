package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "example 3",
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
