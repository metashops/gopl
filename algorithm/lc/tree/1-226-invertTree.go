package main

import (
	`fmt`

	tree "gopl/algorithm/lc/common"
)

// （1）递归版本的前序遍历
func invertTree1(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree1(root.Left)
	invertTree1(root.Right)

	return root
}

// （2）迭代版本的前序遍历
func invertTree2(root *tree.TreeNode) *tree.TreeNode {
	var stack []*tree.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left // 交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}

// 递归前序遍历
func preorderTraversal(root *tree.TreeNode) (res []int) {
	var traversal func(node *tree.TreeNode)
	traversal = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 递归后序遍历
func postorderTraversal(root *tree.TreeNode) (res []int) {
	var traversal func(node *tree.TreeNode)
	traversal = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func main() {
	n6 := tree.TreeNode{Val: 6, Left: nil, Right: nil}
	n4 := tree.TreeNode{Val: 4, Left: nil, Right: &n6}
	n3 := tree.TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := tree.TreeNode{Val: 2, Left: nil, Right: nil}
	n1 := tree.TreeNode{Val: 1, Left: &n2, Right: &n3}

	tn := tree.NewTreeNode(5, &n1, &n4)
	tree1 := invertTree1(tn)

	res := preorderTraversal(tree1)
	fmt.Println(res)
}
