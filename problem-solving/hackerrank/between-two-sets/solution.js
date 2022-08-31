'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

function checkStCondition (num, arr) {
    for (let i = 0; i < arr.length; i++)
        if (num % arr[i]) return false
    return true
}

function checkNdCondition (num, arr) {
    for (let i = 0; i < arr.length; i++)
        if (arr[i] % num) return false
    return true
}

function findMin (arr) {
    return Math.min.apply(Math, arr)
}

function findMax (arr) {
    return Math.max.apply(Math, arr)
}

function getTotalX(a, b) {
    let start = findMax(a)
    let end = findMin(b)
    let count = 0

    for (let i = start; i <= end; i++)
        if (checkStCondition(i, a) && checkNdCondition(i, b)) count +=1

    return count

}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const firstMultipleInput = readLine().replace(/\s+$/g, '').split(' ');

    const n = parseInt(firstMultipleInput[0], 10);

    const m = parseInt(firstMultipleInput[1], 10);

    const arr = readLine().replace(/\s+$/g, '').split(' ').map(arrTemp => parseInt(arrTemp, 10));

    const brr = readLine().replace(/\s+$/g, '').split(' ').map(brrTemp => parseInt(brrTemp, 10));

    const total = getTotalX(arr, brr);

    ws.write(total + '\n');

    ws.end();
}
