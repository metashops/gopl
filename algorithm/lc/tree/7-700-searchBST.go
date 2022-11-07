package main

import (
	tree "gopl/algorithm/lc/common"
)

/**
题目：给定二叉搜索树（BST）的根节点root和一个整数值val。
	 你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。 如果节点不存在，则返回null。
*/

func main() {

}

// 迭代
func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			break
		}
	}
	return nil
}
