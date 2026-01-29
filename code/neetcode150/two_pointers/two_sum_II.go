package twopointers

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.
*/

/*
	sorted list, constant space -> this screams two pointers
	large values are on the right while smaller ones are on the left

	two pointers on left and right - if the sum of our two pointer values is too small,
	we need to move left pointer to go to a bigger value
	and vice versa for too big, right pointer moves left to go to a smaller value
*/

func twoSum(numbers []int, target int) []int {
	// create two pointers
	l, r := 0, len(numbers)-1

	// while loop until l passes r
	for l < r {
		num1 := numbers[l]
		num2 := numbers[r]

		if target > num1+num2 {
			// sum is too small
			l++
		} else if target < num1+num2 {
			// sum is too big
			r--
		} else {
			break
		}
	}

	return []int{l + 1, r + 1}
}
