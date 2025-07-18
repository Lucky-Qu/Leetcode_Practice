package main

import (
	"fmt"
	"strings"
)

//给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
//
//func reverseWords(s string) string {
//	//个人思路：双指针法
//	b := []byte(s)
//	start1, end1, end2, start2 := 0, 1, len(b)-2, len(b)-1
//	mid1 := make([]byte, 0)
//	mid2 := make([]byte, 0)
//	for end1 < end2 {
//		//从前往后匹配一个单词
//		for b[start1] == ' ' || b[end1] != ' ' {
//			if b[start1] == ' ' {
//				start1++
//				end1++
//				continue
//			}
//			if b[end1] != ' ' {
//				end1++
//			}
//		}
//		//从后往前匹配一个单词
//		for b[start2] == ' ' || b[end2] != ' ' {
//			if b[start2] == ' ' {
//				start2--
//				end2--
//				continue
//			}
//			if b[end2] != ' ' {
//				end2--
//			}
//		}
//		//存储单词
//		mid1 = append(mid1, b[start1:end1]...)
//		mid1 = append(mid1, ' ')
//		mid2 = append(mid2, b[end2+1:start2+1]...)
//		mid2 = append(mid2, ' ')
//		fmt.Println(string(mid1))
//		fmt.Println(string(mid2))
//		//进入下一次匹配准备
//		start1 = end1 + 1
//		end1 = start1 + 1
//		start2 = end2 - 1
//		end2 = start2 - 1
//	}
//	//去除前导空格，尾随空格，间隔大于一的空格
//	//返回最终结果
//	return ""
//}
//
//上面蠢完了，做一半突然发现一个倒着遍历单词就解决了
//

//func reverseWords(s string) string {
//	b := []byte(s)
//	b = append([]byte{' ', ' '}, b...)
//	b = append(b, ' ')
//	mid2 := make([]byte, 0)
//	end2, start2 := len(b)-2, len(b)-1
//	for end2 > 0 {
//		for (b[start2] == ' ' || b[end2] != ' ') && end2 > -1 {
//			if b[start2] == ' ' {
//				start2--
//				end2--
//				continue
//			}
//			if b[end2] != ' ' {
//				end2--
//			}
//		}
//		mid2 = append(mid2, b[end2+1:start2+1]...)
//		mid2 = append(mid2, ' ')
//		start2 = end2 - 1
//		end2 = start2 - 1
//	}
//	start2 = len(mid2) - 1
//	for mid2[start2] == ' ' {
//		start2--
//	}
//	return string(mid2[:start2+1])
//}

// 原来有库函数就可以直接实现=.=
func reverseWords(s string) string {
	words := strings.Fields(s) // 步骤1：清理空格并分割单词
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // 步骤2：反转单词顺序
	}
	return strings.Join(words, " ") // 步骤3：连接单词
}

func main() {
	fmt.Println(reverseWords(" a good   example"))
}
