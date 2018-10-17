/*
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

package main

import (
	"fmt"
)

func numMin(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func maxArea(height []int) int {
	startX := 0
	startY := 0
	endX := len(height) - 1
	endY :=0
	width := 0
	dimension := 0
	maxDimension := 0

	for startX != endX {
		startY = height[startX]
		endY = height[endX]
		width = endX - startX
		dimension = width*numMin(startY, endY)

		if dimension > maxDimension {
			maxDimension = dimension
		}

		if startY > endY {
			endX -= 1
		} else {
			startX += 1
		}
	}

	return maxDimension    
}


func main() {
	height := []int{9,4,3,4,8,6}
	maxDimension := maxArea(height)
	fmt.Println(maxDimension)
}