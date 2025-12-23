# https://codeforces.com/problemset/problem/1895/C

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

def solve(tickets):
    cnt = defaultdict(int)
    ans = 0

    for s in tickets:
        n = len(s)
        sum_ = sum(map(int, s))
        cnt[(n, sum_)] += 1

    for s in tickets:
        n = len(s)
        sum_ = sum(map(int, s))

        pre = [0] * (n + 1)
        for i in range(n):
            pre[i + 1] = pre[i] + int(s[i])

        start_h = (n + 1) // 2
        for h in range(start_h, n + 1):
            ans += cnt[(2 * h - n, 2 * pre[h] - sum_)]

        max_k = (n - 1) // 2
        for k in range(1, max_k + 1):
            ans += cnt[(n - 2 * k, sum_ - 2 * pre[k])]

    return ans

n = ii()
s = lsi()
print(solve(s))