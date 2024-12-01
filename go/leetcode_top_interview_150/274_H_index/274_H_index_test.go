package leetcode_top_interview_150

import (
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		expected  int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{0, 1, 3, 5, 6}, 3},
		{[]int{10, 8, 5, 4, 3}, 4},
		{[]int{25, 8, 5, 3, 3}, 3},
	}

	for _, test := range tests {
		result := hIndex(test.citations)
		if result != test.expected {
			t.Errorf("For citations %v, expected %d, but got %d", test.citations, test.expected, result)
		}
	}
}
