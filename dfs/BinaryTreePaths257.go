package main

import (
	"strconv"
	"strings"
	"fmt"
)


//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, path []int, result *[]string) {
	if root == nil {
		return
	}

	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		if len(path) > 0 {
			s := []string{}
			for i := 0; i < len(path); i++ {
				s = append(s, strconv.Itoa(path[i]))
			}
			*result = append(*result, strings.Join(s, "->"))
		}
		return
	}

	helper(root.Left, path, result)
	helper(root.Right, path, result)
}

func binaryTreePaths(root *TreeNode) []string {
	result := &[]string{}
	path := []int{}
	helper(root, path, result)
	return *result

}

func main() {
	root := TreeNode{1,nil,nil}
	firstNode := TreeNode{2,nil,nil}
	secNode := TreeNode{3,nil,nil}
	root.Left = &firstNode
	root.Right = &secNode
	result := binaryTreePaths(&root)
	fmt.Println(result)
}
