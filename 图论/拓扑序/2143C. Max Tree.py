# https://codeforces.com/problemset/problem/2143/C
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
many_si = lambda: input().split()
lii = lambda: list(many_ii())
lsi = lambda: list(many_si())
MOD = 1_000_000_007

def topologicalSort(n: int, edges: List[List[int]]) -> List[int]:
    g = [[] for _ in range(n)]
    in_deg = [0] * n
    for x, y in edges:
        g[x].append(y)
        in_deg[y] += 1  # 统计 y 的先修课数量

    topo_order = []
    q = deque(i for i, d in enumerate(in_deg) if d == 0)  # 没有先修课，可以直接上
    while q:
        x = q.popleft()
        topo_order.append(x)
        for y in g[x]:
            in_deg[y] -= 1  # 修完 x 后，y 的先修课数量减一
            if in_deg[y] == 0:  # y 的先修课全部上完
                q.append(y)  # 加入学习队列

    if len(topo_order) < n:  # 图中有环
        return []
    return topo_order

def solve(n, edges):
    topo_order = topologicalSort(n, edges)
    ans = [0] * n
    for i, x in enumerate(topo_order):
        ans[x] = n - i
    return ans

T = ii()
for _ in range(T):
    n = ii()
    edges = []
    for _ in range(n - 1):
        u, v, x, y = many_ii()
        if x >= y:
            edges.append([u - 1, v - 1])
        else:
            edges.append([v - 1, u - 1])
    ans = solve(n, edges)
    print(' '.join(map(str, ans)))