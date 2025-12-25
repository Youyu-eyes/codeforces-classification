# https://codeforces.com/problemset/problem/2033/E
# 1400
# 置换环问题

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

def solve(n, p):
    for i in range(n):
        p[i] -= 1
    
    vis = [False] * n

    # 递归 DFS 在本题行不通
    # 因为数据范围为10**6，而 python 的递归最大深度为2e5，大约 10**5.3 次
    # 出题人有意卡掉 DFS 的递归写法

    def backtrack():
        ans = 0
        def dfs(i):
            vis[i] = True
            length = 1
            if not vis[p[i]]:
                length += dfs(p[i])
            return length

        for i in range(n):
            if not vis[i]:
                l = dfs(i)
                if l > 2:
                    ans += (l - 1) // 2
        return ans

    def iteration():
        ans = 0
        for i in range(n):
            if not vis[i]:
                length = 0
                j = i
                while not vis[j]:
                    vis[j] = True
                    j = p[j]
                    length += 1
                
                if length > 2:
                    ans += (length - 1) // 2
        return ans

    return iteration()

T = ii()
for _ in range(T):
    n = ii()
    p = lii()
    print(solve(n, p))