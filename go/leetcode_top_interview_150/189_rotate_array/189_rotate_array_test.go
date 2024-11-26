package leetcode_top_interview_150

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{3, 4, 5, 6, 1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	for _, test := range tests {
		rotate(test.nums, test.k)
		if !reflect.DeepEqual(test.nums, test.expected) {
			t.Errorf("rotate(%v, %d) = %v; expected %v", test.nums, test.k, test.nums, test.expected)
		}
	}
}
func TestRotate2(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{3, 4, 5, 6, 1, 2}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	for _, test := range tests {
		rotate2(test.nums, test.k)
		if !reflect.DeepEqual(test.nums, test.expected) {
			t.Errorf("rotate2(%v, %d) = %v; expected %v", test.nums, test.k, test.nums, test.expected)
		}
	}
}
