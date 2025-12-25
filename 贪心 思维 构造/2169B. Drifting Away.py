# https://codeforces.com/contest/2169/problem/B
# 1100 正常写法
# 1250 正则式匹配

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
li = lambda: list(many_ii)
MOD = 1_000_000_007

def solve(s):
    if s == '*':
        return 1
    match = re.match(r'^(<*)(\*?)(>*)$', s)
    if match:
        a = len(match.group(1))
        b = len(match.group(2))
        c = len(match.group(3))
        if a > 0 or c > 0:
            return max(a, c) + b
    return -1

T = ii()
for _ in range(T):
    s = si()
    print(solve(s))