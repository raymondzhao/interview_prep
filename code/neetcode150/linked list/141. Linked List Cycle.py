# Definition for singly-linked list.
from typing import Optional

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        # initialize two pointers at head
        slowPointer = fastPointer = head

        # continue loop since if the loop ends, we dont have a cycle
        while fastPointer and fastPointer.next:

            # move forward pointers (slow + 1, fast + 2)
            slowPointer = slowPointer.next
            if fastPointer.next:
                fastPointer = fastPointer.next.next

            # if fast "catches up" we have a cycle
            if slowPointer == fastPointer:
                return True
        
        return False