/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

// "(" = 40
// ")" = 41
package main

import(
"fmt"
)

type stack []int

func (s stack) Push(v int) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, int) {
    l := len(s)
    if l == 0 {
    	return s, -1
    }
    return  s[:l-1], s[l-1]
}

func (s stack) Peek() int {
    l := len(s)
    if l == 0 {
    	return -1
    }
    return s[l-1]
}

func longestValidParentheses(s string) int {
	var sta stack
	var topIndex int
	maxParentheses := 0
	sta = sta.Push(-1)

	for i,v := range s {
		if v == 40 {
			sta = sta.Push(i)
		} else {
			topIndex = sta.Peek()
			if topIndex >= 0 && s[topIndex] == 40 {
				sta, _ = sta.Pop()
				length := i - sta.Peek()
				if  length > maxParentheses {
					maxParentheses = length
				}
			} else {
				sta = sta.Push(i)
			}
		}
	}
	return maxParentheses
}

func main() {
	test := ")()())()(()()()"
	fmt.Println(" test start = ", test)
	result := longestValidParentheses(test)
	fmt.Println("result = ", result)
}