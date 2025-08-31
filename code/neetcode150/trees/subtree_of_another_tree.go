package trees

/*
Given the roots of two binary trees root and subRoot, return true if there is
a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this
node's descendants. The tree tree could also be considered as a subtree of itself.
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// need to traverse the tree and check for candidates
	// if we have a candidate we traverse that and the subtree to make sure they are equal

	// check edge cases
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	rootQ := []*TreeNode{root} // store visited nodes for tree traversal

	for len(rootQ) > 0 {
		node := rootQ[0]
		rootQ = rootQ[1:] // dequeue after we point

		// found candidate - need to check if values equal compared to subRoot
		if node.Val == subRoot.Val {
			if treesEqual(node, subRoot) {
				return true
			}
		}

		// keep traversing tree via BFS
		if node.Left != nil {
			rootQ = append(rootQ, node.Left)
		}
		if node.Right != nil {
			rootQ = append(rootQ, node.Right)
		}
	}

	return false
}

func treesEqual(a, b *TreeNode) bool {
	// use a tuple so we can check both trees in one comparison
	q := [][2]*TreeNode{{a, b}}

	// traverse both trees (from candidate and subtree)
	for len(q) > 0 {
		pair := q[0]
		q = q[1:]
		n1, n2 := pair[0], pair[1] // compare the nodes

		// its the same keep going
		if n1 == nil && n2 == nil {
			continue
		}
		// one value is nil
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		// add the pairs to the q
		q = append(q, [2]*TreeNode{n1.Left, n2.Left})
		q = append(q, [2]*TreeNode{n1.Right, n2.Right})
	}
	return true
}
