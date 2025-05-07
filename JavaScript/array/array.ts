/**Array
 * Arrays are linear data structures used to store elements in a contiguous memory location. They're indexed starting at 0 and support operations like traversal, insertion, and deletion. Arrays are the foundation for more advanced techniques like two-pointer and sliding window. JavaScript arrays are dynamic and come with built-in methods. Understanding core operations without built-ins is essential for mastering algorithms.
 */

// function reverseArray(arr: number[]) {
//     const result = arr.reverse();
//     console.log(result);
// }

function reverseArray(arr: number[]): number[] {
    let left = 0, right = arr.length - 1;
    while (left < right) {
        [arr[left], arr[right]] = [arr[right], arr[left]]
        left++
        right--
    }
    console.log(arr);
    console.log(arr[0]);
    return arr
}

// reverseArray([1, 2, 3, 4, 5])