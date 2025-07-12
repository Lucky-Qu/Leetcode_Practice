package main

import "fmt"

//给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
//如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
//返回 合并后的字符串。

//func mergeAlternately(word1 string, word2 string) string {
//	//比较长度，将多出的存起来以便追加
//	l1 := len(word1)
//	l2 := len(word2)
//	var a string
//	if l1 > l2 {
//		a = word1[l2:]
//		word1 = word1[:l2]
//	} else {
//		a = word2[l1:]
//		word2 = word2[:l1]
//	}
//
//	mergeSameLenString := func(s1 string, s2 string) string {
//		//合并相同长度字符串
//		b1 := []byte(s1)
//		b2 := []byte(s2)
//		var s string = ""
//		for i := 0; i < len(b1); i++ {
//			s += string(b1[i])
//			s += string(b2[i])
//		}
//		return s
//	}
//
//	return mergeSameLenString(word1, word2) + a
//}

//上面自己第一次做的昏完了属于是，下面是看了题解后的双指针法

func mergeAlternately(word1 string, word2 string) string {
	i, j := len(word1), len(word2)
	result := make([]byte, 0, len(word1)+len(word2))
	for k := 0; k < i || k < j; k++ {
		if k < i {
			result = append(result, word1[k])
		}
		if k < j {
			result = append(result, word2[k])
		}
	}
	fmt.Println(result)
	return string(result)
}

func main() {
	fmt.Println(mergeAlternately("abcdasdd", "de"))
}
