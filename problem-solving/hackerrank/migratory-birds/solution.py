#!/bin/python3

import math
import os
import random
import re
import sys

def find_max_index(arr):
    max_val = arr[0]
    index = 0

    for i in range(len(arr)):
        if arr[i] > max_val:
            max_val = arr[i]
            index = i
        elif arr[i] == max_val:
            if i < index:
                index = i

    return index+1

# Complete the migratoryBirds function below.
def migratoryBirds(arr):
    count = [0, 0, 0, 0, 0]
    index = 0

    for i in arr:
        count[i-1] += 1

    index = find_max_index(count)

    return index

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr_count = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    result = migratoryBirds(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
