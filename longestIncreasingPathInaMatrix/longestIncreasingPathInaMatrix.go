/*
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums = 
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
] 
Output: 4 
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: nums = 
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
] 
Output: 4 
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
*/

package main

import(
"fmt"
)

func max(i int, j int) int {
    if j > i {
        return j
    } 
    return i
}

func dfs(i int,j int,matrix [][]int, 
         record, dir [][]int, dir_x int, dir_y int,
         row_length int, column_length int, first_time int) int {
    if first_time == 0 {
        if i < row_length &&
           j < column_length &&
           i >= 0 && j >= 0 &&
           matrix[i][j] > matrix[i-dir_x][j-dir_y] {
           if record[i][j] != 0 {
                return record[i][j]
           }    
        } else {
            return 0
        }
    }
    max_length := 0
    total_length := 0
    for d:=0; d<len(dir); d++ {
        dir_x_new := dir[d][0]
        dir_y_new := dir[d][1]
        total_length = 1 + dfs(i+dir_x_new, j+dir_y_new, matrix,
                               record, dir, dir_x_new, dir_y_new, row_length,
                               column_length, 0)
        max_length = max(max_length, total_length)
        total_length = 0
    }
    record[i][j] = max_length
    return max_length
}

func longestIncreasingPath(matrix [][]int) int {
    row_length := len(matrix)
    if row_length == 0 {
        return 0
    }
    column_length := len(matrix[0])
    if column_length == 0 {
        return 0
    }
    record := make([][]int, row_length)
    for i := 0; i < row_length; i++ {
        record[i] = make([]int, column_length)
    }
    dir :=[][]int{
        {0,1},
        {0,-1},
        {1,0},
        {-1,0},
    }
    max_length := 0
    for i:=0; i<row_length; i++ {
        for j:=0; j<column_length; j++ {
           total_length := dfs(i, j, matrix, record, dir, 
                               0, 0, row_length, column_length, 1)
           max_length = max(max_length, total_length)
        }
    }
    return max_length
}

func main(){
    a := [][]int{
        {7,6,1,1},
        {2,7,6,0},
        {1,3,5,1},
        {6,6,3,2},
    }
    max_length := longestIncreasingPath(a)
    fmt.Println("max length = ", max_length)
}
