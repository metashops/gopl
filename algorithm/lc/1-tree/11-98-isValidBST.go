package main

import (
	tree "gopl/algorithm/lc/common"
)

/*
思路：中序遍历下，输出的二叉搜索树节点的数值是有序序列，就相当于变成了判断一个序列是否递增的了
*/

// 中序遍历解法
func isValidBST(root *tree.TreeNode) bool {
	var prev *tree.TreeNode
	var travel func(node *tree.TreeNode) bool
	travel = func(node *tree.TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}
