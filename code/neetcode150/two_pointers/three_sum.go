package twopointers

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.
*/
import (
	"sort"
)

func threeSum(nums []int) [][]int {
	// need to sort the array first
	// we can convert the equation to be -nums[i] = nums[j] + nums[k]
	// j and k can be used as two pointers - this only works when the array is sorted

	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ { // len(nums) - 2 bc we have j and k
		// check for duplicate i's since if we don't we get the same triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// make two pointers
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// move j pointer so it's not a duplicate
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// move k pointers so it's not a duplicate
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				// if sum is greater, we need to reduce k since k represents bigger values
				k--
			} else if sum < 0 {
				// if sum is less, we need to increase j since j represents smaller values
				j++
			}
		}
	}
	return res
}
