package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	newNode := head

	for head != nil {
		new := head.Next
		newNode.Next = prev

		prev = newNode

		newNode = new
	}

	return newNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newNode := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}
