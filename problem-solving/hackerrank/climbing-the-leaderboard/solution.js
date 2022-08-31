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

function rejectDuplication (arr) {
    let new_arr = []
    let length = arr.length
    let new_length = new_arr.length

    for (let i = 0; i < length; i++) {
        if (arr[i] !== new_arr[new_length - 1] || !new_length) {
            new_arr[new_length] = arr[i]
            new_length += 1
        }
    }

    return new_arr
}

// Complete the climbingLeaderboard function below.
function climbingLeaderboard(scores, alice) {
    let new_scores = rejectDuplication(scores)
    let ranks = []
    let ranks_length = 0

    for (let i = 0; i < alice.length; i++) {
        let start = 0
        let end = new_scores.length - 1

        while (true) {
            let index = Math.floor((start + end) / 2)

            if (start === end - 1) {
                if (new_scores[start] <= alice[i]) {
                    ranks[ranks_length] = start + 1
                }
                else if (new_scores[end] <= alice[i]) {
                    ranks[ranks_length] = end + 1
                }
                else {
                    ranks[ranks_length] = end + 2
                }
                ranks_length += 1
                break
            }
            else if (new_scores[index] > alice[i]) {
                start = index
            }
            else if (new_scores[index] < alice[i]) {
                end = index
            }
            else {
                ranks[ranks_length] = index + 1
                ranks_length += 1
                break
            }
        }
    }

    return ranks
}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const scoresCount = parseInt(readLine(), 10);

    const scores = readLine().split(' ').map(scoresTemp => parseInt(scoresTemp, 10));

    const aliceCount = parseInt(readLine(), 10);

    const alice = readLine().split(' ').map(aliceTemp => parseInt(aliceTemp, 10));

    let result = climbingLeaderboard(scores, alice);

    ws.write(result.join("\n") + "\n");

    ws.end();
}
