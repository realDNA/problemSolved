/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

package main

import(
"fmt"
"sort"
)

func findSmallestLargerNum(pivotVal int, pivotPoint int, numsLen int, nums []int) int {
	var targetPoint int

	start := pivotPoint
	end := numsLen - 1
	middle := (start+end) / 2
	
	for (middle != start) && (middle != end) {
		if nums[middle] > pivotVal {
			end = middle
		} else {
			start = middle
		}
		middle = (start+end) / 2
	} 

	if nums[start] > pivotVal {
		targetPoint = start
	} else {
		targetPoint = end
	}

	return targetPoint
} 

func nextPermutation(nums []int)  {
    numsLen := len(nums)
    reverseAll := true
    var pivotVal int
    var pivotPoint int 

    for i := numsLen-1; i > 0; i-- {
    	if nums[i] > nums[i-1] {
    		pivotVal = nums[i-1]
    		pivotPoint = i-1
    		reverseAll = false
    		break
    	}
    }

    if reverseAll {
    	sort.Ints(nums)
    } else {
        sort.Ints(nums[pivotPoint+1:])
        fmt.Println(nums)
        targetPoint := findSmallestLargerNum(pivotVal, pivotPoint, numsLen, nums)
        temp := pivotVal
        nums[pivotPoint] = nums[targetPoint]
        nums[targetPoint] = temp
    }
}

func main() {
	test := []int{5,4,7,5,3,2}
	fmt.Println("start run test = ", test)
	nextPermutation(test)
}