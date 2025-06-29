/*
Power Sum - Let’s define a peculiar type of array in which each element is either an integer or another peculiar array.
Assume that a peculiar array is never empty.
Write a function that will take a peculiar array as its input and find the sum of its elements.
If an array is an element in the peculiar array
you have to convert it to it’s equivalent value so that you can sum it with the other elements.
Equivalent value of an array is the sum of its elements raised to the number which represents how far nested it is.
For e.g. [2,3[4,1,2]] = 2+3+ (4+1+2)^2

another example - [1,2,[7,[3,4],2]] = 1 + 2 +( 7+(3+4)^3+2)^2

*/

package main

import (
	"fmt"
	"math"
)

// powerSum recursively sums the elements of nested arrays, applying exponentiation based on depth.
func powerSum(array []any, power int) float64 {
	sum := 0.0
	for _, v := range array {
		switch val := v.(type) {
		case float64: // base case: number
			sum += val
		case int: // in case elements are int instead of float64
			sum += float64(val)
		case []any: // recursive case: nested list
			sum += powerSum(val, power+1)
		default:
			// ignore unsupported types
		}
	}
	return math.Pow(sum, float64(power))
}

func Ex2() {
	// Constructing a nested array like: [1, 2, 3, 4, [8, 4]]
	arr := []any{
		1, 2, 3, 4,
		[]any{8, 4},
	}

	result := powerSum(arr, 1)
	fmt.Println("Result:", result)
}
