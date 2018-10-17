/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
    eIndex := len(nums) - 1
    var result [][]int
    fmt.Println(nums)
    for i := 0; i < eIndex-1; i++ {
    	if((i > 0) && (nums[i] == nums[i-1])) {
    		continue
    	}
    	fmt.Println("i = ",i)
    	go func (i int, eIndex int, result [][]int, nums []int) {
	    	target := -nums[i]
	    	startP := i+1
	    	endP := eIndex
	    	fmt.Println("target = ",target)
	    	for startP < endP {
	    		sum := nums[startP] + nums[endP]
	    		if(sum == target) {
	    			findOne := []int{nums[i],nums[startP],nums[endP]}
	    			result = append(result,findOne)
	    			startP++
	    			endP--
	    			for startP < endP && nums[startP] == nums[startP-1] {
						startP++
					}
					for startP < endP && nums[endP] == nums[endP+1] {
						endP--
					}
	    		} else if(sum > target){
	    			endP--
	    		} else if(sum < target) {
	    			startP++
	    		}	

	    	}
	    	fmt.Println("result routine = ", result)
	    }(i, eIndex, result, nums);
    }

	/*
    sort.Ints(nums)
    eIndex := len(nums) - 1
    var result [][]int
    fmt.Println(nums)
    for i := 0; i < eIndex-1; i++ {
    	if((i > 0) && (nums[i] == nums[i-1])) {
    		continue
    	}
    	target := -nums[i]
    	startP := i+1
    	endP := eIndex

    	for startP < endP {
    		sum := nums[startP] + nums[endP]
    		if(sum == target) {
    			findOne := []int{nums[i],nums[startP],nums[endP]}
    			result = append(result,findOne)
    			startP++
    			endP--
    			for startP < endP && nums[startP] == nums[startP-1] {
					startP++
				}
				for startP < endP && nums[endP] == nums[endP+1] {
					endP--
				}
    		} else if(sum > target){
    			endP--
    		} else if(sum < target) {
    			startP++
    		}	

    	}
    }
    */
	return result
}

func main() {
	a := []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}
	result := threeSum(a)
	fmt.Println("result = ",result)
}