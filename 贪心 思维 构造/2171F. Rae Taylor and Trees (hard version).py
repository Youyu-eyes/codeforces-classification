# https://codeforces.com/contest/2171/problem/F

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
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def solve(n, nums):
    pre = [0] * n
    pre[0] = nums[0]
    for i in range(1, n):
        pre[i] = min(pre[i - 1], nums[i])
    
    suf = [(0, 0)] * n
    suf[-1] = (nums[-1], n - 1)
    for i in range(n - 2, -1, -1):
        if nums[i] > suf[i + 1][0]:
            suf[i] = (nums[i], i)
        else:
            suf[i] = suf[i + 1]

    valid = True
    for i in range(n - 1):
        if pre[i] > suf[i + 1][0]:
            valid = False
            break
    
    if not valid:
        return []

    edges = []
    l = 0
    while l < n:
        r = suf[l][1]
        for i in range(l, r):
            edges.append((nums[i], nums[r]))
        if l > 1:
            edges.append((pre[l - 1], nums[r]))
        l = r + 1
    
    return edges

T = ii()
for _ in range(T):
    n = ii()
    nums = lii()
    ans = solve(n, nums)
    if ans:
        print('Yes')
        for u, v in ans:
            print(u, v)
    else:
        print('No')