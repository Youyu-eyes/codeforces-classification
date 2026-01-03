# https://codeforces.com/problemset/problem/2118/C
# 1300
# 位运算
# 枚举非置位

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

def solve(k, nums):
    ans = 0
    i = 0
    max_len = max(x.bit_length() for x in nums)
    while (k >= 1 << i) or i < max_len:
        for x in nums:
            if x >> i & 1:
                ans += 1
            elif k >= 1 << i:
                ans += 1
                k -= 1 << i
        i += 1

    return ans

T = ii()
for _ in range(T):
    n, k = many_ii()
    nums = lii()
    print(solve(k, nums))