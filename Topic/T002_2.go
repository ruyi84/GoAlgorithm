package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = new(ListNode)
	}
	if l2 == nil {
		l2 = new(ListNode)
	}

	l1.Val += l2.Val
	if l1.Val >= 10 {
		fmt.Println(l1.Val)
		l1.Val %= 10
		if l1.Next == nil {
			l1.Next = &ListNode{
				Val: 1,
			}
		} else {
			l1.Next.Val += 1
		}
	}

	if l1.Next == nil && l2.Next == nil {
		return l1
	}

	l1.Next = addTwoNumbers(l1.Next, l2.Next)

	return l1
}
