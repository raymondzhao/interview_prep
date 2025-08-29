package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/

// recursion example

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {

		// base case - if the node is nil, then return 0
		if node == nil {
			return 0
		}

		// go left subtree
		left := depth(node.Left)

		// go right subtree
		right := depth(node.Right)

		// figure out the diameter computing the left and right sum
		if left+right > diameter {
			diameter = left + right
		}

		// make sure to add the height
		return 1 + max(left, right)
	}

	depth(root)
	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTreeIterative(root *TreeNode) int {
	diameter := 0
	// keep track of visited nodes
	visited := make(map[*TreeNode]bool)
	// calculate depths as we iterate
	depth := make(map[*TreeNode]int)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if curr == nil {
			stack = stack[:len(stack)-1] // pop value if nothing is there
		}

		if visited[curr] {
			// if we've already seen it, then we can calculate depth
			// since the node's children's depth have already been calculated
			stack = stack[:len(stack)-1] // pop

			left, right := depth[curr.Left], depth[curr.Right]

			if left+right > diameter {
				diameter = left + right
			}
			depth[curr] = 1 + max(left, right)
		} else {
			// if we haven't seen it - we still have to map the tree's structure
			visited[curr] = true
			stack = append(stack, curr.Left, curr.Right)
		}
	}

	return diameter
}
