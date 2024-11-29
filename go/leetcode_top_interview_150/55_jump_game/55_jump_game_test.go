package leetcode_top_interview_150

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{2, 0, 0}, true},
		{[]int{1, 1, 1, 1, 1}, true},
		{[]int{1, 2, 3}, true},
		{[]int{1, 0, 1, 0}, false},
	}

	for _, test := range tests {
		result := canJump(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %v, but got %v", test.nums, test.expected, result)
		}
	}
}
