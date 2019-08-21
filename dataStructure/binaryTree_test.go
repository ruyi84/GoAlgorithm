package dataStructure

import (
	"fmt"
	"reflect"
	"testing"
)

type TreeNode struct {
	lefTreeNode *TreeNode
	rightNode   *TreeNode
	value       int
}

func (t TreeNode) IsEmpty() bool {
	return reflect.DeepEqual(t, TreeNode{})
}

func (t TreeNode) fmt() {
	fmt.Println(t.value)
}

func Test_First(t *testing.T) {
	treeNode1 := TreeNode{
		value: 10,
	}
	treeNode2 := TreeNode{
		value: 5,
	}
	treeNode3 := TreeNode{
		value: 15,
	}
	treeNode4 := TreeNode{
		value: 3,
	}
	treeNode5 := TreeNode{
		value: 7,
	}

	treeNode1.lefTreeNode = &treeNode2
	treeNode1.rightNode = &treeNode3

	treeNode3.lefTreeNode = &treeNode4
	treeNode3.rightNode = &treeNode5

	treeNode1.inTraerseBTree()
}

func (a *TreeNode) preTraverseBtree() {
	if a != nil {
		fmt.Println(a.value)
		a.lefTreeNode.preTraverseBtree()
		a.rightNode.preTraverseBtree()
	}

}

func (a *TreeNode) inTraerseBTree() {
	if a != nil {
		a.lefTreeNode.inTraerseBTree()
		a.fmt()
		a.rightNode.inTraerseBTree()
	}
}
