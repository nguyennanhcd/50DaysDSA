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

var count = 0;

function toh(N: number, from: number, aux: number, to: number) {
    if (N == 1) {
        console.log(`Move disc ${N} from ${from} to ${to}`);
        count++;
        return;
    }

    toh(N - 1, from, to, aux);  // Note: aux và to đảo chỗ nhau
    console.log(`Move disc ${N} from ${from} to ${to}`);
    count++;
    toh(N - 1, aux, from, to);
}

toh(5, 1, 2, 3);
console.log(`Total moves: ${count}`);