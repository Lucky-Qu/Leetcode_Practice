package main

import (
	"fmt"
)

// 下方为我第一次做时的题解
//
//	func longestPalindrome(s string) string {
//		b := []byte(s)
//		l := len(b)
//		var rb, re = 0, 0
//		//从第一项开始逐个匹配
//		for i := 0; i < l-(re-rb); i++ {
//			//匹配最后一项是否相等，如果不相等往前一位判定，如果相等检测是否为回文数
//			for j := l - 1; j >= i; j-- {
//				//检测最后一项是否相等
//				if b[i] != b[j] {
//					//往前一位判定
//					continue
//				} else {
//					//检测是否为回文数
//					var isHuiwen = true
//					testHuiwen := b[i : j+1]
//					huiwenLen := len(testHuiwen)
//					for k := 0; k < huiwenLen/2; k++ {
//						if testHuiwen[k] != testHuiwen[huiwenLen-1-k] {
//							isHuiwen = false
//							break
//						}
//					}
//					if isHuiwen {
//						if j-i > re-rb {
//							rb = i
//							re = j
//						}
//					}
//				}
//			}
//		}
//		return string(b[rb : re+1])
//	}
//
// 下方为看了题解后学会的中心扩散法
func longestPalindrome(s string) string {
	b := []byte(s)
	l := len(b)
	if l < 2 {
		return s
	}

	var start, end int

	// 内部函数：给定左右指针向外扩散
	expand := func(left, right int) {
		for left >= 0 && right < l && b[left] == b[right] {
			if right-left > end-start {
				start, end = left, right
			}
			left--
			right++
		}
	}

	for i := 0; i < l; i++ {
		expand(i, i)   // 奇数长度中心
		expand(i, i+1) // 偶数长度中心
	}
	return s[start : end+1]
}
func main() {
	s := "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"
	fmt.Println(longestPalindrome(s))
	c := []byte(s)
	fmt.Println(string(c[0:2]))
}
