/*

Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

*/

function combine(n: number, k: number): number[][] {
    const result: number[][] = [];
    const combination: number[] = [];

    function backtrack(start: number) {
        if (combination.length === k) {
            result.push([...combination]);
            return;
        }
        for (let i = start; i <= n; i++) {
            combination.push(i);
            backtrack(i + 1);
            combination.pop();
        }
    }

    backtrack(1);
    return result;
}

console.log(combine(4, 2)); // Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]