/*Problem 1: Find the Maximum Number in an Array*/
// function findMax(arr) {
//     let max = arr[0]
//     for (let i = 0; i < arr.length; i++) {
//         if (arr[i] > max) {
//             max = arr[i]
//         }

//     }
//     return max
// }
// console.log(findMax([1, 2, 30, 4, 5]));

/*Problem 2: Reverse an Array*/
function reverseArray1(arr) {
    let start = 0;
    let end = arr.length - 1;

    while (start < end) {
        let temp = arr[start];
        arr[start] = arr[end];
        arr[end] = temp;
        start++;
        end--;
    }
    return arr;
}


function reverseArray2(arr) {
    arr = arr.reverse();
    return arr;
}

console.log(reverseArray1([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]));
console.log(reverseArray2([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]));