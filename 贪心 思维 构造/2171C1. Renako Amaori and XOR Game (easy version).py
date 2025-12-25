# https://codeforces.com/contest/2171/problem/C1
# 1100

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
input = lambda: sys.stdin.readline().strip()
MOD = 1_000_000_007

def solve(Ajisai, Mai):
    xor_a = 0
    for x in Ajisai:
        xor_a ^= x
    xor_m = 0
    for x in Mai:
        xor_m ^= x
    for i, (x, y) in enumerate(zip(Ajisai, Mai)):
        if x == y:
            continue
        if (1 - i % 2 and xor_a & 1 == 0) or (1 - i % 2 == 0 and xor_m & 1 == 0):
            xor_a ^= 1
            xor_m ^= 1
    if xor_a > xor_m:
        return 'Ajisai'
    if xor_m > xor_a:
        return 'Mai'
    return 'Tie'

T = int(input())
for _ in range(T):
    n = int(input())
    a = list(map(int, input().split()))
    m = list(map(int, input().split()))
    print(solve(a, m))