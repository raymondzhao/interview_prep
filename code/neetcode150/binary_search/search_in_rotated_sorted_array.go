package binarysearch

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly left rotated at an unknown index
k (1 <= k < nums.length) such that the resulting array is
[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target,
return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	// need 3 pointers - front, middle, back
	// we need to determine which half is sorted and then check if target fits within that range
	// if it is - move the pointers, if its not, it's in the other half

	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start <= end { // iterate through until we have either found the value or our pointers have moved
		middle := (start + end) / 2 // use start and end here so we can keep recalculating the middle

		if nums[middle] == target {
			return middle
		}

		// check if left is sorted
		if nums[start] <= nums[middle] {
			// check if target is in this half - else its in the other half
			if target >= nums[start] && target < nums[middle] {
				end = middle - 1
			} else {
				start = middle + 1
			}
			// right is sorted
		} else {
			// same thing here
			if target > nums[middle] && target <= nums[end] {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}

	return -1
}
