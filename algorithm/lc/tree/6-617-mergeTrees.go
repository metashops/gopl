package main

import (
	tree `gopl/algorithm/lc/common`
)

func main() {

}

// 前序遍历简洁版
func mergeTrees(root1 *tree.TreeNode, root2 *tree.TreeNode) *tree.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
