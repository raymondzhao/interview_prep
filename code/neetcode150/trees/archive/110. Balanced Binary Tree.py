# Definition for a binary tree node.
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        # easiest way is to check heights of each subtree recursively
        def dfs(root) -> []:
            # base case - empty tree will be balanced and have a height of 0
            if not root:
                return [True, 0]
            
            # traverse down left and right subtrees
            left,right = dfs(root.left), dfs(root.right)

            # calculate height
            height = 1 + max(left[1],right[1])
            # calculate if its balanced - balanced binary trees have subtrees balanced so we have to check heights
            balanced = left[0] and right[0] and abs(left[1] - right[1]) <= 1

            return [balanced,height]
        
        return dfs(root)[0]
