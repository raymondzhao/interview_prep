import heapq
from typing import List

class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.h = nums
        self.k = k
        heapq.heapify(self.h)
        while len(self.h) > k:
            heapq.heappop(self.h)


    def add(self, val: int) -> int:
        heapq.heappush(self.h, val)
        if len(self.h) > self.k:
            heapq.heappop(self.h)
        return self.h[0]


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val) 

obj = KthLargest(3,[4, 5, 8, 2])
print(obj.add(3))
print(obj.add(5))
print(obj.add(10))
print(obj.add(9))
print(obj.add(4))