# https://codeforces.com/contest/2165/problem/A
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
from itertools import pairwise
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

def solve(nums):
    nums += [nums[0]]
    ans = -max(nums)
    for x, y in pairwise(nums):
        ans += max(x, y)
    return ans
T = ii()
for _ in range(T):
    n = ii()
    nums = li()
    ans = solve(nums)
    print(ans)