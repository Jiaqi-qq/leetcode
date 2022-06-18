package main

// Definition for a Node.
type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	tmp := Node{Val: x}
	if aNode == nil { // 元素个数 == 0
		tmp.Next = &tmp
		return &tmp
	}
	if aNode.Next == aNode { // 元素个数 == 1
		tmp.Next = aNode
		aNode.Next = &tmp
		return aNode
	}
	// 元素个数 > 1
	cur := aNode
	for cur.Next != aNode {
		if cur.Next.Val >= x && cur.Val <= x { // 有<=x和x<=的值 注意要有等于号
			break
		}
		if cur.Val > cur.Next.Val {
			if x > cur.Val || x < cur.Next.Val { // 没有比x大的 or 没有比x小的 都加在最大最小值中间
				break
			}
		}
		cur = cur.Next
	}
	tmp.Next = cur.Next
	cur.Next = &tmp
	return aNode
}
