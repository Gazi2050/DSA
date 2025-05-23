function replaceDigit(str) {
    return str.replace(/\d/g, digit => digit < '5' ? '0' : '1')
}