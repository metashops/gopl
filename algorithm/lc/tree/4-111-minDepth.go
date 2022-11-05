package main

import (
	"fmt"

	tree "gopl/algorithm/lc/common"
)

// 最小深度：根节点到最近叶子节点的最短路径上的节点数量
// 叶子节点：左右孩子都为空的节点才是叶子节点！

func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*tree.TreeNode{root}
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
}

func main() {
	n6 := tree.TreeNode{Val: 6, Left: nil, Right: nil}
	n4 := tree.TreeNode{Val: 4, Left: nil, Right: &n6}
	n3 := tree.TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := tree.TreeNode{Val: 2, Left: nil, Right: nil}
	n1 := tree.TreeNode{Val: 1, Left: &n2, Right: &n3}

	tn := tree.NewTreeNode(5, &n1, &n4)
	t := minDepth(tn)
	fmt.Println(t)

}
