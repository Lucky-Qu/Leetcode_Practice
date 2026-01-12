package main

import (
	"fmt"
)

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
// 示例 1:
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
//
// 输入：height = [1,1]
// 输出：1

//	func maxArea(height []int) int {
//		// 个人思路：
//		// 容纳水的面积其实就是长方形的面积：较短的高 * 宽
//		// 那一开始可以假定长方形的宽是第一个和最后一个，高是他俩之间较低的哪一个
//		// 然后慢慢从外面往里走，每次都计算面积，遇到大的就替换
//		// 直到两个重叠停止，只需要一次遍历和对应的计算面积
//		// 因为有短板效应，所以挪动较短的那个角标优先
//		left, right := 0, len(height)-1
//		getArea := func(left, right int) int {
//			return min(height[left], height[right]) * (right - left)
//		}
//		area := getArea(left, right)
//		for left < right {
//			if height[left] < height[right] {
//				// 挪动左
//				i := 1
//				for {
//					if left+i >= right {
//						return area
//					}
//					if getArea(left+i, right) > area {
//						left += i
//						area = getArea(left, right)
//						break
//					}
//					// 如果当前移动对应的高度比右边长，转为挪动右
//					if height[left+i] > height[right] {
//						left += i
//						break
//					}
//					i++
//				}
//			} else {
//				// 挪动右
//				i := 1
//				for {
//					if right-i <= left {
//						return area
//					}
//					if getArea(left, right-i) > area {
//						right -= i
//						area = getArea(left, right)
//						break
//					}
//					// 如果当前移动对应的高度比左边长，转为挪动左
//					if height[right-i] > height[left] {
//						right -= i
//						break
//					}
//					i++
//				}
//			}
//		}
//		return area
//	}
//
// 注：上面有点弄巧成拙了，可以通过，但是写的有点丑陋，还是下面的吧
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	getArea := func(left, right int) int {
		return min(height[left], height[right]) * (right - left)
	}
	area := getArea(left, right)
	for {
		// 谁小移动谁
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		// 边界检查
		if left >= right {
			break
		}
		if area < getArea(left, right) {
			area = getArea(left, right)
		}
	}
	return area
}

func main() {
	testArray := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(testArray))
}
