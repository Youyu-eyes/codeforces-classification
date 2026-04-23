# from math import inf
# from math import gcd
# from math import lcm
# from math import sin
# from math import pi
# from math import sqrt
# from math import isqrt

# from heapq import heapify
# from heapq import heappush
# from heapq import heappop
# from heapq import heapreplace
# from heapq import nsmallest
# from heapq import nlargest

# from heapq import heapify_max
# from heapq import heappush_max
# from heapq import heappop_max
# from heapq import heapreplace_max
# from heapq import nsmallest_max
# from heapq import nlargest_max

# from functools import cache

# from itertools import accumulate
# from itertools import zip_longest
# from itertools import pairwise
# from itertools import combinations
# from itertools import permutations

# from bisect import bisect_left
# from bisect import bisect_right

# from collections import defaultdict
# from collections import Counter
# from collections import deque

# import datetime
# from typing import List, Tuple

# MOD = 1_000_000_007

class UnionFind:
    def __init__(self, n: int):
        self._fa = list(range(n))
        self._size = [1] * n
        self.cc = n
        self.lazy = [0] * n
        self.d = [0] * n

    def find(self, x: int) -> int:
        fa = self._fa
        if fa[x] != x:
            root = self.find(fa[x])
            self.d[x] += self.d[fa[x]]
            fa[x] = root
        return fa[x]

    def is_same(self, x: int, y: int) -> bool:
        return self.find(x) == self.find(y)

    def merge(self, from_: int, to: int) -> bool:
        x, y = self.find(from_), self.find(to)
        if x == y:
            return False

        self.d[x] = self.lazy[x] - self.lazy[y]
        self._fa[x] = y
        self._size[y] += self._size[x]
        self.cc -= 1
        return True

    def get_size(self, x: int) -> int:
        return self._size[self.find(x)]

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII(b = 0):
    return list(MII(b))


for _ in range(II()):
    n, Q = MII()
    val = [0] * (n + 1)
    uf = UnionFind(n + 1)
    for q in range(Q):
        cur = LII()
        if cur[0] == 1:
            x, y = cur[1], cur[2]
            uf.merge(x, y)
        elif cur[0] == 2:
            x, a = cur[1], cur[2]
            val[x] += a
        elif cur[0] == 3:
            x, a = cur[1], cur[2]
            root = uf.find(x)
            uf.lazy[root] += a
        else:
            root = uf.find(x)
            ans = val[x] + uf.d[x] + uf.lazy[root]
            print(ans)