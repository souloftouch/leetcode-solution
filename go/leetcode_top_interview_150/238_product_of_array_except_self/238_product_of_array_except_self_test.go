package leetcode_top_interview_150

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
	}

	for _, test := range tests {
		result := productExceptSelf(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums %v, expected %v, but got %v", test.nums, test.expected, result)
		}
	}
}
