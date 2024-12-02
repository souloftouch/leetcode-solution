package leetcode_top_interview_150

import (
	"math/rand"
)

// O(1)时间插入、删除和获取随机元素

// 实现RandomizedSet 类：

// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

type RandomizedSet struct {
	nums    []int
	numsMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 0),
		numsMap: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.numsMap[val]; ok {
		return false
	}
	rs.nums = append(rs.nums, val)
	rs.numsMap[val] = len(rs.nums) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.numsMap[val]; !ok {
		return false
	}
	index := rs.numsMap[val]
	rs.numsMap[rs.nums[len(rs.nums)-1]] = index
	rs.nums[index] = rs.nums[len(rs.nums)-1]
	rs.nums = rs.nums[:len(rs.nums)-1]
	delete(rs.numsMap, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
