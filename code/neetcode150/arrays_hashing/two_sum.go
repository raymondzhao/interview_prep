package arrayshashing

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution,
and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

*/

func twoSum(nums []int, target int) []int {
	n := make(map[int]int) // create map to store value and index in nums
	for i, value := range nums {
		complement := target - value
		// check map - if exists we have a match
		if v, exists := n[complement]; exists {
			return []int{v, i}
		}
		// if no match, update our map
		n[value] = i
	}
	return nil
}
