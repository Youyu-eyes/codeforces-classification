# https://codeforces.com/contest/2152/problem/B
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
from itertools import zip_longest
from itertools import pairwise
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

def judge(n, D, K):
    if D < K:
        return n - D
    elif K < D:
        return D
    else:
        return 0

def solve(n, rK, cK, rD, cD):
    return max(judge(n, rD, rK), judge(n, cD, cK))

T = ii()
for _ in range(T):
    n, rK, cK, rD, cD = many_ii()
    print(solve(n, rK, cK, rD, cD))