from typing import List
import heapq

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        h = list(map(lambda x: x * -1, stones))
        heapq.heapify(h)
        while len(h) > 1:
            # grab the two biggest stones
            stone1 = heapq.heappop(h) * -1
            stone2 = heapq.heappop(h) * -1

            if stone1 != stone2:
                heapq.heappush(h,-1 * abs(stone1-stone2))
            
        if len(h) == 0:
            return 0
        return -1 * heapq.heappop(h) 

print(Solution().lastStoneWeight([2,2]))