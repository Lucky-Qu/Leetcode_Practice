package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int) // 哈希表存储字符及其索引
	left := 0               // 滑动窗口的左边界
	result := 0             // 记录最长子串的长度

	// 遍历字符串中的字符
	for right := 0; right < len(s); right++ {
		// 如果当前字符已经存在于滑动窗口中，则更新左边界
		if idx, ok := m[s[right]]; ok && idx >= left {
			left = idx + 1
		}

		// 更新当前字符的索引
		m[s[right]] = right

		// 更新最长子串的长度
		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
func main() {}
