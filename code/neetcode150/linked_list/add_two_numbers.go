package linked_list

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// create dummy head
	dummy := new(ListNode)
	tail := dummy
	// keep track of carryover
	carry := 0

	// go through the list as long as we have a value from both lists
	for l1 != nil || l2 != nil {
		r1, r2 := 0, 0
		if l1 != nil {
			r1 = l1.Val
			// once we grab value we move
			l1 = l1.Next
		}
		if l2 != nil {
			r2 = l2.Val
			// once we grab value we move
			l2 = l2.Next
		}
		sum := r1 + r2 + carry
		// use int math to see if its 1 or 0
		carry = sum / 10
		// use mod to get the digit in ones place
		digit := sum % 10
		// connect our new node to our tail which acts as a iterator ptr
		tail.Next = &ListNode{
			Val: digit,
		}
		// move ptr to next value
		tail = tail.Next
	}
	// if we still have carry over, we need an extra node at the end
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return dummy.Next
}
