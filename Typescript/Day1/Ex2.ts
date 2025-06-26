function checkMono (num: number[]):boolean {
    let isIncr = true
    let isDecr = true

    for (let i = 0; i < num.length; i++) {
        if( num[i] > num[i+1] ) {
            isIncr = false
        }

        if( num[i] < num[i+1] ) {
            isDecr = false
        }        
    }

    return  isIncr || isDecr
}



const monoIncr = [1, 2, 3, 4, 5, 6]
const monoDecr = [9, 8, 7, 6, 5, 4]
const monoNone = [1, 3, 5, 2, 9, 4, 11, 6]

console.log(checkMono(monoIncr))
console.log(checkMono(monoDecr))
console.log(checkMono(monoNone))