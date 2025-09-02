from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        # deal with base case
        if not root: return

        # swap left and right
        left = root.left
        root.left = root.right
        root.right = left

        # recursively go down both subtrees
        self.invertTree(root.left)
        self.invertTree(root.right)

        return root