# from math import inf, gcd, lcm, sin, pi, sqrt, isqrt

# from heapq import heapify, heappush, heappop, heapreplace, nsmallest, nlargst

# from heapq import heapify_max, heappush_max, heappop_max, heapreplace_max, nsmallest_max, nlargest_max

# from functools import cache

# from itertools import accumulate, zip_longest, pairwise, combinations, permutations

# from bisect import bisect_left, bisect_right

# from collections import defaultdict, Counter, deque

# from sortedcontainers import SortedSet, SortedDict, SortedList
# import datetime
# from typing import List, Tuple

# MOD = 1_000_000_007

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
    print()