package arrayshashing

/*
Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

// every answer[i] is the product of every number before i and every product after i
// [1,2,3,4] = answer[3] = 1 x 2 x 4 aka left (1,2) x right (4)
// need to do two passes - each pass keeping track of a running product as we go from left to right
// and vice versa
func productExceptSelf(nums []int) []int {
	answers := make([]int, len(nums))

	// left pass - update the running product with the current index for the next calc
	product := 1 // start at 1
	for i := 0; i < len(nums); i++ {
		answers[i] = product
		product *= nums[i]
	}

	// right pass - same thing but we need to multiple the answers itself
	// since it has our left pass
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answers[i] *= product
		product *= nums[i]
	}

	return answers
}
