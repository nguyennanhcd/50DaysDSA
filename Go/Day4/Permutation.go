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

func backtrack(path []int, used map[int]bool) {
	if len(path) == 3 {
		fmt.Println(path)
		return
	}
	for i := 1; i <= 3; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, i)
		backtrack(path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func Permutation() {
	fmt.Println("Backtracking:")
	backtrack([]int{}, map[int]bool{})
}
