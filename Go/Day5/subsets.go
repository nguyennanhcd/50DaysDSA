/*

Given an integer array unique elements, return all possible subsets ( the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

*/

package main

import "fmt"

func subsets(nums []int) [][]int {
	output := [][]int{}

	var helper func(nums []int, i int, subset []int)
	helper = func(nums []int, i int, subset []int) {
		if i == len(nums) {
			// Create a new slice to avoid reference issues
			temp := make([]int, len(subset))
			copy(temp, subset)
			output = append(output, temp)
			return
		}

		// Exclude current element
		helper(nums, i+1, subset)

		// Include current element
		subset = append(subset, nums[i])
		helper(nums, i+1, subset)

		// Backtrack by removing the last element
		subset = subset[:len(subset)-1]
	}

	helper(nums, 0, []int{})
	return output
}

func subsetEx() {
	nums := []int{1, 8, 7}
	powerSet := subsets(nums)
	fmt.Println(powerSet)
}
