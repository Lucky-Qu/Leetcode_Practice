package main

import (
	"fmt"
)

//泰波那契序列 Tn 定义如下：
//
//T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
//
//给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

//	func tribonacci(n int) int {
//		switch n {
//		case 0:
//			return 0
//		case 1:
//			return 1
//		case 2:
//			return 1
//		}
//		return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
//	}
//
// 上面这段方法超时了，所以在想怎么加速
// 想到这个函数最终都会变为角标0，1，2的相加
// F3 = F2 + F1 + F0
// F4 = F3 + F2 + F1 = F2 + F1 + F0 + F2 + F1 = 2F2 + 2F1 + F0
// F5 = F4 + F3 + F2 = F3 + F2 + F1 + F2 + F1 + F0 + F2 = F2 + F1 + F0 + F2 + F1 + F2 + F1 + F0 + F2 = 4F2 + 3F1 + F0
// F6 = 7F2 + 6F1 + 3F0
// 因为F0 = 0，所以统计F1和F2的数量后相加即是答案
// 思路错了，这样和上面的思路是基本一样的
// 新思路
// 因为是三个数的和，所以从0，1，2开始记录值，递增直到等于n
func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	x, y, z := 0, 1, 1
	for i := 3; i <= n; i++ {
		x, y, z = y, z, x+y+z
	}
	return z
}

func main() {
	fmt.Println(tribonacci(37))
}
