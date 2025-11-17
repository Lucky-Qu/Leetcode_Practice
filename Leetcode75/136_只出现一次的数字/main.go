package main

//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

// 考查位运算符，没什么说的
func singleNumber(nums []int) int {
	num := 0
	for _, v := range nums {
		num ^= v
	}
	return num
}

func main() {

}
