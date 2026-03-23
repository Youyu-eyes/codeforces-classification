# https://codeforces.com/contest/2112/problem/C
# 双指针做法
# 时间复杂度 O(n^2)，python 能过
# 倒序枚举 j 能让 k 递增，递减也可，但是不如递增好写

# from math import inf
# from math import gcd
# from math import lcm
# from math import sin
# from math import pi

# from heapq import heapify
# from heapq import heappush
# from heapq import heappop
# from heapq import heapreplace
# from heapq import nsmallest
# from heapq import nlargest

# from functools import cache

# from itertools import accumulate
# from itertools import zip_longest
# from itertools import pairwise

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

def II():
    return int(input())

def MII(b = 0):
    return map(lambda x: int(x) - b, input().split())

def LII():
    return list(MII())

for _ in range(II()):
    n = II()
    nums = LII()
    mx = nums[-1]
    ans = 0
    for i in range(2, n):
        k = 0
        for j in range(i - 1, 0, -1):
            while k < j and nums[k] <= max(mx, 2 * nums[i]) - nums[i] - nums[j]:
                k += 1
            if j == k:
                break
            ans += j - k
    print(ans)