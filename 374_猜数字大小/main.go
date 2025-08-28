package main

//我们正在玩猜数字游戏。猜数字游戏的规则如下：
//
// 我会从 1 到 n 随机选择一个数字。 请你猜选出的是哪个数字。（我选的数字在整个游戏中保持不变）。
//
// 如果你猜错了，我会告诉你，我选出的数字比你猜测的数字大了还是小了。
//
// 你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有三种可能的情况：
//
// -1：你猜的数字比我选出的数字大 （即 num > pick）。
// 1：你猜的数字比我选出的数字小 （即 num < pick）。
// 0：你猜的数字与我选出的数字相等。（即 num == pick）。
// 返回我选出的数字。

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2 // 取中间，避免溢出
		res := guess(mid)
		if res == 0 {
			return mid // 猜对了
		} else if res < 0 { // mid 比目标大
			high = mid - 1
		} else { // mid 比目标小
			low = mid + 1
		}
	}
	return -1 // 理论上不会到这里
}
func main() {

}
