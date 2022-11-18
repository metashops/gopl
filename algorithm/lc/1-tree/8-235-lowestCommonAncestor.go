package main

import (
	tree `gopl/algorithm/lc/common`
)

// 235. 二叉搜索树的最近公共祖先

// 利用BSL的性质（前序遍历有序）
func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val { // 当前节点的值大于给定的值，则说明满足条件的在左边
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val { // 当前节点的值小于各点的值，则说明满足条件的在右边
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	} // 当前节点的值在给定值的中间（或者等于），即为最深的祖先
}
