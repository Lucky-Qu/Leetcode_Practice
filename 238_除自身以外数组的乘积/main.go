package main

import "fmt"

//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

//func productExceptSelf(nums []int) []int {
//	result := make([]int, len(nums), len(nums))
//	mid := 1
//	for i, _ := range nums {
//		for j, w := range nums {
//			if i == j {
//				continue
//			}
//			mid *= w
//		}
//		result[i] = mid
//		mid = 1
//	}
//	return result
//}

// 上面是第一种个人思路，可以解决问题但是时间复杂度在O(N2)，优化一下
// 个人思路历程：
// 既然要求时间复杂度不高于ON，并且考虑到要处理数组中的全部元素必定会遍历一次，遍历循环的复杂度为O(N)，说明不能嵌套循环
// 那么第一次正着遍历，第二次反着遍历，将每一项两次遍历的结果相乘，会得到正确结果多乘自己这个数的平方的结果
// 那么进一步思考，将两次遍历错开两位相乘即可得到正确答案（如思路图）
//
//func productExceptSelf(nums []int) []int {
//	result := make([]int, len(nums)+2, len(nums)+2)
//	mid1 := make([]int, len(nums), len(nums))
//	mid2 := make([]int, len(nums), len(nums))
//	mid3 := 1
//	for i := 0; i < len(nums); i++ {
//		mid3 *= nums[i]
//		mid1[i] = mid3
//	}
//	mid3 = 1
//	for i := len(nums) - 1; i >= 0; i-- {
//		mid3 *= nums[i]
//		mid2[i] = mid3
//	}
//	mid1 = append([]int{1, 1}, mid1...)
//	mid2 = append(mid2, 1, 1)
//	for i := range result {
//		result[i] = mid1[i] * mid2[i]
//	}
//	fmt.Println(mid1, mid2)
//	return result[1 : len(result)-1]
//}
//这种思路是对的，并且也符合了题目要求，但是时间空间仍是不理想的，在上面的实现中使用了两个append操作，产生了多余的空间浪费
//那么将思路抽离出核心，关键思路就是将一个数的前缀积和后缀积相乘，优化代码后如下：

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	mid := make([]int, len(nums), len(nums))
	//初始化
	result[0] = 1
	mid[len(nums)-1] = 1
	//前缀积
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]
	}
	//后缀积相乘
	for i := len(nums) - 1 - 1; i >= 0; i-- {
		mid[i] = mid[i+1] * nums[i+1]
		result[i] *= mid[i]
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
