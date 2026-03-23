# https://codeforces.com/contest/2161/problem/C
# 1200

import sys
from math import inf
from math import gcd
from math import lcm
from math import sin
from math import pi
import heapq as h
from functools import cache
from itertools import accumulate
from itertools import zip_longest
from bisect import bisect_left
from bisect import bisect_right
from collections import defaultdict
from collections import Counter
from collections import deque as q
import datetime
import re
input = lambda: sys.stdin.readline().strip()
si = lambda: input()
ii = lambda: int(input())
many_ii = lambda: map(int, input().split())
li = lambda: list(many_ii())
MOD = 1_000_000_007

def solve(nums, X):
    n = len(nums)
    cur = s = l = 0
    r = n - 1
    nums.sort()
    profit = 0
    ans = [0] * n
    for i in range(n):
        if (s + nums[r]) // X > 0:
            profit += nums[r]
            cur += (s + nums[r]) // X
            s = (s + nums[r]) % X
            ans[i] = nums[r]
            r -= 1
        else:
            s = (s + nums[l]) % X
            ans[i] = nums[l]
            l += 1
    return profit, ans

T = ii()
for _ in range(T):
    n, X = many_ii()
    nums = li()
    profit, ans = solve(nums, X)
    print(profit)
    for x in ans:
        print(x, end = ' ')
    print()