'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.replace(/\s*$/, '')
        .split('\n')
        .map(str => str.replace(/\s*$/, ''));

    main();
});

function readLine() {
    return inputString[currentLine++];
}

// Complete the appendAndDelete function below.
function appendAndDelete(s, t, k) {
    let l1 = s.length
    let l2 = t.length
    let index = l2 - 1
    let min = -1

    for(let i = 0; i < l2; i++) {
        if(s[i] != t[i]) {
            index = i
        }
    }

    if(k < (l1 + l2)) {
        min = (l1 - index) + (l2 - index)

        if((k - min) % 2 === 0 && k >= min) {
            return 'Yes'
        }
        else {
            return 'No'
        }
    }
    else {
        return 'Yes'
    }
}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const s = readLine();

    const t = readLine();

    const k = parseInt(readLine(), 10);

    let result = appendAndDelete(s, t, k);

    ws.write(result + "\n");

    ws.end();
}
