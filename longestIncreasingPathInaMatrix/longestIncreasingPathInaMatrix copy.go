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

func dfs(i int,j int,matrix [][]int, path_length int, 
         record [][]int, dir [][]int, dir_x int, dir_y int,
         row_length int, column_length int, first_time int) int {
    fmt.Println("i = ", i)
    fmt.Println("j = ", j)
    fmt.Println("record = ", record)
    if first_time == 0 {
        if i < row_length &&
           j < column_length &&
           i >= 0 && j >= 0 &&
           record[i][j] != 1 &&
           matrix[i][j] > matrix[i-dir_x][j-dir_y] {
           record[i][j] = 1
           path_length += 1
        } else {
            return path_length
        }
    } else {
        record[i][j] = 1
    }
    max_length := 0
    for d:=0; d<len(dir); d++ {
        dir_x_new := dir[d][0]
        dir_y_new := dir[d][1]
        total_length := dfs(i+dir_x_new, j+dir_y_new, matrix, path_length,
                            record, dir, dir_x_new, dir_y_new, row_length,
                            column_length, 0)
        fmt.Println("total_lengeth = ",total_length)
        max_length = max(max_length, total_length)
    }

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
    dir :=[][]int{
        {0,1},
        {0,-1},
        {1,0},
        {-1,0},
    }
    max_length := 0
    for i:=0; i<row_length; i++ {
        for j:=0; j<column_length; j++ {
           path_length := 1
           for i := 0; i < row_length; i++ {
               record[i] = make([]int, column_length)
           }
           fmt.Println("change next -------------") 
           fmt.Println("intitial record === ", record)   
           total_length := dfs(i, j, matrix, path_length, record,dir, 
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
