package leetcode_top_interview_150

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{5, 6, 1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 6},
	}

	for _, test := range tests {
		result := jump(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
