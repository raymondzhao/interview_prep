package trees

/*
Given a binary search tree (BST),
find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined
between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// BST - left is less than the root, and the right is greater than the root
	// for a BST, we just need to find the split aka p < node > q - bc thats the LCA
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			// go left
			curr = curr.Left

		} else if p.Val > curr.Val && q.Val > curr.Val {
			// go right
			curr = curr.Right
		} else {
			// otherwise return the current
			return curr
		}
	}
	return nil
}
