/*

Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.


Example:

nums = [3,3,2]

Output :

[[3,3,2] , [3,2,3], [2,3,3] ]

*/

package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int

	var helper func(int)
	helper = func(i int) {
		// base case
		if i == len(nums)-1 {
			// copy slice để tránh thay đổi do backtracking
			perm := make([]int, len(nums))
			copy(perm, nums)
			res = append(res, perm)
			return
		}

		// avoid duplicates by using a map
		used := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			if !used[nums[j]] {
				used[nums[j]] = true
				nums[i], nums[j] = nums[j], nums[i]
				helper(i + 1)
				nums[i], nums[j] = nums[j], nums[i] // backtrack
			}
		}
	}

	helper(0)
	return res
}

func Permutation2() {
	nums := []int{3, 3, 2}
	result := permuteUnique(nums)
	for _, perm := range result {
		fmt.Println(perm)
	}
}
