/*

Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example :

nums = [1,3,3]

Output =

[

[],

[1],

[3],

[1,3],

[3,3],

[1,3,3]


*/

package main

import (
	"fmt"
	"sort"
)

// SubsetsWithDup generates all unique subsets of nums, handling duplicates
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var subsets func(index int, curr []int)
	subsets = func(index int, curr []int) {
		if index == len(nums) {
			// Create a new slice to avoid modifying the shared curr
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			res = append(res, tmp)
			return
		}

		// Include the current number
		curr = append(curr, nums[index])
		subsets(index+1, curr)
		curr = curr[:len(curr)-1] // Remove last element

		// Exclude the current number and skip all duplicates
		for index < len(nums)-1 && nums[index] == nums[index+1] {
			index++
		}
		subsets(index+1, curr)
	}

	subsets(0, []int{})
	return res
}

func subsets2() {
	nums := []int{1, 2, 2}
	result := subsetsWithDup(nums)
	fmt.Println("Subsets with duplicates:", result)
}
