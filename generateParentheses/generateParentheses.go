/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

import(
"fmt"
)

func generateParenthesisHelper(n int, s string, left_n int, right_n int, length int, result *[]string) {
	if length == 2*n {
		*result = append(*result, s)
	}

	if left_n < n {
		generateParenthesisHelper(n, s+"(", left_n+1, right_n, length+1, result)
	}
	if left_n > right_n && right_n < n {
		generateParenthesisHelper(n, s+")", left_n, right_n+1, length+1, result)
	}
}

func generateParenthesis(n int) []string {
 	var result []string
	generateParenthesisHelper(n, "", 0, 0, 0, &result)
	return result
}

func main() {
	result := generateParenthesis(27)
	fmt.Println(result)
}