package main

import "fmt"

//对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
//给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。

func gcdOfStrings(str1 string, str2 string) string {
	//引入Euclid
	euclid := func(num1 int, num2 int) int {
		mid := 0
		for num1%num2 != 0 {
			mid = num2
			num2 = num1 % num2
			num1 = mid
		}
		return num2
	}
	//判断是否有最大公因子
	l1, l2 := len(str1), len(str2)
	if l1 > l2 {
		for i, v := range str1 {
			if str2[i%l2] != uint8(v) {
				return ""
			}
		}
		if str1[:euclid(l1, l2)] != str2[l2-euclid(l1, l2):] {
			return ""
		}
		return str1[:euclid(l1, l2)]
	} else {
		for i, v := range str2 {
			if str1[i%l1] != uint8(v) {
				return ""
			}
		}
		if str2[:euclid(l2, l1)] != str1[l1-euclid(l2, l1):] {
			return ""
		}
		return str1[:euclid(l2, l1)]
	}
}

func main() {
	fmt.Println(gcdOfStrings("AABBAABBAA", "AABB"))
}
