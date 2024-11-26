package leetcode_top_interview_150

import "slices"

// 轮转数组

//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

func rotate(nums []int, k int) {
	t := k % len(nums)
	temp := make([]int, len(nums)-t)
	copy(temp, nums[0:len(nums)-t])
	for i := 0; i < len(nums); i++ {
		if i < t {
			nums[i] = nums[i+len(temp)]
		} else {
			// print(i, k, temp[i-k], " ")
			nums[i] = temp[i-t]
		}
	}
}

func rotate2(nums []int, k int) {
	k = k % len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
