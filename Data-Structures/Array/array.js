// Creating arrays
const createArray = Array.from('12345'); // ['1', '2', '3', '4', '5']
const arr = [1, 2, 3, 4, 5];
const isArray = Array.isArray(arr); // true
const arrLength = arr.length; // 5

// Adding and removing elements
const addFirstElement = arr.unshift(0); // [0, 1, 2, 3, 4, 5]
const removeFirstElement = arr.shift(); // [1, 2, 3, 4, 5]
const removeLastElement = arr.pop(); // [1, 2, 3, 4]
const addLastElement = arr.push(6); // [1, 2, 3, 4, 6]

// Accessing elements
const findIndex = arr.at(2); // 3 (element at index 2)
const findElement = arr.indexOf(3); // 2 (index of element 3)
const isElementExist = arr.includes(5); // true

// Combining arrays
const anotherArray = [6, 7, 8];
const combinedArray = arr.concat(anotherArray); // [1, 2, 3, 4, 6, 6, 7, 8]

// Copying within an array
const copyWithinArray = arr.copyWithin(0, 2, 4); // [3, 4, 3, 4, 6]

// Filling an array
const filledArray = arr.fill(0, 1, 3); // [3, 0, 0, 4, 6]

// Reversing an array
const reversedArray = arr.reverse(); // [6, 4, 0, 0, 3]

// Slicing an array
const slicedArray = arr.slice(1, 3); // [4, 0]

// Sorting an array
const sortedArray = arr.sort(); // [0, 0, 3, 4, 6]

// Iterating over arrays
arr.forEach(element => console.log(element)); // Logs each element

// Mapping over an array
const mappedArray = arr.map(element => element * 2); // [0, 0, 6, 8, 12]

// Filtering an array
const filteredArray = arr.filter(element => element > 2); // [3, 4, 6]

// Reducing an array
const reducedValue = arr.reduce((acc, curr) => acc + curr, 0); // 13
const reducedRightValue = arr.reduceRight((acc, curr) => acc + curr, 0); // 13

// Finding an element
const foundElement = arr.find(element => element > 2); // 3
const foundIndex = arr.findIndex(element => element > 2); // 2

// Checking conditions
const allAboveZero = arr.every(element => element > 0); // false
const someAboveFour = arr.some(element => element > 4); // true

// Flattening arrays
const nestedArray = [1, [2, [3, [4, 5]]]];
const flatArray = nestedArray.flat(2); // [1, 2, 3, [4, 5]]
const flatMappedArray = arr.flatMap(element => [element, element * 2]); // [0, 0, 0, 0, 3, 6, 4, 8, 6, 12]

// Array iterators
const arrayKeys = arr.keys(); // Array Iterator
const arrayValues = arr.values(); // Array Iterator
const arrayEntries = arr.entries(); // Array Iterator

// Other methods
const arrayToString = arr.toString(); // '0,0,3,4,6'
const arrayJoin = arr.join('-'); // '0-0-3-4-6'
const arrayLastIndexOf = arr.lastIndexOf(3); // 2
const arrayFindLast = arr.findLast(element => element > 2); // 6
const arrayFindLastIndex = arr.findLastIndex(element => element > 2); // 4

// Adding methods using prototype
Array.prototype.customMethod = function () {
    return 'This is a custom method added to all arrays!';
};

const customResult = arr.customMethod(); // 'This is a custom method added to all arrays!'
