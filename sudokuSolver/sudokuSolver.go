package main

import (
	"fmt"
)

const sudokuMaxNum = 9

func getColumn(board [][]byte, column int) []byte {
	s := make([]byte, sudokuMaxNum)
	for i, v := range board {
		s[i] = v[column]
	}
	return s
}

func isInArray(target byte, byteArray []byte) int {
	for i, v := range byteArray {
		if target == v {
			return i
		}
	}
	return -1
}

func countRemainNumbers(sudokuLine []byte) []byte {
	nums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, v := range sudokuLine {
		matchIndex := isInArray(v, nums)
		if matchIndex >= 0 {
			nums = append(nums[:matchIndex], nums[matchIndex+1:]...)
		}
	}
	return nums
}

func collectNumbersApplyRow(sudokuRow []byte) []byte {
	return countRemainNumbers(sudokuRow)
}

func collectNumbersApplycolumn(sudokucolumn []byte) []byte {
	return countRemainNumbers(sudokucolumn)
}

func collectNumbersApplyGrid(board [][]byte, row int, column int) []byte {
	gridRowNum := row / 3
	gridColumnNum := column / 3
	gridRow := make([]byte, sudokuMaxNum)
	n := 0
	for i := gridRowNum * 3; i <= gridRowNum*3+2; i++ {
		for j := gridColumnNum * 3; j <= gridColumnNum*3+2; j++ {
			gridRow[n] = board[i][j]
			n += 1
		}
	}
	return countRemainNumbers(gridRow)
}

func numCanFillInCell(rowNums []byte, columnNums []byte, gridNums []byte) []byte {
	var setNums []byte
	isInColumn := false
	isInGrid := false

	for _, vRow := range rowNums {
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
	return setNums
}

func solveSudokuHelper(board [][]byte, directlyFillMultipleCell bool) bool {
	minNumsLen := sudokuMaxNum
	minI := sudokuMaxNum
	minJ := sudokuMaxNum
	isAnyEmpty := false
	var minCellNums []byte
	var cellNumsisOnePointArray [][]int
	var cellNumsisOneNumArray [][]byte

	for i := 0; i < sudokuMaxNum; i++ {
		for j := 0; j < sudokuMaxNum; j++ {
			if board[i][j] == '.' {
				isAnyEmpty = true
				rowNums := collectNumbersApplyRow(board[i])
				columnNums := collectNumbersApplycolumn(getColumn(board, j))
				gridNums := collectNumbersApplyGrid(board, i, j)
				cellNums := numCanFillInCell(rowNums, columnNums, gridNums)
				cellNumsLen := len(cellNums)

				if cellNumsLen == 0 {
					return false
				}

				if cellNumsLen < minNumsLen {
					minNumsLen = cellNumsLen
					minI = i
					minJ = j
					minCellNums = cellNums
				}

				if cellNumsLen == 1 && directlyFillMultipleCell {
					cellNumsisOnePoint := []int{i, j}
					cellNumsisOnePointArray = append(cellNumsisOnePointArray, cellNumsisOnePoint)
					cellNumsisOneNumArray = append(cellNumsisOneNumArray, cellNums)
				}
			}
		}
	}

	if !isAnyEmpty {
		return true
	}

	cellNumsisOnePointArrayLen := len(cellNumsisOnePointArray)
	if cellNumsisOnePointArrayLen != 0 {
		for i := 0; i < cellNumsisOnePointArrayLen; i++ {
			rowNum := cellNumsisOnePointArray[i][0]
			rowColumn := cellNumsisOnePointArray[i][1]
			board[rowNum][rowColumn] = cellNumsisOneNumArray[i][0]
		}
		// just one way, no need to check
		solveSudokuHelper(board, true)
	} else {
		for i, v := range minCellNums {
			board[minI][minJ] = v

			if i == len(minCellNums) {
				directlyFillMultipleCell = true
			} else {
				directlyFillMultipleCell = false
			}

			isWork := solveSudokuHelper(board, directlyFillMultipleCell)
			if isWork == false {
				board[minI][minJ] = '.'
			} else {
				return true
			}
		}
	}

	return false
}

func solveSudoku(board [][]byte) {
	directlyFillMultipleCell := true
	solveSudokuHelper(board, directlyFillMultipleCell)
	fmt.Println(board)
}

func main() {
	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	solveSudoku(board)
}
