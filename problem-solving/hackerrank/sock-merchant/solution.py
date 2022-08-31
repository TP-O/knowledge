#!/bin/python3

import math
import os
import random
import re
import sys

def find_index(val, ar):
    for i in range(len(ar)):
        if ar[i] == val:
            return i

    return None

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    socks = []
    count = []
    pair = 0

    for i in range(len(ar)):
        index = find_index(ar[i], socks)

        if index is not None:
            count[index] += 0.5
        else:
            socks.append(ar[i])
            count.append(0.5)

    for i in count:
        pair += int(i)

    return pair

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)

    fptr.write(str(result) + '\n')

    fptr.close()
