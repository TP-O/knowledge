#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the breakingRecords function below.
def breakingRecords(scores):
    min = scores[0]
    max = scores[0]
    min_count = 0
    max_count = 0
    for i in scores:
        if i < min:
            min = i
            min_count += 1
        elif i > max:
            max = i
            max_count += 1

    return str(max_count)+' '+str(min_count)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = breakingRecords(scores)

    fptr.write(result)
    fptr.write('\n')

    fptr.close()
