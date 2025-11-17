package main

import "fmt"

// 给你一个整数数组 arr，如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	check := make(map[int]bool)
	for _, v := range arr {
		m[v]++
	}
	for _, v := range m {
		if !check[v] {
			check[v] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	if uniqueOccurrences([]int{1, 2, 4, 5, 6}) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
