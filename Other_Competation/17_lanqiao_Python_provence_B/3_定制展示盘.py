# from math import inf
# from math import gcd
# from math import lcm
# from math import sin
# from math import pi
from math import isqrt
from math import sqrt

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

# from collections import defaultdict
# from collections import Counter
# from collections import deque

# import datetime
# import re
# from typing import List, Tuple

# MOD = 1_000_000_007

import sys
input = lambda: sys.stdin.readline().rstrip('\r\n')
print = lambda *args, end='\n', sep=' ': sys.stdout.write(
    sep.join(map(str, args)) + end
)

def is_prime(n: int) -> bool:
    for i in range(2, isqrt(n) + 1):
        if n % i == 0:
            return False
    return n >= 2


def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII(b = 0):
    return list(MII(b))

for _ in range(II()):
    n = II()
    if n < 3:
        print(4)
        continue
    if n & 1 and is_prime(n):
        print(n + 1)
    else:
        print(n)