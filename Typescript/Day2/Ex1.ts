/* 
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row.
Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01,
and each occurrence of 1 with 10.  For example, for n = 3, the 1st row is 0, the 2nd row is 01, 
and the 3rd row is 0110. Given two integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.
*/

function KthSymBol(n: number, k: number):number {
    if (n==1) {
        return 0
    }

    const length: number = 2 ** (n-1)
    const mid = length/2

    if (k <= mid) {
        return KthSymBol(n-1, k)
    } else {
        return KthSymBol(n-1, k-mid)
    }

}


console.log(KthSymBol(5, 4))