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


function permuteUnique(nums: number[]): number[][] {
    const res: number[][] = [];

    function helper(i: number): void {
        if (i === nums.length - 1) {
            res.push([...nums]); // copy mảng để tránh dùng chung tham chiếu
            return;
        }

        const used: Set<number> = new Set();
        for (let j = i; j < nums.length; j++) {
            if (!used.has(nums[j])) {
                used.add(nums[j]);
                [nums[i], nums[j]] = [nums[j], nums[i]];
                helper(i + 1);
                [nums[i], nums[j]] = [nums[j], nums[i]]; // backtrack
            }
        }
    }

    helper(0);
    return res;
}

// Test
console.log(permuteUnique([3, 3, 2]));
