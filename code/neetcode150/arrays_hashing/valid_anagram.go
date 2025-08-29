package arrayshashing

/*
Given two strings s and t, return true if t is an anagram
of s, and false otherwise.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false
*/

func isAnagram(s string, t string) bool {
	// check lengths for easy check
	if len(s) != len(t) {
		return false
	}

	// build freq map of letters to how many times they are used
	sMap := make(map[rune]int)

	for _, r := range s {
		if count, exists := sMap[r]; exists {
			sMap[r] = count + 1
		} else {
			sMap[r] = 1
		}
	}

	// decrement the counts if exists and if its negative, return false
	// also return false if value is not in the freq map
	for _, r := range t {
		count, exists := sMap[r]
		if !exists {
			return false
		}
		sMap[r] = count - 1
		if count-1 < 0 {
			return false
		}
	}

	return true
}
