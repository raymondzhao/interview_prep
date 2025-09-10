package trees

/*
Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).
*/

import (
	"container/list" // using list for o(1) queue operations
)

// create an object for keeping track of the level
type NodeLevel struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{} // each index represents the level
	if root == nil {
		return res
	}

	// create list
	q := list.New()
	q.PushBack(NodeLevel{root, 0})

	// bfs traversal
	for q.Len() > 0 {
		front := q.Front()
		n := front.Value.(NodeLevel)
		q.Remove(front)

		// keep track of the level
		level := n.Level

		// need to append to res so we can represent the new level
		if level == len(res) {
			res = append(res, []int{})
		}

		// process node and add val to res
		res[level] = append(res[level], n.Node.Val)

		if n.Node.Left != nil {
			q.PushBack(NodeLevel{n.Node.Left, level + 1}) // add 1 to the level
		}
		if n.Node.Right != nil {
			q.PushBack(NodeLevel{n.Node.Right, level + 1})
		}
	}

	return res
}
