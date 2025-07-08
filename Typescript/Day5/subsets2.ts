/*

Given an integer array nums that may contain duplicates, return all possiblesubsets

(the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example : 

nums = [1,3,3]

Output =

[

[],

[1],

[3],

[1,3],

[3,3],

[1,3,3]


*/

function subsetsWithDup(nums: number[]): number[][] {
    let res: number[][] = [];
    nums.sort((a: number, b: number) => a - b);
 
    function subsets(index: number, curr: number[]): void {
        if (index === nums.length) {
            res.push([...curr]);
            return;
        }
        // Include the current number
        curr.push(nums[index]);
        subsets(index + 1, curr);
        curr.pop();
        // Exclude the current number and skip all duplicates
        while (index < nums.length - 1 && nums[index] === nums[index + 1]) {
            index++;
        }
        subsets(index + 1, curr);
    }
 
    subsets(0, []);
    return res;
}

console.log(subsetsWithDup([1, 3, 3,4]));