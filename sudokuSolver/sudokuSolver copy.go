package main

import(
"fmt"
)

func getColumn(board [][]byte, column int) []byte {
    s := make([]byte, 9)
    for i,v := range board {
        s[i] = v[column]
    }
    return s
}

func isInArray(target byte, byteArray []byte) int {
    for i,v := range byteArray {
        if target == v {
            return i
        }
    }
    return -1
}

func countRemainNumbers(sudokuLine []byte) []byte {
    nums := []byte{'1','2','3','4','5','6','7','8','9'}
    for _,v := range sudokuLine {
        matchIndex := isInArray(v, nums)
        if matchIndex >= 0 { 
            nums = append(nums[:matchIndex], nums[matchIndex+1:]...)
        }
    }
    //fmt.Println("final nums = ", nums)
    return nums
}

func collectNumbersApplyRow(sudokuRow []byte) []byte {
    //fmt.Println("sudokuRow = ",sudokuRow)
    return countRemainNumbers(sudokuRow)
}

func collectNumbersApplycolumn(sudokucolumn []byte) []byte {
    //fmt.Println("sudokucolumn = ",sudokucolumn)
    return countRemainNumbers(sudokucolumn)
}

func collectNumbersApplyGrid(board [][]byte, row int, column int) []byte {
    gridRowNum := row/3
    gridColumnNum := column/3
    gridRow := make([]byte, 9)
    n := 0
    for i := gridRowNum*3; i <= gridRowNum*3+2; i++ {
        for j:= gridColumnNum*3; j <= gridColumnNum*3+2; j++ {
           gridRow[n] = board[i][j]
           n += 1
        }
    }
    //fmt.Print("gridRow = ",gridRow)
    return countRemainNumbers(gridRow)
}

func numCanFillInCell(rowNums []byte, columnNums []byte, gridNums[]byte) []byte {
    //fmt.Println("rowNums = ",rowNums)
    //fmt.Println("columnNums = ",columnNums)
    //fmt.Println("gridNums = ",gridNums)
    var setNums []byte
    isInColumn := false
    isInGrid := false
    
    for _,vRow := range rowNums {
        for _, vColumn := range columnNums {
            if vRow == vColumn {
                isInColumn = true
            }
        } 
        for _, vGrid := range gridNums {
            if vRow == vGrid {
                isInGrid = true
            }
        }
        if isInColumn && isInGrid {
            setNums = append(setNums, vRow)
        }
        isInColumn = false
        isInGrid = false
    }
    //fmt.Println("setNums = ",setNums)
    return setNums
}

func solveSudokuHelper(board [][]byte) bool {
    minNumsLen := 10
    minI := 10
    minJ := 10
    var minCellNums []byte

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                rowNums := collectNumbersApplyRow(board[i])
                columnNums := collectNumbersApplycolumn(getColumn(board, j))
                gridNums := collectNumbersApplyGrid(board, i, j)
                cellNums := numCanFillInCell(rowNums, columnNums, gridNums)
                cellNumsLen := len(cellNums)
                if cellNumsLen == 0 {
                    fmt.Printf("(%d, %d) cannot be applied number\n", i ,j )
                    return false
                }
                if cellNumsLen < minNumsLen {
                    minNumsLen = cellNumsLen
                    minI = i
                    minJ = j
                    minCellNums = cellNums
                }
                fmt.Printf("(%d, %d) = %d\n",i, j, cellNumsLen)
            }
        }
    }
    fmt.Printf("(%d, %d ) has min len = %d, %v\n", minI, minJ, minNumsLen, minCellNums)
    if minI > 8 || minJ > 8 {
        return true
    }
    for _,v := range minCellNums {
        board[minI][minJ] = v
        isWork := solveSudokuHelper(board)
        if isWork == false {
            fmt.Println("cannot work")
            board[minI][minJ] = '.'
        } else {
            fmt.Println("can work")
            return true
        }
    }
    return false
}

func solveSudoku(board [][]byte) {
    solveSudokuHelper(board)
    fmt.Println(board)
}

func main() {
    board := [][]byte {
        {'5','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'},
    }
    solveSudoku(board)
}
