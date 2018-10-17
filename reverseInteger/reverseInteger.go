/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range:
[−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

package main

import (
	"fmt"
)

func reverse(x int) int {
	// define border
	top_limit := 1<<31 - 1
	bottom_limit := -1 << 31

	// start to reverse int
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	// check whether overflow or not
	if result > top_limit || result < bottom_limit {
		result = 0
	}

	return result
}

func main() {
	x := 1534236469
	x_rev := 0
	x_rev = reverse(x)
	fmt.Println(x_rev)
}
