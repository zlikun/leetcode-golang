package main

import (
	"fmt"
)

func main() {
	var r []int
	var nums []int

	nums = []int{3, 2, 4}
	r = twoSum1(nums, 6)
	fmt.Println(nums[r[0]]+nums[r[1]] == 6)

	nums = []int{3, 3}
	r = twoSum2(nums, 6)
	fmt.Println(nums[r[0]]+nums[r[1]] == 6)

	nums = []int{3, 4, 3}
	r = twoSum3(nums, 6)
	fmt.Println(nums[r[0]]+nums[r[1]] == 6)
}

// 解法一：两次循环，元素之和相等则返回，
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	for i, m := range nums {
		for j, n := range nums[i+1:] {
			if m+n == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return make([]int, 2)
}

// 两遍哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum2(nums []int, target int) []int {
	// 每种输入只取一个答案，所以可以将值与索引反转，使用map映射
	var data = make(map[int]int)
	// 注意转换为map时，只记录第一个元素索引（过滤重复值）
	for i, v := range nums {
		if _, ok := data[v]; !ok {
			data[v] = i
		}
	}
	for i, m := range nums {
		// i != j 用于处理 nums 包含相同元素时用
		if j, ok := data[target-m]; ok && i != j {
			return []int{i, j}
		}
	}
	return make([]int, 2)
}

// 一遍哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum3(nums []int, target int) []int {
	// 每种输入只取一个答案，所以可以将值与索引反转，使用map映射
	var data = make(map[int]int)
	// 注意转换为map时，只记录第一个元素索引（过滤重复值）
	for i, v := range nums {
		// 实际上可以用当前值来计算已有哈希表（map）中是否有匹配项，从而一轮迭代完成
		if j, ok := data[target-v]; ok {
			return []int{i, j}
		}
		if _, ok := data[v]; !ok {
			data[v] = i
		}
	}
	return make([]int, 2)
}
