/*
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
*/

package main

import (
	"fmt"
	"slices"
)

func RemoveIndex(s []int, index int) []int {
	return slices.Delete(s, index, index+1)
}

func findTheWinner(players []int, startingPosition int, k int) int {
	if len(players) == 1 {
		return players[0]
	}

	indexToRemove := (startingPosition + k - 1) % len(players)
	newPlayers := RemoveIndex(players, indexToRemove)
	newStartingPosition := indexToRemove % len(newPlayers)

	return findTheWinner(newPlayers, newStartingPosition, k)
}

func findTheWinner2(n int, k int) int {
	var josephus func(int) int
	josephus = func(n int) int {
		if n == 1 {
			return 0
		}
		return (josephus(n-1) + k) % n
	}
	return josephus(n) + 1
}

func findTheWinner3(n int, k int) int {
	survivor := 0

	for i := 2; i <= n; i++ {
		survivor = (survivor + k) % i
	}

	return survivor + 1
}

func Ex2() {
	n := 5
	k := 2

	players := make([]int, n)
	for i := range n {
		players[i] = i + 1
	}

	result := findTheWinner(players, 0, k)
	result2 := findTheWinner2(n, k)
	result3 := findTheWinner3(n, k)

	fmt.Println("Winner is:", result)
	fmt.Println("Winner2 is:", result2)
	fmt.Println("Winner3 is:", result3)
}
