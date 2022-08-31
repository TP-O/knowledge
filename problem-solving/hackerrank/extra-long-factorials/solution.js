'use strict';

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

// Complete the extraLongFactorials function below.
function extraLongFactorials(n) {
    let end = n
    let r = []
    while(n > 0) {
        if(n >= 10) {
            r.unshift(n % 10)
        }
        else {
            r.unshift(n)
        }
        n = Math.floor(n / 10)
    }

    for(let d = end - 1; d > 1; d--) {
        let st = r
        let w = d
        let nd = []
        let e = []

        while(w > 0) {
            if(w >= 10) {
                nd.unshift(w % 10)
            }
            else {
                nd.unshift(w)
            }
            w = Math.floor(w / 10)
        }

        for(let i = nd.length - 1; i >= 0; i--) {
            let a = []
            let z, carry = 0

            for(let j = st.length - 1; j >= 0; j--) {
                z = (nd[i] * st[j]) + carry

                if(z >= 10) {
                    carry = Math.floor(z / 10)
                    a.unshift(z % 10)
                }
                else {
                    carry = 0
                    a.unshift(z)
                }
            }
            if(carry) {
                if(carry >= 10) {
                    a.unshift(carry % 10)
                    a.unshift(Math.floor(carry / 10))
                }
                else {
                    a.unshift(carry)
                }
            }
            e.push(a)
        }

        r = e[0]
        e.forEach((x, k) => {
            if(k) {
                let t = r
                let z, carry = 0
                let l1 = t.length - 1
                let l2 = x.length - 1
                r = []

                for(let i = 0; i < k; i++) {
                    r.unshift(t[l1 - i])
                }

                for(let i = l1 - k; i >= 0; i--) {
                    z = t[i] + x[l2] + carry

                    if(z >= 10) {
                        carry = 1;
                        r.unshift(z % 10)
                    }
                    else {
                        carry = 0;
                        r.unshift(z);
                    }

                    l2--
                }

                for(let j = l2; j >= 0; j--) {
                    z = x[j] + carry

                    if(z >= 10) {
                        carry = 1
                        r.unshift(z % 10)
                    }
                    else {
                        carry = 0
                        r.unshift(z)
                    }
                }

                if(carry) {
                    if(carry >= 0) {
                        r.unshift(carry % 10)
                        r.unshift(Math.floor(carry / 10))
                    }
                    else {
                        r.unshift(carry)
                    }
                }
            }
        })
    }

    r = r.join('')

    console.log(r)
}

function main() {
    const n = parseInt(readLine(), 10);

    extraLongFactorials(n);
}
