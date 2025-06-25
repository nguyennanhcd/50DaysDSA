'''
You are given an array of integers in which each subsequent value is not less than the previous
value. Write a function that takes this array as an input and returns a new array with the squares
of each number sorted in ascending order.
'''
from typing import List

def sorted_squares(nums: List[int]) -> List[int]:
    n = len(nums)
    result = [0] * n
    left, right = 0, n - 1
    pos = n - 1

    while left <= right:
        left_sq = nums[left] ** 2
        right_sq = nums[right] ** 2

        if left_sq > right_sq:
            result[pos] = left_sq
            left += 1
        else:
            result[pos] = right_sq
            right -= 1

        pos -= 1

    return result

nums = [-5, 3, 4, 8]
print(sorted_squares(nums))
