package main

import "fmt"

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

//func reverseVowels(s string) string {
//	//个人解题思路：先遍历字符串，将元音字母的索引存在一个数组中，再将数组反转，对应的值进行互换
//	index := make([]int, 0, 10)
//	for i, v := range s {
//		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
//			index = append(index, i)
//		}
//	}
//	reverseIndex := make([]int, len(index))
//	for i, v := range index {
//		reverseIndex[len(index)-i-1] = v
//	}
//	result := []byte(s)
//	for i, v := range index {
//		result[v] = s[reverseIndex[i]]
//	}
//	return string(result)
//}

//上面的方法可行，但是时间复杂度和空间复杂度都很高，采用双指针进行优化会更好

// 双指针思路：边找边换，一次遍历解决
func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	byteS := []byte(s)
	for left < right {
		if vowels[byteS[left]] && vowels[byteS[right]] {
			byteS[left], byteS[right] = byteS[right], byteS[left]
			left++
			right--
			continue
		}
		if vowels[byteS[left]] {
			right--
			continue
		}
		left++
	}
	return string(byteS)
}

func main() {
	fmt.Println(reverseVowels("LeetCode"))
}
