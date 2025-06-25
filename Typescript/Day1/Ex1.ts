/* You are given an array of integers in which each subsequent value is not less than the previous
value. Write a function that takes this array as an input and returns a new array with the squares
of each number sorted in ascending order.
*/

const squares = (num: number[]): number[]=> {
	const n = num.length
    let left = 0
    let right = n-1
    let pos = n-1
    let result: number[] = new Array(n)

    while (left <= right) {
        let leftValue = num[left] **2
        let rightValue = num[right] **2

        if (leftValue > rightValue) {
            result[pos] = leftValue
            left++
        } else {
            result[pos] = rightValue
            right --
        }
        pos --
    }

    return result
}

const num = [-5, 3, 4, 8]
const result = squares(num)
console.log(result)
