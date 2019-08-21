package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {

	var i int
	var f func(root *TreeNode, L int, R int)
	f = func(root *TreeNode, L int, R int) {
		if root != nil {
			if root.Val < R && root.Val > L {
				i += root.Val
				fmt.Println(i)
			}
			f(root.Left, L, R)
			f(root.Right, L, R)
		}
	}

	f(root, L, R)

	return i
}
