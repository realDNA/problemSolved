/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given 'abcabcbb', the answer is 'abc', which the length is 3.

Given 'bbbbb', the answer is 'b', with the length of 1.

Given 'pwwkew', the answer is 'wke', with the length of 3. Note that the answer must be a substring, 'pwke' is a subsequence and not a substring.
*/

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	count := 0
	start := 0
	currMaxcount := 0
	mapStore := make(map[rune]int)

	for pos, char := range s {
		// create map to store character as key:value
		if _, exist := mapStore[char]; exist {

			// "start index" to record the current start character
			if mapStore[char] >= start {
				start = mapStore[char] + 1
			}

			// the total number of currnet non repeat character
			count = pos - start + 1
		} else {

			// increase number of non repeat character
			count++
		}

		// record current max number of non repeat character
		if count > currMaxcount {
			currMaxcount = count
		}

		// store character as key value
		mapStore[char] = pos
	}
	return currMaxcount
}

func main() {
	count := 0
	p := fmt.Println
	test := "abcabcbb"
	count = lengthOfLongestSubstring(test)
	p(count)
}
