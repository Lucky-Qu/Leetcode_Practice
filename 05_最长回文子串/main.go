package main

import "fmt"

func longestPalindrome(s string) string {
	b := []byte(s)
	l := len(b)
	var rb, re = 0, 0
	//从第一项开始逐个匹配
	for i := 0; i < l-(re-rb); i++ {
		//匹配最后一项是否相等，如果不相等往前一位判定，如果相等检测是否为回文数
		for j := l - 1; j >= i; j-- {
			//检测最后一项是否相等
			if b[i] != b[j] {
				//往前一位判定
				continue
			} else {
				//检测是否为回文数
				var isHuiwen = true
				testHuiwen := b[i : j+1]
				huiwenLen := len(testHuiwen)
				for k := 0; k < huiwenLen/2; k++ {
					if testHuiwen[k] != testHuiwen[huiwenLen-1-k] {
						isHuiwen = false
						break
					}
				}
				if isHuiwen {
					if j-i > re-rb {
						rb = i
						re = j
					}
				}
			}
		}
	}
	return string(b[rb : re+1])
}
func main() {
	s := "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"
	fmt.Println(longestPalindrome(s))
	c := []byte(s)
	fmt.Println(string(c[0:2]))
}
