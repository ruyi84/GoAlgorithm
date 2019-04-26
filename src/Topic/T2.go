package main

import "fmt"

Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println(l1)
	return l1
}

func main() {
	addTwoNumbers([2, 1, 3], [5, 6, 4])
}
