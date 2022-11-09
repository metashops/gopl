package main

import (
	"fmt"

	tree "gopl/algorithm/lc/common"
)

// 完全二叉树的节点个数
func countNodes(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	LeftH, RightH := 0, 0 // 计算两边高度
	ln, rn := root, root
	for ln != nil {
		LeftH++
		ln = ln.Left
	}
	for rn != nil {
		RightH++
		rn = rn.Right
	}
	if LeftH == RightH {
		return 1<<LeftH - 1 // 左移n位就是乘以2的n次方
	}
	// 当前子树不是完美二叉树，只是完全二叉树，递归处理左右子树
	// + 1 就是跟节点
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	n6 := tree.TreeNode{Val: 6, Left: nil, Right: nil}
	n4 := tree.TreeNode{Val: 4, Left: nil, Right: &n6}
	n3 := tree.TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := tree.TreeNode{Val: 2, Left: nil, Right: nil}
	n1 := tree.TreeNode{Val: 1, Left: &n2, Right: &n3}

	tn := tree.NewTreeNode(5, &n1, &n4)
	t := countNodes(tn)
	fmt.Println(t)

}
