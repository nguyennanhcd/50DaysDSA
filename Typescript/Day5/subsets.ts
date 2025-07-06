/*

Given an integer array unique elements, return all possible subsets ( the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

*/

function subsets(nums: number[]): number[][] {
    const output: number[][] = [];
    
    function helper(nums: number[], i: number, subset: number[]): void {
        if (i === nums.length) {
            output.push([...subset]);
            return;
        }
        
        // Exclude current element
        helper(nums, i + 1, subset);
        
        // Include current element
        subset.push(nums[i]);
        helper(nums, i + 1, subset);
        
        // Backtrack
        subset.pop();
    }
    
    helper(nums, 0, []);
    return output;
}

// Example usage
const nums: number[] = [1, 8, 7];
const powerSet: number[][] = subsets(nums);
console.log(powerSet);