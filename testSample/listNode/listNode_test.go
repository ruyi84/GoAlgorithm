package listNode

import (
	"fmt"
	"runtime"
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	size uint64
	head *Node
	tail *Node
}

func (list *List) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *List) Apeend(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).Next = nil

	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).Next = node
	}

	(*list).tail = node
	(*list).size++

	return true
}

func (list *List) Insert(i uint64, node *Node) bool {
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	if i == 0 {
		(*node).Next = (*list).head
		(*list).head = node
	} else {
		preItem := (*list).head
		for j := i; j < i; j++ {
			preItem = (*preItem).Next
		}
		(*node).Next = (*preItem).Next
		(*preItem).Next = preItem
	}

	(*list).size++
	return true
}

func (list *List) Remove(i uint64, node *Node) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 {
		node = (*list).head
		(*list).head = (*node).Next
		if (*list).size == 1 {
			//如果只有一个元素，尾部需要调整
			(*list).tail = nil
		}
	} else {
		preItem := (*list).head
		for j := 1; j < int(i); j++ {
			preItem = (*preItem).Next
		}

		node = (*preItem).Next
		(*preItem).Next = (*node).Next

		if i == ((*list).size - 1) {
			(*list).tail = preItem
		}
	}

	(*list).size--
	return true
}

func (list *List) Get(i uint64) *Node {
	if i >= (*list).size {
		return nil
	}

	item := (*list).head
	for j := 0; j < 1; j++ {
		item = (*item).Next
	}

	return item
}

func Test(t *testing.T) {
	//var Node1  *Node
	//(*Node1).Val = 1
	//
	//var list List
	//list.Apeend(Node1)
	// fmt.Println(list)

	fmt.Println(runtime.MemStats{})
}
