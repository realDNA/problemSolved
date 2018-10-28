/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/

package main

import(
"fmt"
)

func isOtherQueenInRow(sRow string, col int, boardLen int) bool {
    for i := col; i < boardLen; i++ {
        if sRow[i] == 'Q' {
            return true
        }
    }
    return false
}

func isOtherQueenInDiagonal(board []string, row int, col int, boardLen int) bool {
    rowU, rowD := row, row
    colU, colD := col, col

    for i := rowU; i > 0; i-- {
        if colU < boardLen-1 {
            colU++
        } else {
            break
        }
        if board[i-1][colU] == 'Q' {
            return true
        }
    }
    
    for i := rowD; i < boardLen-1; i++ {
        if colD < boardLen-1 {
            colD++
        } else {
            break
        }
        if board[i+1][colD] == 'Q' {
            return true
        }
    }
    return false
}

func solveNQueensHelper(board []string, col int, boardLen int, allBoard *[][]string) bool {
    if col == -1 {
        copyBoard := make([]string, boardLen)
        copy(copyBoard, board)
        *allBoard = append(*allBoard, copyBoard)
        return true
    }

    for row := 0; row < boardLen; row++ {
        if !isOtherQueenInRow(board[row], col, boardLen) && 
           !isOtherQueenInDiagonal(board, row, col, boardLen) {
            board[row] = board[row][:col] + "Q" + board[row][col+1:]
            solveNQueensHelper(board, col-1, boardLen, allBoard)
            board[row] = board[row][:col] + "." + board[row][col+1:]
        }
    }

    return false
}

func solveNQueens(n int) [][]string {

    var board []string
    var allBoard[][]string

    stringTemplate := ""

    for s := 0; s < n; s++ {
        stringTemplate += "." 
    }
    for b := 0; b < n; b++ {
        board = append(board, stringTemplate)
    }

    col := n-1 // collumn, will decrease, begin from last collumn
    boardLen := n // fixed

    solveNQueensHelper(board, col, boardLen, &allBoard)

    return allBoard
}

func main() {
    n := 6
    board := solveNQueens(n)
    fmt.Println(board)
}
