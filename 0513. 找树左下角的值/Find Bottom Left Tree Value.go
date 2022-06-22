package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历 先加入右子树 每层的最左边节点一定是最后访问到的
func findBottomLeftValue(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		ans = node.Val
	}
	return
}
