package main

import "fmt"

//给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：
//answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
//answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
//注意：列表中的整数可以按 任意 顺序返回。

func findDifference(nums1 []int, nums2 []int) [][]int {
	result1 := make(map[int]bool)
	result2 := make(map[int]bool)
	checkNum1 := make(map[int]bool)
	checkNum2 := make(map[int]bool)
	for _, v := range nums1 {
		checkNum1[v] = true
	}
	for _, v := range nums2 {
		checkNum2[v] = true
	}
	for _, v := range nums2 {
		if !checkNum1[v] {
			result1[v] = true
		}
	}
	for _, v := range nums1 {
		if !checkNum2[v] {
			result2[v] = true
		}
	}
	return func() [][]int {
		r1, r2 := make([]int, 0), make([]int, 0)
		for k := range result1 {
			r1 = append(r1, k)
		}
		for k := range result2 {
			r2 = append(r2, k)
		}
		return [][]int{r2, r1}
	}()
}

func main() {
	ans := findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
	fmt.Println(ans)
}
