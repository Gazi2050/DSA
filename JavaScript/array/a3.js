//Searching Methods
var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];

const index = arr.indexOf(2);
const index2 = arr.indexOf(22);
console.log(index);
console.log(index2);

const isExist = arr.includes(7)
const isExist2 = arr.includes(77)
console.log(isExist);
console.log(isExist2);

const find1 = arr.find(e => e > 5)
const find2 = arr.findIndex(e => e > 5)
console.log(find1);
console.log(find2);