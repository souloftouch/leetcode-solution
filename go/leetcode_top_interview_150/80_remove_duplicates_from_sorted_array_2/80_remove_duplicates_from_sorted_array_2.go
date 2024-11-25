package leetcode_top_interview_150

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates2(nums []int) int {
	for i, j := 2, len(nums)-1; i < len(nums); i++ {
		if nums[i] == nums[i-2] {
			for k := i; k < j; k++ {
				nums[k] = nums[k+1]
			}
			j--
			i--
		}
		if i >= j {
			return i + 1
		}
	}
	return len(nums)
}
