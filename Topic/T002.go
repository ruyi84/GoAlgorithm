package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := 0
	num1 := 0

	x := 0
	y := 0

	l3 := new(ListNode)
	l4 := l3

	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum = x + y + num1
		fmt.Println(sum, x, y, num1)
		num1 = sum / 10
		l3.Val += sum % 10

		if l1 == nil && l2 == nil {
			break
		}

		l3.Next = new(ListNode)
		l3 = l3.Next
		x, y = 0, 0
	}

	if num1 != 0 {
		l3.Next = new(ListNode)
		l3 = l3.Next
		l3.Val = num1
	}

	return l4
}
