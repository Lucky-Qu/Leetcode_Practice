package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点，用于简化链表操作
	current := dummy     // 当前节点指针
	carry := 0           // 进位

	for l1 != nil || l2 != nil || carry > 0 {
		// 获取当前位的值
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next // 移动到下一个节点
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next // 移动到下一个节点
		}

		// 计算总和和新的进位
		sum := val1 + val2 + carry
		carry = sum / 10

		// 创建新节点并移动指针
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next // 返回结果链表的头节点
}
func main() {

}
