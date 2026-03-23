# https://codeforces.com/contest/1526/problem/B
# 1400

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

def solve(nums):
    if nums % 11 == 0 or nums % 111 == 0:
        return True
    n = nums // 111
    for i in range(n + 1):
        if (nums - i * 111) % 11 == 0:
            return True
    return False

T = ii()
for _ in range(T):
    n = ii()
    print('YES' if solve(n) else 'NO')