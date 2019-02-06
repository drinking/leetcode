package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(list []int) *TreeNode {

	var nodes []TreeNode
	for _, v := range list {
		node := TreeNode{
			v,
			nil,
			nil,
		}
		nodes = append(nodes, node)
	}

	length := len(nodes)
	for i, v := range nodes {
		if i*2+1 < length {
			v.Left = &nodes[i*2+1]
		}
		if i*2+2 < length {
			v.Right = &nodes[i*2+2]
		}
		nodes[i] = v
	}
	return &nodes[0]

}

func fmtNode(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Val == -1 {
		root.Left = nil
	}

	if root.Right != nil && root.Right.Val == -1 {
		root.Right = nil
	}

	if root.Left != nil {
		fmtNode(root.Left)
	}
	if root.Right != nil {
		fmtNode(root.Right)
	}

}

func smallestFromLeaf(root *TreeNode) string {

	var smallestStr string = ""
	var y func(r *TreeNode, result string)
	y = func(r *TreeNode, result string) {

		result = string('a'+r.Val) + result

		if r.Left != nil {
			y(r.Left, result)
		}

		if r.Right != nil {
			y(r.Right, result)
		}

		if r.Left == nil && r.Right == nil {
			if smallestStr == "" {
				smallestStr = result
			} else if smallestStr > result {
				smallestStr = result
			}
		}
	}

	y(root, "")
	return smallestStr
}

func main() {

	list := []int{2, 2, 1, -1, 1, 0, -1, 0}
	root := buildTree(list)
	fmtNode(root)
	smallest := smallestFromLeaf(root)
	fmt.Println(smallest)

}
