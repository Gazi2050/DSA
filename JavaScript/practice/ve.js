//very easy 
function sum(a, b) {
    return a + b;
}

console.log(sum(2, 3));

//function for check even or odd
function isEven(num) {
    if (num % 2 === 0) {
        return 'even'
    }
    if (num % 2 !== 0) {
        return 'odd'
    }
}
console.log(isEven(9));

function isBig(a, b) {
    if (a > b) {
        return `${a} is bigger then ${b}`
    }
    else if (a === b) {
        return 'equal'
    }
    else if (a < b) {
        return `${b} is bigger then ${a}`
    }
}
console.log(isBig(5, 10))
