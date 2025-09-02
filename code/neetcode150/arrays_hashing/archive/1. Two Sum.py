from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            num1 = nums[i]
            for j in range(len(nums)):
                if j != i and num1 + nums[j] == target:
                    return [i,j]
        return []
    
    def betterTwoSum(self, nums: List[int], target: int) -> List[int]:
        # create hashMap to keep track of values
        valIndex = {}
        for i,num in enumerate(nums):
            # use the difference between target and num
            if target - num in valIndex:
                return [valIndex[target - num], i]
            else:
                # add to hashmap if not found so later on we might find it
                valIndex[num] = i
        return


sol = Solution()
print(sol.twoSum([3,2,4],6))