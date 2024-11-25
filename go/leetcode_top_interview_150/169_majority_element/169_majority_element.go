package leetcode_top_interview_150

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func majorityElement(nums []int) int {
	ret := 0
	countMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
		if countMap[nums[i]] > countMap[ret] {
			ret = nums[i]
		}
	}
	return ret
}

func majorityElement2(nums []int) int {
	candidate, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
