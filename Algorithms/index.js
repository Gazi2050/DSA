function summmation(n) {
    let sum = 0;
    for (let i = 1; i <= n; i++) {
        sum += i;
    }
    return sum;
}

console.log(summmation(5));

function sum(n) {
    return (n * (n + 1)) / 2
}
console.log(sum(5));