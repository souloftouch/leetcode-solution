package leetcode_top_interview_150

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{1, 1, 1, 1}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 6, 5},
		{[]int{}, 1, 0},
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		result := removeElement(numsCopy, test.val)
		if result != test.expected {
			t.Errorf("removeElement(%v, %d) = %d; expected %d", test.nums, test.val, result, test.expected)
		}
	}
}
