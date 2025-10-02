package arrayshashing

import (
	"strconv"
	"strings"
)

/*
Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

Please implement encode and decode

Example 1:

Input: ["neet","code","love","you"]

Output:["neet","code","love","you"]
*/

// we can represent these strings using something like a number to represent length of string and
// delimiter like '#'
// example above could be 4#neet4#code4#love3#you

func Encode(strs []string) string {
	var b strings.Builder // using string builder for perf since strings are immutable in Go
	for _, str := range strs {
		length := strconv.Itoa(len(str))
		b.WriteString(length)
		b.WriteString("#")
		b.WriteString(str)
	}
	return b.String()
}

func Decode(s string) []string {
	curr := 0
	ret := []string{}
	for curr < len(s) {
		// find length prefix - while loop until we found delim
		j := curr
		for s[j] != '#' {
			j++
		}
		l, _ := strconv.Atoi(s[curr:j])

		// skip '#'
		j++

		// grab str using length
		str := s[j : j+l]
		ret = append(ret, str)

		// update curr
		curr = j + l

	}
	return ret
}
