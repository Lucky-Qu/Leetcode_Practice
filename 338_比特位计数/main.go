package main

import "fmt"

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

func countBits(n int) []int {
	result := make([]int, n+1)
	howManyOneInIt := func(num int) int {
		count := 0
		for num != 0 {
			if num%2 != 0 {
				count++
			}
			num /= 2
		}
		return count
	}
	for i := 0; i <= n; i++ {
		result[i] = howManyOneInIt(i)
	}
	return result
}

func main() {
	fmt.Println(countBits(8))
}
