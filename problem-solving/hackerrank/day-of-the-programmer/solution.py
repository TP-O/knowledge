#!/bin/python3

import math
import os
import random
import re
import sys

def leap_year(year):
    if year < 1919:
        if year%4 == 0:
            return True
    elif year%400 == 0 or (year%4 == 0 and year%100 != 0):
            return True

    return False

# Complete the dayOfProgrammer function below.
def dayOfProgrammer(year):
    if year == 1918:
        day = 26
    elif leap_year(year):
        day = 12
    else:
        day = 13

    date = str(day)+'.09.'+str(year)

    return date

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    year = int(input().strip())

    result = dayOfProgrammer(year)

    fptr.write(result + '\n')

    fptr.close()
