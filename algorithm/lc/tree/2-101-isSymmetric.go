package main

import (
	tree `gopl/algorithm/lc/common`
)

// 对称二叉树
func isSymmetric(root *tree.TreeNode) bool {
	var defs func(left *tree.TreeNode, right *tree.TreeNode) bool
	defs = func(left *tree.TreeNode, right *tree.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return defs(left.Left, right.Right) && defs(right.Left, left.Right)
	}
	return defs(root.Left, root.Right)
}

// func main() {
// 	n6 := tree.TreeNode{Val: 6, Left: nil, Right: nil}
// 	n4 := tree.TreeNode{Val: 4, Left: nil, Right: &n6}
// 	n3 := tree.TreeNode{Val: 3, Left: nil, Right: nil}
// 	n2 := tree.TreeNode{Val: 2, Left: nil, Right: nil}
// 	n1 := tree.TreeNode{Val: 1, Left: &n2, Right: &n3}
//
// 	tn := tree.NewTreeNode(5, &n1, &n4)
// 	tree := isSymmetric(tn)
// 	fmt.Println(tree)
// }
