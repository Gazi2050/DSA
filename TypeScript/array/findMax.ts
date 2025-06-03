// solution-01 
// function findMax(arr: number[]): number {
//     return Math.max(...arr);
// }

// solution-02
// function findMax(arr: number[]): number {
//     let max = -Infinity;
//     for (let i = 0; i < arr.length; i++) {
//         if (arr[i] > max) {
//             max = arr[i]
//         }
//     }
//     return max;
// }

// solution-03
function findMax(arr: number[]): number {
    let max = -Infinity;
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] > max) {
            max = arr[i]
        }
    }
    return max;
}

const arr = [1, 2, 3, 4, 5, 6];
const result = findMax(arr);
console.log(result);
