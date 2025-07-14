package main

import "fmt"

//有 n 个有糖果的孩子。给你一个数组 candies，其中 candies[i] 代表第 i 个孩子拥有的糖果数目，和一个整数 extraCandies 表示你所有的额外糖果的数量。
//
//返回一个长度为 n 的布尔数组 result，如果把所有的 extraCandies 给第 i 个孩子之后，他会拥有所有孩子中 最多 的糖果，那么 result[i] 为 true，否则为 false。
//
//注意，允许有多个孩子同时拥有 最多 的糖果数目。

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := 0
	for _, v := range candies {
		if maxNum < v {
			maxNum = v
		}
	}
	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= maxNum {
			result[i] = true
		}
	}
	return result
}

//和官方题解思路一致

func main() {
	fmt.Println(kidsWithCandies([]int{1, 2, 3}, 2))
}
