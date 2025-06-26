/*
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row.
Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01,
and each occurrence of 1 with 10.  For example, for n = 3, the 1st row is 0, the 2nd row is 01,
and the 3rd row is 0110. Given two integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.
*/

// the first half of the current row equal to the previous row
// the second half of the current row equal to the inversion of the previous row and also r

package main

import (
	"fmt"
	"math"
)

func kthRow(n, k int) int {
	if n == 1 {
		return 0
	}

	length := int(math.Pow(2, float64(n-1)))
	mid := length / 2
	if k <= mid {
		return kthRow(n-1, k)
	} else {
		return 1 - kthRow(n-1, k-mid)
	}
}

func Ex1() {
	n := 5
	k := 3
	result := kthRow(n, k)
	fmt.Println(result)
}
