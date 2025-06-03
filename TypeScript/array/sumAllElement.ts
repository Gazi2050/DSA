const arr: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9];

//solution-01
//Using for loop
function sumByLoop(array: number[]) {
    let sum: number = 0
    for (let i = 0; i < array.length; i++) {
        sum += arr[i]
    }
    console.log(sum);
}
// sumByLoop(arr)

//solution-02
//Using forEach() Method
function sumByForEach(array: number[]) {
    let sum = 0;
    array.forEach(num => sum += num);
    console.log(sum);
}

// sumByForEach(arr)

//solution-03
//Using Recursion
function sumByRecursion(array: number[], idx: number): number {
    if (idx === array.length) {
        return 0;
    }

    const sum = array[idx] + sumByRecursion(array, idx + 1);

    if (idx === 0) {
        console.log(sum);
    }

    return sum;
}

sumByRecursion(arr, 0)