# https://codeforces.com/contest/2145/problem/C

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
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(s):
    d = s.count('a') - s.count('b')
    if d == 0:
        return 0
    cnt = defaultdict(int)
    cnt[d - 1] = -1
    pre = 0
    ans = inf
    for i, c in enumerate(s):
        if c == 'b':
            pre += 1
        x = i - 2 * pre
        if x in cnt:
            ans = min(ans, i - cnt[x])
        cnt[d + x] = i
    return ans if ans < n else -1


T = ii()
for _ in range(T):
    n = ii()
    s = si()
    print(solve(s))