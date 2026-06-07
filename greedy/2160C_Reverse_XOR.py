# https://codeforces.com/contest/2160/problem/C
# 1300

import sys
from math import inf
from math import gcd
from math import lcm
from math import sin
from math import pi
from heapq import heapify
from heapq import heappush
from heapq import heappop
from heapq import heapreplace
from heapq import nsmallest
from heapq import nlargest
from functools import cache
from itertools import accumulate
from itertools import pairwise
from itertools import zip_longest
from bisect import bisect_left
from bisect import bisect_right
from collections import defaultdict
from collections import Counter
from collections import deque
import datetime
import re
from typing import List, Tuple
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(n):
    s = bin(n)[2:]
    for i in range(len(s) - 1, -1, -1):
        if s[i] == '1':
            s = s[:i + 1]
            break
    if s != s[::-1]:
        return False
    if len(s) & 1 and s[len(s) // 2] == '1':
        return False
    return True

T = ii()
for _ in range(T):
    n = ii()
    if solve(n):
        print('YES')
    else:
        print('NO')