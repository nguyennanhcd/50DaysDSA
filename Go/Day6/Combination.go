/*
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
*/

package main

import "fmt"

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	result := [][]int{}
	var helper func(start int, path []int)
	helper = func(start int, path []int) {
		if len(path) == k {
			combination := make([]int, k)
			copy(combination, path)
			result = append(result, combination)
			return
		}
		for i := start; i <= n; i++ {
			helper(i+1, append(path, i))
		}
	}
	helper(1, []int{})
	return result
}

func Combination() {

	n := 4
	k := 2
	combinations := combine(n, k)
	fmt.Println("Combinations of", k, "numbers from 1 to", n, ":", combinations)

}
