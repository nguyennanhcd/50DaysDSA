'''

Given an integer array unique elements, return all possible subsets ( the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

'''


def subsets(nums):
    output = []
    def helper(nums, i, subset):
        if i == len(nums):
            output.append(subset.copy())
            return
        helper(nums, i+1, subset)
        subset.append(nums[i])
        helper(nums, i+1, subset)
        subset.pop()
    helper(nums, 0, [])
    return output

power_set = subsets([1,8,7])
print(power_set)