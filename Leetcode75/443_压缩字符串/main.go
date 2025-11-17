package main

import (
	"fmt"
)

func compress(chars []byte) int {
	editPointer := 0
	positionPointer := 0

	for positionPointer < len(chars) {
		char := chars[positionPointer]
		count := 0

		// 统计连续字符
		for positionPointer < len(chars) && chars[positionPointer] == char {
			positionPointer++
			count++
		}

		// 写字符
		chars[editPointer] = char
		editPointer++

		// 写数字（如果 count > 1）
		if count > 1 {
			start := editPointer
			for c := count; c > 0; c /= 10 {
				chars[editPointer] = byte(c%10) + '0'
				editPointer++
			}
			reverse(chars, start, editPointer-1)
		}
	}

	return editPointer
}

func reverse(chars []byte, l, r int) {
	for l < r {
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
}

func main() {
	chars := []byte("abcccccc")
	n := compress(chars)
	fmt.Println(n, string(chars[:n])) // 4 abc6
}
