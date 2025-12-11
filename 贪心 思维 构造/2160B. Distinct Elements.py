# https://codeforces.com/contest/2160/problem/B

import sys
from math import inf
from math import gcd
from math import lcm
from math import sin
from math import pi
import heapq as h
from functools import cache
from itertools import accumulate
from itertools import zip_longest
from bisect import bisect_left
from bisect import bisect_right
from collections import defaultdict
from collections import Counter
from collections import deque as q
import datetime
import re
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(n, b):
    a = [0] * n
    a[0] = 1
    max_n = 2
    for i in range(1, n):
        j = i + b[i - 1] - b[i]
        if j < 0:
            a[i] = max_n
            max_n += 1
        else:
            a[i] = a[j]
    return a
        
T = ii()
for _ in range(T):
    n = ii()
    b = li()
    a = solve(n, b)
    for x in a:
        print(x, end = ' ')
    print()