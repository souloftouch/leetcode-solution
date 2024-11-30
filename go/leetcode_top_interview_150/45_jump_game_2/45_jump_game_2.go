package leetcode_top_interview_150

// 跳跃游戏 Ⅱ

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]

// i + j < n

// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

func jump(nums []int) int {
	n := len(nums)
	count, current, next := 0, 0, 0
	for i := 0; i < n-1; i++ {
		if i+nums[i] > next {
			next = i + nums[i]
		}
		if i == current {
			count++
			current = next
		}
	}
	return count
}
