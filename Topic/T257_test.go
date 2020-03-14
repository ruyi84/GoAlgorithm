package main

import (
	"fmt"
	"testing"
)

type TreeNode1 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//[1,2,3,null,5]
func Test_257(t *testing.T) {
	var node *TreeNode1
	node.Val = 1
	node.Left.Val = 2
	node.Right.Val = 3
	node.Left.Right.Val = 5

	paths := binaryTreePaths(node)
	fmt.Println(paths)
}

func binaryTreePaths(root *TreeNode1) []string {

}

func getNext(root *TreeNode1) (*TreeNode1, bool) {

	return
}
