/* You are given an array of integers in which each subsequent value is not less than the previous
value. Write a function that takes this array as an input and returns a new array with the squares
of each number sorted in ascending order.
*/

package main

import (
	"fmt"
)

func squareSort(num []int) []int {
	n := len(num)
	left, right := 0, n-1
	result := make([]int, n)
	pos := n - 1

	for left <= right {
		leftValue := num[left] * num[left]
		rightValue := num[right] * num[right]

		if leftValue > rightValue {
			result[pos] = leftValue
			left++
		} else {
			result[pos] = rightValue
			right--
		}
		pos--
	}
	return result
}

func Ex1() {
	num := []int{-5, 3, 4, 6}
	result := squareSort(num)
	fmt.Println(result)
}
