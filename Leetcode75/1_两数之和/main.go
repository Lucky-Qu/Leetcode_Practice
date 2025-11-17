package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		} else {
			m[v] = i
		}
	}
	return nil
}

func main() {
	var a = []int{5, 4, 3, 2, 1}
	fmt.Println(twoSum(a, 5))
}
