package main

// 快慢指针
// 同时定义两个指针，分别为快指针和慢指针，快指针每次移动两个节点，慢指针每次移动一个节点
// 所以，当快指针移动完整个链表时，慢指针应该正处于整个链表的中间部位
func middleNode(head *ListNode) *ListNode {
	first, last := head, head

	for first != nil && first.Next != nil {
		first = first.Next.Next
		last = last.Next
	}

	return last
}

func middleNode1(head *ListNode) *ListNode {
	cur := head
	num := 0
	for cur != nil {
		cur = cur.Next
		num++
	}

	for i := 0; i < num/2; i++ {
		head = head.Next
	}
	return head
}
