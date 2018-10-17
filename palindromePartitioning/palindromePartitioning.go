/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

package main

import(
"fmt"
)

func isPalindrome(s string) bool {
    sLength := len(s)
    for i:=0; i<sLength/2; i++ {
        if s[i] != s[sLength-1-i] {
            return false
        }
    }
    return true
}

func partitionHelper(s string, row []string, result *[][]string ) {
    sLen := len(s)
    if sLen == 0 {
        copyRow := make([]string, len(row))
        copy(copyRow, row) // prevent use same memory for row
        *result = append(*result, copyRow)
        return
    }
    for i:=1; i<=sLen; i++ {
        s1 := s[0:i]
        if !isPalindrome(s1) {
            continue
        }
        s2 := s[i:sLen]
        partitionHelper(s2, append(row,s1), result )    
    }
}

func partition(s string) [][]string  {
    var result [][]string
    var row []string
    partitionHelper(s,row,&result)
    return result
}

func main(){
    testS := "ababbbabbaba"
    result1 := partition(testS)
    fmt.Println("result = ",result1)
}
