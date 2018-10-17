/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/

package main

import(
"fmt"
)

func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    numsLength := len(nums)
    first := nums[0]
    last := nums[numsLength-1]
    startP := 0
    endP := numsLength-1    
    var middle int

    if numsLength == 1 {
        if target == nums[0] {
            return 0
        } else {
            return -1
        }
    }

    for startP != endP {
        middle = (startP + endP)/2
        if startP == endP-1 {
            if target == nums[startP] {
                return startP
            } else if target == nums[endP] {
                return endP
            } else {
                return -1
            }
        }
        if nums[middle] >= first {
            if target >= first {
                if target > nums[middle] {
                    startP = middle
                } else if target < nums[middle] {
                    endP = middle
                } else if target == nums[middle] {
                    return middle
                }
            } else if target <= last {
                startP = middle
            } else if target<first && target>last {
                return -1
            }
        } else if nums[middle] <= last {
            if target >= first {
                endP = middle
            } else if target <= last {
                if target < nums[middle] {
                    endP = middle 
                } else if target > nums[middle] {
                    startP = middle
                } else if target == nums[middle] {
                    return middle
                }
            } else if target<first && target>last {
                return -1
            }
        }        
    }
    return -1
}

func main() {
    nums := []int{1}
    target := 1
    result := search(nums, target)
    fmt.Println("result = ", result)
}
