package leetcode_top_interview_150

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := removeDuplicates(test.nums)
		if result != test.expected {
			t.Errorf("removeDuplicates(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
