/*

The tower of Hanoi is a famous puzzle where we have three rods and N disks.
The objective of the puzzle is to move the entire stack to another rod.
You are given the number of discs N. Initially, these discs are in the rod 1.
You need to print all the steps of discs movement so that all the discs reach the 3rd rod.
 Also, you need to find the total moves.

Note: The discs are arranged such that the top disc is numbered 1 and the bottom-most disc is numbered N.
Also, all the discs have different sizes and a bigger disc cannot be put on the top of a smaller disc.
Refer the provided link to get a better clarity about the puzzle.

*/

package main

import (
	"fmt"
)

// count sẽ được dùng để đếm số bước di chuyển
var count int

func toh(N int, from, to, aux string) {
	// Base case: Nếu chỉ có 1 đĩa, di chuyển trực tiếp từ 'from' sang 'to'
	if N == 1 {
		fmt.Printf("Move disc %d from %s to %s\n", N, from, to)
		count++ // Tăng số bước di chuyển
		return
	}

	// Bước 1: Di chuyển N-1 đĩa từ 'from' sang 'aux' (sử dụng 'to' làm cọc trung gian)
	toh(N-1, from, aux, to)

	// Bước 2: Di chuyển đĩa lớn nhất (thứ N) từ 'from' sang 'to'
	fmt.Printf("Move disc %d from %s to %s\n", N, from, to)
	count++ // Tăng số bước di chuyển

	// Bước 3: Di chuyển N-1 đĩa từ 'aux' sang 'to' (sử dụng 'from' làm cọc trung gian)
	toh(N-1, aux, to, from)
}

func TowerOfHanoi() {
	numDiscs := 5 // Số đĩa bạn muốn di chuyển
	toh(numDiscs, "A", "C", "B")
	fmt.Printf("\nTotal moves for %d discs: %d\n", numDiscs, count)
}
