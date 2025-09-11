package arrayshashing

/*
Given an array of strings strs, group the together. You can return the answer in any order.
*/

// need to create a key for any anagram
// one key is to sort the string by rune/char so all anagrams wil point to the same key

// another optimization: use a string of 26 len with index being char's count
// that could be our key and would perform better since it would be O(n) instead of O(n log n)

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// make a hashmap - k: anagram key, v: list of the anagrams
	m := make(map[string][]string)
	for _, v := range strs {
		// sort string to create key
		runes := []rune(v)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sorted := string(runes)
		// add anagram to hashmap using new key
		m[sorted] = append(m[sorted], v)
	}

	// iterate through hashmap to get groups
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
