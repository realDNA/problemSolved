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

func isOtherQueenInRow(sRow []string, col int, boardLen int) bool {
    for i := col; i < boardLen; i++ {
        if sRow[i] == "Q" {
            return true
        }
    }
    return false
}

func isOtherQueenInDiagonal(board [][]string, row int, col int, boardLen int) bool {
    rowU, rowD := row, row
    colU, colD := col, col

    //fmt.Println("upper area ------------ ")
    for i := rowU; i > 0; i-- {
        if colU < boardLen-1 {
            colU++
        } else {
            break
        }
        if board[i-1][colU] == "Q" {
            return true
        }
        //fmt.Println("rowU = ", i-1)
        //fmt.Println("colU = ", colU)
    }
    //fmt.Println("lower area ------------")
    for i := rowD; i < boardLen-1; i++ {
        if colD < boardLen-1 {
            colD++
        } else {
            break
        }
        if board[i+1][colD] == "Q" {
            return true
        }
        //fmt.Println("rowD = ", i+1)
        //fmt.Println("colD = ", colD)
    }
    return false
}

func solveNQueensHelper(board [][]string, col int, boardLen int) bool {
    if col == -1 {
        return true
    }

    for row := 0; row < boardLen; row++ {
        if !isOtherQueenInRow(board[row], col, boardLen) && 
           !isOtherQueenInDiagonal(board, row, col, boardLen) {
            board[row][col] = "Q"
            isSuccess := solveNQueensHelper(board, col-1, boardLen)
            if isSuccess {
                return true   
            } else {
                board[row][col] = "."
            }
        }
    }

    return false
}

func solveNQueens(n int) [][]string {

    // make queen chess board
    board := make([][]string, n, n)

    for b := 0; b < n; b++ {
        board[b] = make([]string, n)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }

    fmt.Println(board)

    //isOtherQueenInDiagonal(board, 0, 0, n)

    col := n-1 // collumn, will decrease, begin from last collumn
    boardLen := n // fixed

    solveNQueensHelper(board, col, boardLen)

    return board
}

func main() {
    n := 6
    board := solveNQueens(n)
    fmt.Println(board)
}
