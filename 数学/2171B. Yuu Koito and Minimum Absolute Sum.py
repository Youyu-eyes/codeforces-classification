# https://codeforces.com/contest/2171/problem/B

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
input = lambda: sys.stdin.readline().strip()
MOD = 1_000_000_007

def solve(nums):
    diff = abs(nums[0] - nums[-1])
    if nums[0] == -1 and nums[-1] >= 0:
        nums[0] = nums[-1]
        diff = 0
    if nums[-1] == -1 and nums[0] >= 0:
        nums[-1] = nums[0]
        diff = 0
    for i, x in enumerate(nums):
        if x == -1:
            nums[i] = 0
    return nums, diff

T = int(input())
for _ in range(T):
    n = int(input())
    nums, diff = solve(list(map(int, input().split())))
    print(diff)
    for x in nums:
        print(x, end = ' ')
    print()