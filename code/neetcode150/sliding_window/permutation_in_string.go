package slidingwindow

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.

A permutation is a rearranging of characters
*/

func checkInclusion(s1 string, s2 string) bool {
	// need to grab the runes from s1 into a freq map
	m := make(map[rune]int)
	// map holds info about what we need, dont need, what the window is doing
	for _, rune := range s1 {
		m[rune]++
	}
	runes2 := []rune(s2)
	l := 0
	missing := len(s1) // keep track of missing characters
	for r := 0; r < len(runes2); r++ {
		// add right character (r)
		if m[runes2[r]] > 0 {
			// found a char
			missing--
		}
		m[runes2[r]]--

		window := r - l + 1
		// shrink if window is too large
		if window > len(s1) {
			// char was needed/part of the map
			// negative means we had extra copies of char
			if m[runes2[l]] >= 0 {
				missing++
			}
			m[runes2[l]]++
			l++
		}

		if missing == 0 {
			return true
		}

	}
	return false
}
