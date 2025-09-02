from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        # keep track of seen values using hashmap
        seenValues = set()

        # go through List and add values in hash - if there is duplicate, return out
        for num in nums:
            if num in seenValues:
                return True
            seenValues.add(num)
        
        return False

sol = Solution()
print(sol.containsDuplicate([1,2,3,3]))