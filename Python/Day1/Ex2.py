'''
An array is monotonic if it is either monotone increasing or monotone decreasing.
An array is monotone increasing if all its elements form left to right are non-decreasing.
Given an integer array return true if the given array is monotonic, or false otherwise
'''

def checkMono(arr):
    increasing = True
    decreasing = True

    for i in range(len(arr) - 1):
        if arr[i] < arr[i + 1]:
            decreasing = False
        if arr[i] > arr[i + 1]:
            increasing = False

    return increasing or decreasing



monoIncr = [1,2,3,4,5,6]
monoDecr = [9,8,7,6,5,4]
monoNone = [1,3,5,2,9,4,11,6]

print(checkMono(monoIncr))
print(checkMono(monoDecr))
print(checkMono(monoNone)) 