package main

import (
	tree `gopl/algorithm/lc/common`
)

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		root = &tree.TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
