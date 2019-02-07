//https://leetcode.com/problems/smallest-string-starting-from-leaf/

package main

import (
	"fmt"
	"strings"
)

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

func printNode(root *TreeNode, level int) {

	fmt.Println(strings.Repeat("\t", level), root.Val)
	if root.Left != nil {
		printNode(root.Left, level+1)
	}
	if root.Right != nil {
		printNode(root.Right, level+1)
	}

}

func smallestFromLeaf2(root *TreeNode) string {

	if root == nil {
		return ""
	}

	left := smallestFromLeaf2(root.Left)
	right := smallestFromLeaf2(root.Right)

	if left == "" {
		return right + string('a'+root.Val)
	} else if right == "" {
		return left + string('a'+root.Val)
	} else if left < right {
		return left + string('a'+root.Val)
	} else {
		return right + string('a'+root.Val)
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

// unfinished
func smallestFromLeaf3(root *TreeNode) string {

	var stack []*TreeNode
	var smallestStr = ""
	var popStr = ""
	for {

		if len(stack) == 0 && root == nil {
			break
		}

		if root != nil {
			stack = append(stack, root)
			popStr = string('a'+root.Val) + popStr
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			if root.Right == nil && root.Left == nil {
				if smallestStr == "" || smallestStr > popStr {
					smallestStr = popStr
				}
			}

			stack = stack[:len(stack)-1]
			if len(stack)-2 > -1 && root != stack[len(stack)-2] && root.Right == nil {
				popStr = popStr[1:len(popStr)]
			}

			root = root.Right
			if root == nil {
				popStr = popStr[1:len(popStr)]
			}
		}

	}

	return smallestStr
}

func main() {

	list := []int{25, 1, 3, 1, 3, 0, 2}
	root := buildTree(list)
	fmtNode(root)
	printNode(root, 0)
	smallest := smallestFromLeaf(root)
	fmt.Println(smallest)

}
