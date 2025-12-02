package solutions

import "testing"

func Test_repeatedTwice(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
		// TODO: Add test cases.
		{"11", 11, true},
		{"12", 12, false},
		{"121", 121, false},
		{"112233112233", 112233112233, true},
		{"1234512345", 1234512345, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatedTwice(tt.n)
			if got != tt.want {
				t.Errorf("repeatedTwice() = %v, want %v", got, tt.want)
			}
		})
	}
}
