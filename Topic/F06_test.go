package main

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	num := 0
	arr := make([]int, 10001)

	for head != nil {
		num++

		arr[10001-num] = head.Val
		head = head.Next
	}

	return arr[10000-num:]
}
