/*

Coding Exercise ( Permutations)
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

nums = [1,4]

Output :[[1,4],[4,1]]

Example 2:

nums = [1,4,5]

Output :[[1,4,5],[1,5,4],[4,1,5],[4,5,1],[5,1,4],[5,4,1]]

*/

package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	result := [][]int{}
	used := make(map[int]bool, len(nums))
	path := []int{}

	var backtrack func()
	backtrack = func() {
		// Base case: if path length equals input array length, add to result
		if len(path) == len(nums) {
			// Create a copy of path to avoid reference issues
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		// Try each number in nums
		for _, num := range nums {
			if used[num] {
				continue
			}
			// Mark number as used and add to path
			used[num] = true
			path = append(path, num)

			// Recursive call
			backtrack()

			// Backtrack: remove last number and mark as unused
			path = path[:len(path)-1]
			used[num] = false
		}
	}

	backtrack()
	return result
}

func Permutation() {
	// Test cases
	nums1 := []int{1, 4}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", permute(nums1))

	nums2 := []int{1, 4, 5}
	fmt.Println("\nInput:", nums2)
	fmt.Println("Output:", permute(nums2))
}
