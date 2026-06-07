# https://codeforces.com/problemset/problem/2170/C
# 好题
# 1300
# 贪心 + 式子变形：x = q(r + 1) + r <= k ==> qr + q + r <= k ==> (q + 1)(r + 1) <= (k + 1) ==> (r + 1) <= (k + 1) // (q + 1)
# 排序 + 双指针 + 二分: 在 r 数组中二分，如果 r[p2] 不满足题意 q[p1]，q[p1] 变大，r[p2] 更不满足条件，可以直接缩小二分范围

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

def solve(n, k, lq, lr):
    p1 = 0
    p2 = n - 1
    ans = 0
    lq.sort()
    lr.sort()
    while p1 < n and p2 >= 0:
        q = lq[p1]
        p2 = bisect_left(lr, (k + 1) // (q + 1) + 1, 0, p2 + 1) - 1
        if p2 >= 0:
            ans += 1
            p2 -= 1
        p1 += 1
    return ans

T = ii()
for _ in range(T):
    n, k = many_ii()
    q = lii()
    r = lii()
    for i in range(n):
        r[i] += 1
    print(solve(n, k, q, r))