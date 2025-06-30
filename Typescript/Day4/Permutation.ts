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

var permute = function(nums: number[]): number[][] {
    let res: number[][] = [];
    let n: number = nums.length;

    function helper(i: number): void {
        if (i === n - 1) {
            res.push([...nums]);
            return;
        }
        for (let j: number = i; j < n; j++) {
            [nums[i], nums[j]] = [nums[j], nums[i]];
            helper(i + 1);
            [nums[i], nums[j]] = [nums[j], nums[i]];
        }
    }

    helper(0);
    return res;
};



