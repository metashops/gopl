package main

import (
	tree `gopl/algorithm/lc/common`
)

// 450. 删除二叉搜索树中的节点
func deleteNode(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	minnode := root.Right
	for minnode.Left != nil {
		minnode = minnode.Left
	}
	root.Val = minnode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *tree.TreeNode) *tree.TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}
