package main

import "fmt"

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。

//func moveZeroes(nums []int) {
//	moveByIndex := func(index int) {
//		num := nums[index]
//		for i := index; i < len(nums)-1; i++ {
//			nums[i] = nums[i+1]
//		}
//		nums[len(nums)-1] = num
//	}
//	//发现0
//	zeroCount := 0
//	for i := 0; i < len(nums)-zeroCount; {
//		if nums[i] == 0 {
//			zeroCount++
//			moveByIndex(i)
//			continue
//		}
//		i++
//	}
//	fmt.Println(nums)
//}

//上面的方法空间复杂度为O(1)，但是时间复杂度很高，可以采用双指针法

func moveZeroes(nums []int) {
	//定义左右指针
	left, right := 0, 0
	//右指针遍历数组
	for right < len(nums) {
		//遍历到非零元素时赋值给左指针
		if nums[right] != 0 {
			//避免自己给自己赋值
			if left == right {
				left++
				right++
				continue
			}
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
			continue
		}
		right++
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{2, 1})
}
