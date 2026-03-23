# https://codeforces.com/contest/2169/problem/C
# 1300

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

def f(x, s):
    return pow(x, 2) - x - s

def solve(nums):
    s = 0
    mn = 0
    ans = 0
    for i, x in enumerate(nums):
        s += x
        ans = max(ans, f(i + 2, s) - mn)
        mn = min(mn, f(i + 2, s))
    return sum(nums) + ans

T = int(input())
for _ in range(T):
    n = int(input())
    nums = list(map(int, input().split()))
    print(solve(nums))