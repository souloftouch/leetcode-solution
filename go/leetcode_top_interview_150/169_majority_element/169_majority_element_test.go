package leetcode_top_interview_150

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{1, 1, 2, 2, 2}, 2},
	}

	for _, test := range tests {
		result := majorityElement(test.nums)
		if result != test.expected {
			t.Errorf("majorityElement(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}

func TestMajorityElement2(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{1, 1, 2, 2, 2}, 2},
	}

	for _, test := range tests {
		result := majorityElement2(test.nums)
		if result != test.expected {
			t.Errorf("majorityElement2(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
