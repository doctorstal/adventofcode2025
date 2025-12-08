package solutions

import (
	"testing"

	"gotest.tools/assert"
)

func Test_markTimelines(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		manifold [][]int
		wanted   []int
	}{
		{
			"simple",
			[][]int{
				{0, 1, 0},
				{0, -1, 0},
				{0, 0, 0},
			},
			[]int{1, 0, 1},
		},
		{
			"simple1.5",
			[][]int{
				{0, 0, 1, 0},
				{0, 0, -1, 0},
				{0, -1, 0, 0},
				{0, 0, 0, 0},
			},
			[]int{1, 0, 1, 1},
		},
		{
			"simple2",
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, -1, 0, 0},
				{0, -1, 0, -1, 0},
				{0, 0, 0, 0, 0},
			},
			[]int{1, 0, 2, 0, 1},
		},
		{
			"simple3",
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, -1, 0, 0},
				{0, -1, 0, -1, 0},
				{0, 0, -1, 0, 0},
				{0, 0, 0, 0, 0},
			},
			[]int{1, 2, 0, 2, 1},
		},
		{
			"simple4",
			[][]int{
				{0, 0, 0, 1, 0, 0},
				{0, 0, 0, -1, 0, 0},
				{0, 0, -1, 0, 0, 0},
				{0, -1, 0, 0, -1, 0},
				{0, 0, 0, 0, 0, 0},
			},
			[]int{1, 0, 1, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markTimelines(tt.manifold)
			assert.DeepEqual(t, tt.manifold[len(tt.manifold)-1], tt.wanted)
		})
	}
}
