package main

import "fmt"

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
// 任何误差小于 10-5 的答案都将被视为正确答案。
func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := left + k
	maxSumValue := func() int {
		sum := 0
		for i := left; i < right; i++ {
			sum += nums[i]
		}
		return sum
	}()
	sumValue := maxSumValue
	for right < len(nums) {
		sumValue -= nums[left]
		sumValue += nums[right]
		if sumValue > maxSumValue {
			maxSumValue = sumValue
		}
		left++
		right++
	}
	return float64(maxSumValue) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 1, 34, 1, 1, 1}, 3))
}
