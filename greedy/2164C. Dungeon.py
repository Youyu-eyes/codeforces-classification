# https://codeforces.com/contest/2164/problem/C
# 1400

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
from typing import List, Tuple
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(la, lb, lc):
    heapify(la)
    reserved = []
    lst0 = []
    lst = []
    for b, c in zip(lb, lc):
        if c == 0:
            lst0.append(b)
        else:
            lst.append((b, c))

    lst.sort(key = lambda p: (p[0], -p[1]))
    ans = 0
    for b, c in lst:
        while la:
            a = heappop(la)
            if a < b:
                reserved.append(a)
                continue
            a = max(a, c)
            heappush(la, a)
            ans += 1
            break
        if not la:
            break
    
    while la:
        reserved.append(heappop(la))

    lst0.sort()
    reserved.sort()

    # 注意，这里用不了二分，因为每次要删除lst[i]，而python中pop函数加参数的开销很大，可以达到O(n)

    i, j = 0, 0
    while i < len(reserved) and j < len(lst0):
        if reserved[i] >= lst0[j]:
            ans += 1
            j += 1
        i += 1
    
    return ans

T = ii()
for _ in range(T):
    n, m = many_ii()
    a = lii()
    b = lii()
    c = lii()
    print(solve(a, b, c))