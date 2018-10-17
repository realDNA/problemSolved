/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

func twoSum(numbs []int, target int) []int {
	mapStore := make(map[int]int)
	for index1, value1 := range numbs {
		findKey := target - value1
		if _, exist := mapStore[findKey]; exist {
			return []int{mapStore[findKey], index1}
		} else {
			mapStore[value1] = index1
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := []int{}

	result = twoSum(nums, target)
	fmt.Print(result)
}
