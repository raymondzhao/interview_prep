package arrayshashing

/*
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.
*/

func topKFrequent(nums []int, k int) []int {
	// map number to count of number in nums
	freqMap := make(map[int]int)
	// need a list of list where i is the freq of the number.
	// the list inside is the numbers at that freq
	buckets := make([][]int, len(nums)+1)
	res := []int{}
	// iterate through making key: num, val: counts
	for _, val := range nums {
		count, exists := freqMap[val]
		if exists {
			freqMap[val] = count + 1
		} else {
			freqMap[val] = 1
		}
	}
	// iterate through map to fill the buckets with numbers
	for key, count := range freqMap {
		buckets[count] = append(buckets[count], key)
	}
	// start at largest index since that's where most freq nums are
	for i := len(buckets) - 1; i >= 0; i-- {
		// skip empty buckets
		if len(buckets[i]) == 0 {
			continue
		}
		// walk backwards in list to grab up to k values
		for j := len(buckets[i]) - 1; j >= 0; j-- {
			if len(res) < k {
				res = append(res, buckets[i][j])
			} else {
				return res
			}
		}
	}
	return res
}
