
function reverseStr(str) {
    return str.split('').reverse().join('')
}
// console.log(reverseStr('gazi nahian'));

function findMissingNumbers(arr) {
    let missing = []; // Step 1: Create empty array for missing numbers

    // Step 2: Check if the array is sorted
    let isSorted = arr.every((_, i) => i === 0 || arr[i] >= arr[i - 1]);

    // Step 3: Sort the array if it's unsorted
    if (!isSorted) {
        arr.sort((a, b) => a - b);
    }

    // Step 4: Find missing numbers
    for (let i = 0; i < arr.length - 1; i++) {
        let diff = arr[i + 1] - arr[i];
        if (diff > 1) {
            for (let j = 1; j < diff; j++) {
                missing.push(arr[i] + j);
            }
        }
    }

    // Step 5: Return a single number or an array
    return missing.length === 1 ? missing[0] : missing;
}


console.log(findMissingNumbers([1, 2, 3, 4, 5, , 7, 8, 9]));