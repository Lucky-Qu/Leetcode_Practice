package main

// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。

// 暴力思路：遍历出每一种走法，比较花费，返回最少的花费
func minCostClimbingStairs(cost []int) int {

	return 0
}

func main() {
	//如果只看后两位，1 + 100 + 1000 + 1
	//实际应当是，1 + 1000 + 1
	minCostClimbingStairs([]int{1, 1, 100, 1000, 10000, 1})
}
