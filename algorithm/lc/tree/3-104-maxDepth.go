package main

import (
	tree "gopl/algorithm/lc/common"
)

func maxDepth(root *tree.TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*tree.TreeNode{root}

	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	return depth
}

// func main() {
// 	n6 := tree.TreeNode{Val: 6, Left: nil, Right: nil}
// 	n4 := tree.TreeNode{Val: 4, Left: nil, Right: &n6}
// 	n3 := tree.TreeNode{Val: 3, Left: nil, Right: nil}
// 	n2 := tree.TreeNode{Val: 2, Left: nil, Right: nil}
// 	n1 := tree.TreeNode{Val: 1, Left: &n2, Right: &n3}
//
// 	tn := tree.NewTreeNode(5, &n1, &n4)
// 	t := maxDepth(tn)
// 	fmt.Println(t)
//
// }
