from math import inf, gcd, lcm, sin, pi, sqrt, isqrt

from heapq import heapify, heappush, heappop, heapreplace, nsmallest, nlargst

from heapq import heapify_max, heappush_max, heappop_max, heapreplace_max, nsmallest_max, nlargest_max

from functools import cache, cmp_to_key

from itertools import accumulate, zip_longest, pairwise, combinations, permutations, groupby

from bisect import bisect_left, bisect_right

from collections import defaultdict, Counter, deque

from typing import List, Tuple

MOD = 1_000_000_007

from sortedcontainers import SortedSet, SortedDict, SortedList
from datetime import *
t1 = date(year = 2025, month = 3, day = 1)
t2 = date(year = 2025, month = 3, day = 17)
print(t2 - t1)  # 16 days, 0:00:00
print((t2 - t1).days)  # 16

# 参数也可简化
t1 = date(2025, 1, 1)
t2 = date(2025, 3, 17)
print((t2 - t1).days)  # 75

# 获取当前日期
print(date.today())

t1 = date(2025, 3, 17)
print(t1 + timedelta(days = 2))  # 2025-03-19
print(t1 + timedelta(1))  # 2025-03-18
print(t1.weekday())  # 0-6, 0代表周一

# bitwise 注意点
# Python  ^ 运算 比 & 运算低一级

# 自定义排序
# https://leetcode.cn/problems/maximum-value-of-concatenated-binary-segments/description/
MOD = 1_000_000_007

def compare(a, b):
    a1, a0 = a
    b1, b0 = b
    val_a = ((1 << a1) - 1) << a0
    val_b = ((1 << b1) - 1) << b0
    len_a = a1 + a0
    len_b = b1 + b0
    ab = (val_a << len_b) + val_b
    ba = (val_b << len_a) + val_a

    if ab > ba:
        return -1
    elif ab < ba:
        return 1
    return 0

class Solution:
    def maxValue(self, nums1: list[int], nums0: list[int]) -> int:
        n = len(nums1)

        cur = [(nums1[i], nums0[i]) for i in range(n)]
        cur.sort(key = cmp_to_key(compare))

        ans = 0
        for x, y in cur:
            ans = (ans * pow(2, x + y, MOD) % MOD + (pow(2, x, MOD) - 1) * pow(2, y, MOD)) % MOD
        return ans