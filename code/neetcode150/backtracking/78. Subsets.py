from typing import List

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        all = []

        subset = []
        def dfs(i):
            # base case - if there are no more values to check decisions in nums
            if i >= len(nums):
                all.append(subset.copy()) # add subset bc we are done with all these
                return
            
            # what are our decisions/branches
            # add the next val @ i
            subset.append(nums[i])
            dfs(i+1) # continue down the tree since there are more vals

            # don't add the next i
            subset.pop()
            dfs(i+1) # continue down the other side of tree
            
        dfs(0) # start the recursion at first value
        return all