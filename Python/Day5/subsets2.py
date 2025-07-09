'''

Given an integer array nums that may contain duplicates, return all possible subsets
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


'''


def subsets_with_dup(nums: list[int]) -> list[list[int]]:
    res: list[list[int]] = []
    nums.sort()
    
    def subsets(index: int, curr: list[int]) -> None:
        if index == len(nums):
            res.append(curr[:])
            return
        
        # Include the current number
        curr.append(nums[index])
        subsets(index + 1, curr)
        curr.pop()
        
        # Exclude the current number and skip all duplicates
        while index < len(nums) - 1 and nums[index] == nums[index + 1]:
            index += 1
        subsets(index + 1, curr)
    
    subsets(0, [])
    return res


print(subsets_with_dup([1, 2, 2])) 