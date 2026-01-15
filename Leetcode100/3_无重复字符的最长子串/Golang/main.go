package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	func lengthOfLongestSubstring(s string) int {
//		left := 0
//		curLen := 0
//		maxLen := 0
//		isDup := func(s string) bool {
//			m := make(map[rune]bool)
//			for _, v := range s {
//				if _, ok := m[v]; ok {
//					return true
//				}
//				m[v] = true
//			}
//			return false
//		}
//		for right := range s {
//			if !isDup(s[left : right+1]) {
//				curLen = (right + 1) - left
//				if curLen > maxLen {
//					maxLen = curLen
//				}
//			} else {
//				left++
//			}
//		}
//		return maxLen
//	}
//
// 上面的时间复杂度和空间复杂度都太高了，重做一下
func lengthOfLongestSubstring(s string) int {
	// 其实核心思想就是一个窗口
	m := make(map[byte]int)
	left, maxLen := 0, 0
	for right := 0; right < len(s); right++ {
		if index, ok := m[s[right]]; ok && index >= left {
			left = m[s[right]] + 1
		}
		m[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
