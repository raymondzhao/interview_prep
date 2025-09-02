from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        depth = 0
        # base case when there is no node
        if not root:
            return depth
        # process current node 
        depth += 1

        lDepth, rDepth = 0,0

        # look left
        if root.left:
            lDepth = self.maxDepth(root.left)
        # look right       
        if root.right:
            rDepth = self.maxDepth(root.right)
        
        return depth + max(lDepth, rDepth)