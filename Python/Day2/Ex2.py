'''
There are n friends that are playing a game. The friends are sitting in a circle and are numbered from 1 to n in clockwise order.
More formally, moving clockwise from the ith friend brings you to the (i+1)th friend for 1 <= i < n,
and moving clockwise from the nth friend brings you to the 1st friend.
The rules of the game are as follows:
Start at the 1st friend. Count the next k friends in the clockwise direction including the friend you started at.
The counting wraps around the circle and may count some friends more than once.
The last friend you counted leaves the circle and loses the game.
If there is still more than one friend in the circle,
go back to step 2 starting from the friend immediately clockwise of the friend who just lost and repeat.
Else, the last friend in the circle wins the game. Given the number of friends, n, and an integer k, return the winner of the game.

'''


def findTheWinner(n, k):
    arr = [i+1 for i in range (n)]

    def helper(arr, startIndex):
        #base case
        if len(arr) == 1:
            return arr[0]
        
        # recursive case 
        index_to_remove = (startIndex + k -1) %len(arr)
        del arr[index_to_remove]
        return helper(arr, index_to_remove)

    return helper(arr, 0)


def findTheWinner2(n, k):
    def josephus(n):
        if n== 1:
            return 0
        return (josephus(n-1) +k) %n

    return josephus(n) +1


def findTheWinner3(n, k):
    survivor = 0

    for i in range (2, n+1):
        survivor = (survivor + k ) %i
    return survivor +1


n = 4
k = 4
result = findTheWinner3(n, k)
print(result)
    
