# https://codeforces.com/contest/2172/problem/M

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

def solve(n, g, types, k):
    dist = [-1] * (n + 1)
    dist[1] = 0
    q = deque([1])
    
    while q:
        u = q.popleft()
        for v in g[u]:
            if dist[v] == -1:
                dist[v] = dist[u] + 1
                q.append(v)

    max_dist = [-1] * (k + 1)
    for i in range(1, n + 1):
        t = types[i - 1]
        max_dist[t] = max(max_dist[t], dist[i])
    
    return max_dist[1:]

n, m, k = many_ii()
a = li()
g = [[] for _ in range(n + 1)]
for _ in range(m):
    x, y = many_ii()
    g[x].append(y)
    g[y].append(x)
ans = solve(n, g, a, k)
print(' '.join(map(str, ans)))