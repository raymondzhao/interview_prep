package slidingwindow

/*
Given a string s, find the length of the longest substring without duplicate characters.

A substring is a contiguous sequence of characters within a string.
*/

func lengthOfLongestSubstring(s string) int {
	// sliding window problem - but we need to update keep track of our last seen duplicate index
	// to move the left ptr so we can update the window
	left, right, maxLen := 0, 0, 0
	charMap := make(map[byte]int) // need to store char to the last seen index

	for right < len(s) {
		rightChar := s[right]
		val, exists := charMap[rightChar]
		// update the left pointer if we find a duplicate but only if the
		// update maintains the window
		if exists && val >= left {
			left = val + 1
		}
		// update last seen index
		charMap[rightChar] = right
		// make sure we keep track of the max length
		maxLen = max(maxLen, right-left+1)
		// move the right pointer forward until we hit end of the string
		right++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
