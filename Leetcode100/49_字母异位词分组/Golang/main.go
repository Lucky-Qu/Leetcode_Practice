package main

import "fmt"

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//
//
//示例 1:
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
//解释：
//
//在 strs 中没有字符串可以通过重新排列来形成 "bat"。
//字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
//字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	m := make(map[[26]int][]string)
	for i, v := range strs {
		set := [26]int{0}
		for j := 0; j < len(v); j++ {
			set[v[j]-'a']++
		}
		m[set] = append(m[set], strs[i])
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "nat", "bat"}))
}
