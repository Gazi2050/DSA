// solution-01
// function removeDuplicates(arr: number[]): number[] {
//     const result: number[] = [];
//     const seen = new Set();
//     for (let i = 0; i < arr.length; i++) {
//         if (!seen.has(arr[i])) {
//             seen.add(arr[i]);
//             result.push(arr[i]);
//         }
//     }
//     return result;
// }

// solution-02
// function removeDuplicates(arr: number[]): number[] {
//     return Array.from(new Set(arr));
// }

// solution-03
function removeDuplicates(arr: number[]): number[] {
    return [...new Set(arr)];
}

const arr = [1, 2, 3, 4, 5, 6, 6]
const result = removeDuplicates(arr);
console.log(result);