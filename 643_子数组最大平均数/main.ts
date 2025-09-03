// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
// 任何误差小于 10-5 的答案都将被视为正确答案。

function findMaxAverage(nums: number[], k: number): number {
    let windowsSum: number = 0
    let maxSum: number = 0
    let left: number = 0
    let right: number = left + k
    for (let i = left; i < right; i++) {
        windowsSum += nums[i]
    }
    maxSum = windowsSum
    for (;right < nums.length;){
        windowsSum -= nums[left]
        windowsSum += nums[right]
        if (windowsSum > maxSum){
            maxSum = windowsSum
        }
        left++
        right++
    }
    return maxSum / k
}
