package main

import "fmt"

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 个人思路：参考实现的整体削平，但是好像难点在于使用O(N)实现，我实现的最坏情况复杂度是O(N2)

//func trap(height []int) int {
//	left, right := 0, len(height)-1
//	depth := 0
//	sum := 0
//	for left < right {
//		// 找到左边支柱
//		for height[left] == 0 && left < right {
//			left++
//		}
//		// 找到右边支柱
//		for height[right] == 0 && left < right {
//			right--
//		}
//		// 确定深度
//		depth = min(height[left], height[right])
//		// 计算并记录
//		for i := left; i <= right; i++ {
//			// 比深度高则削高，比深度低则记录
//			if height[i] >= depth {
//				height[i] -= depth
//			} else {
//				sum += depth - height[i]
//				height[i] = 0
//			}
//		}
//	}
//	return sum
//}

func trap(height []int) int {
	// 新思路是左右遍历一遍
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	sum := 0
	for left < right {
		// 短板移动
		if leftMax <= rightMax {
			left++
			// 比较最值
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum += rightMax - height[right]
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
