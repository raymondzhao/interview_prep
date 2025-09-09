package trees

/*
Given the roots of two binary trees p and q,
write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical,
and the nodes have the same value.
*/

// using recursion to practice
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// base case is if both nodes are null - they are the same
	if p == nil && q == nil {
		return true
	}
	// if one of them is null, then return false
	if p == nil || q == nil {
		return false
	}
	// recurse down left and right subtrees and make sure to use &&
	// we need everything to match to return true
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
