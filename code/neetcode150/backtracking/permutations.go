package backtracking

/*
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.
*/

func permute(nums []int) [][]int {
	var all [][]int
	used := make([]bool, len(nums))
	var path []int
	backtrack(nums, &all, path, used)
	return all
}

func backtrack(nums []int, all *[][]int, path []int, used []bool) {
	// if index is at last value in nums, we're done
	if len(path) == len(nums) {
		c := make([]int, len(path))
		copy(c, path)
		*all = append(*all, c)
		return
	}

	// need to iterate over every other value in nums
	for i := 0; i < len(nums); i++ {
		// we've already used that value so skip
		if used[i] {
			continue
		}

		// include current num
		path = append(path, nums[i])
		used[i] = true

		backtrack(nums, all, path, used)

		// backtrack when done
		path = path[:len(path)-1]
		used[i] = false
	}
}
