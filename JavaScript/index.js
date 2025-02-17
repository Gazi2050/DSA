function outer() {
    let count = 0;
    return function inner() {
        count++;
        console.log(count);
    }
}

function stringToObjectArray(str) {
    return str.split("").filter(char => /[a-zA-Z0-9]/.test(char)).map((char, index) => ({
        id: index + 1,
        char: char
    }));
}

function reverseString(str) {
    return str.split('').reverse().join('')
}

function isPalindrome(str) {
    let normalStr = str.toLowerCase().replace(/[^a-z0-9]/g, "");
    let reverseStr = normalStr.split('').reverse().join('');

    return normalStr === reverseStr;
}

function isAnagram(str1, str2) {
    let string1 = str1.toLowerCase().replace(/[^a-z0-9]/g, "").split('').sort().join('');
    let string2 = str2.toLowerCase().replace(/[^a-z0-9]/g, "").split('').sort().join('');

    return string1 === string2;
}

function countVowels(str) {
    return (str.match(/[aeiou]/gi) || []).length;
}

function removeDuplicates(str) {
    return [...new Set(str)].join('')
}

function charCount(str) {
    let cleanStr = str.toLowerCase().replace(/[^a-z0-9]/g, "");
    let charCount = [];
    for (let char of cleanStr) {
        let found = charCount.find(obj => obj.char === char);
        if (found) {
            found.count++;
        } else {
            charCount.push({ id: charCount.length + 1, char: char, count: 1 });
        }
    }

    return charCount;
}
console.log(charCount('hello world'));