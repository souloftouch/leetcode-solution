package leetcode_top_interview_150

import (
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		{[]int{1, 1, 2, 2, 2, 3, 3}, 6},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 2},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := removeDuplicates2(test.nums)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
