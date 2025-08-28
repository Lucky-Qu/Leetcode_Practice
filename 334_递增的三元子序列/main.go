package main

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

// 个人思路，定义三个数i，j，k，判断数组三个数是否递增，递增返回true，不递增全部加加
//func increasingTriplet(nums []int) bool {
//	i, j, k := 0, 1, 2
//	for k < len(nums) {
//		if nums[i] < nums[j] && nums[j] < nums[k] {
//			return true
//		}
//		i++
//		j++
//		k++
//	}
//	return false
//}

// 思路错了，他没有要求相邻
// 新思路是穷举
// TODO
//func increasingTriplet(nums []int) bool {
//	for i, v := range nums {
//
//	}
////}
//
//func main() {
//	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
//}
