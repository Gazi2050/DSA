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