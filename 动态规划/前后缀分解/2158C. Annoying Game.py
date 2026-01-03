# https://codeforces.com/problemset/problem/2158/C
# 1400
# 博弈论，博弈DP，前后缀分解

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
from copy import deepcopy
from typing import List, Tuple, Callable
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(n, k, a, b):
    suf = [0] * n
    suf[-1] = a[-1]
    for i in range(n - 2, -1, -1):
        suf[i] = max(suf[i + 1] , 0) + a[i]

    if k & 1 == 0:
        return max(suf)
    
    pre = 0
    ans = -inf
    for i, x in enumerate(a):
        pre = max(pre, 0) + x
        ans = max(ans, pre + suf[i] - x + b[i])
    
    return ans


T = ii()
for _ in range(T):
    n, k = many_ii()
    a = lii()
    b = lii()
    print(solve(n, k, a, b))