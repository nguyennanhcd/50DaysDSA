/* An array is monotonic if it is either monotone increasing or monotone decreasing.
An array is monotone increasing if all its elements form left to right are non-decreasing.
Given an integer array return true if the given array is monotonic, or false otherwise
*/

package main

import "fmt"

func checkMono(num []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := range len(num) - 1 {
		if num[i] < num[i+1] {
			isDecreasing = false
		}
		if num[i] > num[i+1] {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

func Ex2() {
	monoIncr := []int{1, 2, 3, 4, 5, 6}
	monoDecr := []int{9, 8, 7, 6, 5, 4}
	monoNone := []int{1, 3, 5, 2, 9, 4, 11, 6}

	fmt.Println((checkMono(monoIncr)))
	fmt.Println((checkMono(monoDecr)))
	fmt.Println((checkMono(monoNone)))
}
