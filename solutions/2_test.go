package solutions

import "testing"

func Test_repeatedTwice(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
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

func Test_isSillyPattern(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
		{"111", 111, true},
		{"112112", 112112, true},
		{"2121212121", 2121212121, true},
		{"824824824", 824824824, true},
		{"565656", 565656, true},
		{"999", 999, true},
		{"12312", 12312, false},
		{"123", 123, false},
		{"123321", 123321, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSillyPattern(tt.n)
			if got != tt.want {
				t.Errorf("isSillyPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsSillyPatternN(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		p    string
		n    int
		want bool
	}{
		{"111", "111", "1", 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsSillyPatternN(tt.str, tt.p, tt.n)
			if got != tt.want {
				t.Errorf("containsSillyPatternN() = %v, want %v", got, tt.want)
			}
		})
	}
}
