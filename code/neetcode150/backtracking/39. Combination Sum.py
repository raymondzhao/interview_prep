from typing import List

class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        all = []
        
        def dfs(i: int,cur: List[int], total: int):
            # base case - if we sum to target, return
            if total == target:
                all.append(cur.copy())
                return
            
            if i >= len(candidates) or total > target:
                return

            # include the candidate
            cur.append(candidates[i])
            dfs(i, cur, total + candidates[i])

            cur.pop()
            # i + 1 bc we can't do anything with this candidate
            dfs(i + 1, cur, total)

        dfs(0, [], 0) 

        return all

print(Solution().combinationSum([2,3,6,7],7))
            


                



        
