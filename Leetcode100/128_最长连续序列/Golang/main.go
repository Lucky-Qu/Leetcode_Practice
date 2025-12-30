package main

import "fmt"

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//示例 1：
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//示例 3：
//输入：nums = [1,0,1,2]
//输出：3

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	maxNum := 0
	// 去重
	for _, v := range nums {
		m[v] = true
	}
	for v := range m {
		if m[v] && !m[v-1] {
			i := 0
			for j := v; m[j]; j++ {
				i++
				if i > maxNum {
					maxNum = i
				}
			}
		}
	}
	return maxNum
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 2, 1, 4, 9, 4, 5, 3}))
}
