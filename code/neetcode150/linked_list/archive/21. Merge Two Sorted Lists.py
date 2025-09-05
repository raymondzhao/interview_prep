from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    # list1 = [1 -> 2 -> 4], list2 = [1 -> 3 -> 4]
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        # create dummy node in case of inserting into empty list
        dummy = ListNode()
        # create return list
        tail = dummy
        # compare values of list1 and list2 and add to return list
        while list1 and list2:
            if list1.val < list2.val:
                tail.next = list1
                list1 = list1.next
            else:
                tail.next = list2
                list2 = list2.next
            tail = tail.next
        # chain the rest of whatever list to the end 
        if list1:
            tail.next = list1
        else:
            tail.next = list2

        return dummy.next