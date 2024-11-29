package leetcode_top_interview_150

// 跳跃游戏

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

func canJump(nums []int) bool {
	n := len(nums)
	last := n - 1
	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last == 0
}
