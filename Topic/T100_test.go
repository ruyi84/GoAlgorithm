package main

import (
	"testing"
)

/*
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode_100 struct {
	Val   int
	Left  *TreeNode_100
	Right *TreeNode_100
}

func isSameTree(p *TreeNode_100, q *TreeNode_100) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return !isSameTree(p.Right, q.Right) && !isSameTree(p.Left, q.Left)
}

func isSameTree2(p *TreeNode_100, q *TreeNode_100) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	arrp, arrq := []*TreeNode_100{p}, []*TreeNode_100{q}
	for len(arrp) > 0 && len(arrq) > 0 {
		node1, node2 := arrp[0], arrq[0]

		arrp, arrq = arrp[1:], arrq[1:]

		if node1.Val != node2.Val {
			return false
		}

		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right

		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}

		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}

		if left1 != nil {
			arrp = append(arrp, left1)
		}
		if right1 != nil {
			arrp = append(arrp, right1)
		}
		if left2 != nil {
			arrq = append(arrq, left2)
		}
		if right2 != nil {
			arrq = append(arrq, right2)
		}
	}
	return len(arrp) == 0 && len(arrq) == 0
}

func Test_T100(t *testing.T) {
	var a TreeNode_100
	a.Val = 1
	a.Left = &TreeNode_100{
		Val: 2,
	}
	a.Right = &TreeNode_100{
		Val: 3,
	}

	var b TreeNode_100
	b.Val = 1
	b.Left = &TreeNode_100{
		Val: 2,
	}
	b.Right = &TreeNode_100{
		Val: 3,
	}

	isSameTree(&a, &b)
}
