package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
//满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
//请你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
//示例 1：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//示例 2：
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//示例 3：
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		left, right := i+1, n-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				lv, rv := nums[left], nums[right]
				for left < right && nums[left] == lv {
					left++
				}
				for left < right && nums[right] == rv {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

//func threeSum(nums []int) [][]int {
// 初步个人思路：
// 求两数之和为0只需要一个map来判定即可
// 将三个数变为两层两数之和，用Map来解决
// 第一个Map存放一个容量为2的数组，Key为每两个数之和，值为两个数的索引
// 然后遍历数组，如果和为0就检查，然后返回角标对应的值
// 为什么不直接存两个数的值？防止遍历时用同一个索引的值

// 修正个人思路：
// 按照上面的实现的话有一个问题是相同和的数组数可能互相覆盖，最后只留下一组
// 所以应该用数组作为Key，和为Value，同时借助第二个Key为值Value为索引的Map来比对

// 二次修正：
// 这样还是无法规避覆盖问题，决定使用map存储答案来去重

//m := make(map[[2]int]int)
//a := make(map[int]int)
//result := make(map[[3]int]bool)
//for i := 0; i < len(nums); i++ {
//	a[nums[i]] = i
//	for j := i + 1; j < len(nums); j++ {
//		m[[2]int{i, j}] = nums[i] + nums[j]
//	}
//}
//for k, v := range m {
//	if index, ok := a[0-v]; ok {
//		// 索引唯一
//		if index != k[0] && index != k[1] {
//			result[[3]int{min(-v, nums[k[0]], nums[k[1]]), 0 - min(-v, nums[k[0]], nums[k[1]]) - max(-v, nums[k[0]], nums[k[1]]), max(-v, nums[k[0]], nums[k[1]])}] = true
//		}
//	}
//}
//resultArr := make([][]int, 0)
//for k := range result {
//	resultArr = append(resultArr, k[:])
//}
//return resultArr

// 思路肯定错了，推倒重来
//}

func main() {
	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}))
}
