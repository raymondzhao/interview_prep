package arrayshashing

/*
Given an unsorted array of integers nums,
return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/

// need to use a hashmap to store all the values

// since its unsorted we'll need to find a start num
// and then follow the sequence using the hashmap
func longestConsecutive(nums []int) int {
	// we add all values to a hashmap
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	maxSeq := 0
	for v := range m {
		// we need to check if this number is a start
		// if it is we dont care/skip
		if m[v-1] {
			continue
		}
		// found a start
		num := v + 1
		seq := 1
		// follow the sequence in the hashmap and increase seq
		for m[num] {
			seq++
			num++
		}
		// keep track of the maxSeq
		maxSeq = max(maxSeq, seq)
	}

	return maxSeq
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
