package slidingwindow

/*
Given two strings s and t of lengths m and n respectively,
return the minimum window of s such that every character
in t (including duplicates) is included in the window.

If there is no such substring, return the empty string "".

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
*/

func minWindow(s string, t string) string {
	// need to get chars and their counts in t
	required := 0 // track all unique characters
	tChars := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if c, exists := tChars[t[i]]; exists {
			tChars[t[i]] = c + 1
		} else {
			tChars[t[i]] = 1
			required += 1
		}
	}

	// need to create a map showing chars and counts for the window
	window := make(map[byte]int)
	formed := 0                       // track how many t's chars are in window
	bestLen, bestIndex := len(s)+1, 0 // track the best substring's index and the length to calc result
	left, right := 0, 0               // our pointers
	for right < len(s) {
		// add letter to window
		window[s[right]] += 1

		// check if window satisfies t's chars - increment formed
		if window[s[right]] == tChars[s[right]] {
			formed += 1
		}

		// if window contains all chars in t
		if formed == required {
			// shrink sliding window until we don't match string t
			for formed == required {
				if right-left+1 < bestLen {
					bestIndex = left
					bestLen = right - left + 1
				}
				window[s[left]] -= 1
				// if we shrink window and result is less than chars in T, we stop moving left ptr
				if window[s[left]] < tChars[s[left]] {
					formed -= 1
				}
				left += 1
			}
		}
		// move right pointer
		right += 1
	}
	// if we dont see any substring, return empty
	if bestLen == len(s)+1 {
		return ""
	}
	// return answer is substring
	return s[bestIndex : bestIndex+bestLen]
}
