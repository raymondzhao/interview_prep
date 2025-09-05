from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # need two pointers - prev and current - rewrite .next to point to prev as we go along the list
        prev, curr = None, head
        while curr:
            # save curr.next as temp so we can keep iterating
            temp = curr.next
            # rewrite pointers to look at prev value
            curr.next = prev

            # update our pointers to move to next prev and curr
            prev = curr
            curr = temp

        # solution is in prev bc that shows the list all the way back
        return prev

print(Solution().reverseList(ListNode(1,(ListNode(2,ListNode(3,None))))))