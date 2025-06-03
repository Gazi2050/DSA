const arr: number[] = [3, 6, 2, 8, 1, 6];

function isSorted(array: number[]) {
    const sortedArray = [...array].sort((a, b) => a - b); // copy and sort
    const isEqual = array.every((value, index) => value === sortedArray[index]);
    console.log(isEqual);
}
isSorted(arr)