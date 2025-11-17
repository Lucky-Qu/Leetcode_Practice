// 可以遍历中往map里写数据，占用会少点
package main

import "fmt"

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
//你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	// 将数组内数字存到map
	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}
	// 遍历，检查是否有符合的
	for i, v := range nums {
		if j, ok := numsMap[target-v]; ok {
			if j == i {
				continue
			}
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
