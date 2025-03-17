var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];

//Slicing & Splicing Methods
var arr2 = arr.slice(1, 4) //returns a new array with elements from index 1 to 3 (excluding index 4)
var arr3 = arr.splice(1, 4) //removes elements from index 1 to 3 (excluding index 4) and returns a new array with the removed elements

console.log(arr2);
console.log(arr3);