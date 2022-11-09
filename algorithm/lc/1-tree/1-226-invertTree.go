package main

import (
	tree "gopl/algorithm/lc/common"
)

type RN struct {
}

func NewRN() *RN {
	return &RN{}
}

// InvertTree1 （1）递归版本的前序遍历
func (r *RN) InvertTree1(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	r.InvertTree1(root.Left)
	r.InvertTree1(root.Right)

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

// PreorderTraversal 递归前序遍历
func (r *RN) PreorderTraversal(root *tree.TreeNode) (res []int) {
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
