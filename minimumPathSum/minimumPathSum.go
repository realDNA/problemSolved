/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minPathSum(grid [][]int) int {
    gridRowLength := len(grid)
    gridColumnLength := len(grid[0])
    for m := 0; m < gridRowLength; m++ {
        for n := 0; n < gridColumnLength; n++ {
            if m == 0 && n-1 >= 0 {
                grid[m][n] += grid[m][n-1]
            } else if n == 0 && m-1 >= 0 {
                grid[m][n] += grid[m-1][n]
            } else if m-1 >= 0 && n-1 >= 0{
                grid[m][n] += min(grid[m-1][n], grid[m][n-1])
            }
        }
    }
    return grid[gridRowLength-1][gridColumnLength-1]
}
