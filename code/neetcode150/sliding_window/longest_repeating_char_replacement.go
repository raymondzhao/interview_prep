package slidingwindow

/*
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
*/

/*
	this is a sliding window problem

	since we have k times we can replace chars,
	window size - most freq char count = the number of chars we need to replace to make it all the same
	that value is compared to k which is how many chars we CAN replace

	if it's bigger than k, we need to make the window smaller since we need more replacements than alloted

	at the end of the day we are looking for the biggest window that has the most same chars after k replacements
*/

func characterReplacement(s string, k int) int {
	l, maxFreq, ans := 0, 0, 0
	m := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		// add new char to map
		m[s[r]]++
		// need to check max since new char added to window might be max
		maxFreq = max(maxFreq, m[s[r]])

		window := r - l + 1
		// check how many replacements we need and compare with k
		for window-maxFreq > k {
			// shrink window by moving left ptr
			// update freq map
			m[s[l]]--
			l++
			window--
		}

		// compare our max window size
		ans = max(window, ans)
	}
	return ans
}

/*
helper function in a prev exercise
	func max(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
*/
