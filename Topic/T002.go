package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) apeend(l1 *ListNode) {
	l.Next = l1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := l1.Val
	num2 := l1.Val
	L1 := l1.Next
	L2 := l1.Next
	if L1 != nil {
		num1 += L1.Val + num1*10
		num2 += L2.Val + num2*10
		L1 = L1.Next
		L2 = L2.Next
	}
	num := num1 + num2
	var L *ListNode
	L.Val = num / 100
	L.Next.Val = num / 10 % 10
	L.Next.Next.Val = num % 100 % 10
	return L
}

func main() {

}

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解决方法，第一时间想到的时便利，然后相加重新赋值
也就是说，第一次进行遍历，分别取出两个链表的三个值，然后组成一个三位数，然后相加
这时会产生一个新的三位数，然后将这个三位数重新组成一个新的链表
*/
