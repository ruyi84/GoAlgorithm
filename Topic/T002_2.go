package main

func addTwoNumbers_2(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val

	nextNode := addTwoNumbers_2(l1, l2)

	if sum < 10 {
		return &ListNode{Val: sum, Next: nextNode}
	} else {

	}

	return nextNode
}
